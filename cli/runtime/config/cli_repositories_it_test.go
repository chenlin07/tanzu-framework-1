// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	configapi "github.com/vmware-tanzu/tanzu-framework/cli/runtime/apis/config/v1alpha1"
)

func setUpRepositoriesData() (string, string, string, string) {
	cfg := `clientOptions:
  cli:
    discoverySources:
      - oci:
          name: default
          image: "update-default-image"
          unknown: cli-unknown
        contextType: k8s
      - local:
          name: default-local
        contextType: k8s
      - local:
          name: admin-local
          path: admin
      - oci:
          name: new-default
          image: new-default-image
    repositories:
      - gcpPluginRepository:
          bucketName: tanzu-cli-framework
          name: core
          unknown: cli-unknown
servers:
  - name: test-mc
    type: managementcluster
    managementClusterOpts:
      endpoint: updated-test-endpoint
      path: updated-test-path
      context: updated-test-context
      annotation: one
      required: true
    discoverySources:
      - gcp:
          name: test
          bucket: updated-test-bucket
          manifestPath: updated-test-manifest-path
          annotation: one
          required: true
        contextType: tmc
current: test-mc
contexts:
  - name: test-mc
    target: kubernetes
    group: one
    clusterOpts:
      isManagementCluster: true
      annotation: one
      required: true
      annotationStruct:
        one: one
      endpoint: updated-test-endpoint
      path: updated-test-path
      context: updated-test-context
    discoverySources:
      - gcp:
          name: test
          bucket: updated-test-bucket
          manifestPath: updated-test-manifest-path
          annotation: one
          required: true
        contextType: tmc
      - gcp:
          name: test-two
          bucket: updated-test-bucket
          manifestPath: updated-test-manifest-path
          annotation: two
          required: true
        contextType: tmc
currentContext:
  kubernetes: test-mc
`
	expectedCfg := `clientOptions:
    cli:
        discoverySources:
            - oci:
                name: default
                image: "update-default-image"
                unknown: cli-unknown
              contextType: k8s
            - local:
                name: default-local
              contextType: k8s
            - local:
                name: admin-local
                path: admin
            - oci:
                name: new-default
                image: new-default-image
        repositories:
            - gcpPluginRepository:
                bucketName: update-bucket
                name: core
                unknown: cli-unknown
                rootPath: new-root-path
            - gcpPluginRepository:
                name: new-repo
                bucketName: new-bucket
                rootPath: new-root-path
servers:
    - name: test-mc
      type: managementcluster
      managementClusterOpts:
        endpoint: updated-test-endpoint
        path: updated-test-path
        context: updated-test-context
        annotation: one
        required: true
      discoverySources:
        - gcp:
            name: test
            bucket: updated-test-bucket
            manifestPath: updated-test-manifest-path
            annotation: one
            required: true
          contextType: tmc
current: test-mc
`

	cfg2 := `contexts:
  - name: test-mc
    target: kubernetes
    group: one
    clusterOpts:
      isManagementCluster: true
      annotation: one
      required: true
      annotationStruct:
        one: one
      endpoint: updated-test-endpoint
      path: updated-test-path
      context: updated-test-context
    discoverySources:
      - gcp:
          name: test
          bucket: updated-test-bucket
          manifestPath: updated-test-manifest-path
          annotation: one
          required: true
        contextType: tmc
      - gcp:
          name: test-two
          bucket: updated-test-bucket
          manifestPath: updated-test-manifest-path
          annotation: two
          required: true
        contextType: tmc
currentContext:
    kubernetes: test-mc
`
	expectedCfg2 := `contexts:
    - name: test-mc
      target: kubernetes
      group: one
      clusterOpts:
        isManagementCluster: true
        annotation: one
        required: true
        annotationStruct:
            one: one
        endpoint: updated-test-endpoint
        path: updated-test-path
        context: updated-test-context
      discoverySources:
        - gcp:
            name: test
            bucket: updated-test-bucket
            manifestPath: updated-test-manifest-path
            annotation: one
            required: true
          contextType: tmc
        - gcp:
            name: test-two
            bucket: updated-test-bucket
            manifestPath: updated-test-manifest-path
            annotation: two
            required: true
          contextType: tmc
currentContext:
    kubernetes: test-mc
`

	return cfg, expectedCfg, cfg2, expectedCfg2
}

func TestCLIRepositoriesIntegration(t *testing.T) {
	// Setup config data
	cfg, expectedCfg, cfg2, expectedCfg2 := setUpRepositoriesData()
	cfgFiles, cleanUp := setupTestConfig(t, &CfgTestData{cfg: cfg, cfgNextGen: cfg2})

	defer func() {
		cleanUp()
	}()

	// Get CLI Repositories
	repos, err := GetCLIRepositories()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(repos))
	// Add new CLI Repository
	newRepo := &configapi.PluginRepository{
		GCPPluginRepository: &configapi.GCPPluginRepository{
			Name:       "new-repo",
			BucketName: "new-bucket",
			RootPath:   "new-root-path",
		},
	}
	err = SetCLIRepository(*newRepo)
	assert.NoError(t, err)
	repos, err = GetCLIRepositories()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(repos))
	// Should not persist on adding same CLI Repository
	err = SetCLIRepository(*newRepo)
	assert.NoError(t, err)
	repos, err = GetCLIRepositories()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(repos))
	// Update existing CLI Repository
	existingRepo := &configapi.PluginRepository{
		GCPPluginRepository: &configapi.GCPPluginRepository{
			Name:       "core",
			BucketName: "update-bucket",
			RootPath:   "new-root-path",
		},
	}
	err = SetCLIRepository(*existingRepo)
	assert.NoError(t, err)
	repo, err := GetCLIRepository("core")
	assert.Nil(t, err)
	assert.NotNil(t, repo)
	assert.Equal(t, existingRepo.GCPPluginRepository, repo.GCPPluginRepository)

	file, err := os.ReadFile(cfgFiles[0].Name())
	assert.NoError(t, err)
	assert.Equal(t, expectedCfg, string(file))

	file, err = os.ReadFile(cfgFiles[1].Name())
	assert.NoError(t, err)
	assert.Equal(t, expectedCfg2, string(file))

	// Delete CLI Repository
	err = DeleteCLIRepository("core")
	assert.NoError(t, err)
	repos, err = GetCLIRepositories()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(repos))
	repo, err = GetCLIRepository("core")
	assert.Equal(t, errors.New("cli repository not found").Error(), err.Error())
	assert.Nil(t, repo)
}
