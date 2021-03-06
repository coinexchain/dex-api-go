// Code generated by go-swagger; DO NOT EDIT.

package distribution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetDistributionParamsParams creates a new GetDistributionParamsParams object
// with the default values initialized.
func NewGetDistributionParamsParams() *GetDistributionParamsParams {

	return &GetDistributionParamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDistributionParamsParamsWithTimeout creates a new GetDistributionParamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDistributionParamsParamsWithTimeout(timeout time.Duration) *GetDistributionParamsParams {

	return &GetDistributionParamsParams{

		timeout: timeout,
	}
}

// NewGetDistributionParamsParamsWithContext creates a new GetDistributionParamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDistributionParamsParamsWithContext(ctx context.Context) *GetDistributionParamsParams {

	return &GetDistributionParamsParams{

		Context: ctx,
	}
}

// NewGetDistributionParamsParamsWithHTTPClient creates a new GetDistributionParamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDistributionParamsParamsWithHTTPClient(client *http.Client) *GetDistributionParamsParams {

	return &GetDistributionParamsParams{
		HTTPClient: client,
	}
}

/*GetDistributionParamsParams contains all the parameters to send to the API endpoint
for the get distribution params operation typically these are written to a http.Request
*/
type GetDistributionParamsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get distribution params params
func (o *GetDistributionParamsParams) WithTimeout(timeout time.Duration) *GetDistributionParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get distribution params params
func (o *GetDistributionParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get distribution params params
func (o *GetDistributionParamsParams) WithContext(ctx context.Context) *GetDistributionParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get distribution params params
func (o *GetDistributionParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get distribution params params
func (o *GetDistributionParamsParams) WithHTTPClient(client *http.Client) *GetDistributionParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get distribution params params
func (o *GetDistributionParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDistributionParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
