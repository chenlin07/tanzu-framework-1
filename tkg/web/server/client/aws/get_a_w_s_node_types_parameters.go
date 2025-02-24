// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAWSNodeTypesParams creates a new GetAWSNodeTypesParams object
// with the default values initialized.
func NewGetAWSNodeTypesParams() *GetAWSNodeTypesParams {
	var ()
	return &GetAWSNodeTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAWSNodeTypesParamsWithTimeout creates a new GetAWSNodeTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAWSNodeTypesParamsWithTimeout(timeout time.Duration) *GetAWSNodeTypesParams {
	var ()
	return &GetAWSNodeTypesParams{

		timeout: timeout,
	}
}

// NewGetAWSNodeTypesParamsWithContext creates a new GetAWSNodeTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAWSNodeTypesParamsWithContext(ctx context.Context) *GetAWSNodeTypesParams {
	var ()
	return &GetAWSNodeTypesParams{

		Context: ctx,
	}
}

// NewGetAWSNodeTypesParamsWithHTTPClient creates a new GetAWSNodeTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAWSNodeTypesParamsWithHTTPClient(client *http.Client) *GetAWSNodeTypesParams {
	var ()
	return &GetAWSNodeTypesParams{
		HTTPClient: client,
	}
}

/*
GetAWSNodeTypesParams contains all the parameters to send to the API endpoint
for the get a w s node types operation typically these are written to a http.Request
*/
type GetAWSNodeTypesParams struct {

	/*Az
	  AWS availability zone, e.g. us-west-2

	*/
	Az *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get a w s node types params
func (o *GetAWSNodeTypesParams) WithTimeout(timeout time.Duration) *GetAWSNodeTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get a w s node types params
func (o *GetAWSNodeTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get a w s node types params
func (o *GetAWSNodeTypesParams) WithContext(ctx context.Context) *GetAWSNodeTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get a w s node types params
func (o *GetAWSNodeTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get a w s node types params
func (o *GetAWSNodeTypesParams) WithHTTPClient(client *http.Client) *GetAWSNodeTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get a w s node types params
func (o *GetAWSNodeTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAz adds the az to the get a w s node types params
func (o *GetAWSNodeTypesParams) WithAz(az *string) *GetAWSNodeTypesParams {
	o.SetAz(az)
	return o
}

// SetAz adds the az to the get a w s node types params
func (o *GetAWSNodeTypesParams) SetAz(az *string) {
	o.Az = az
}

// WriteToRequest writes these params to a swagger request
func (o *GetAWSNodeTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Az != nil {

		// query param az
		var qrAz string
		if o.Az != nil {
			qrAz = *o.Az
		}
		qAz := qrAz
		if qAz != "" {
			if err := r.SetQueryParam("az", qAz); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
