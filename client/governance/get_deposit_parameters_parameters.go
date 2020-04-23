// Code generated by go-swagger; DO NOT EDIT.

package governance

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

// NewGetDepositParametersParams creates a new GetDepositParametersParams object
// with the default values initialized.
func NewGetDepositParametersParams() *GetDepositParametersParams {

	return &GetDepositParametersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDepositParametersParamsWithTimeout creates a new GetDepositParametersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDepositParametersParamsWithTimeout(timeout time.Duration) *GetDepositParametersParams {

	return &GetDepositParametersParams{

		timeout: timeout,
	}
}

// NewGetDepositParametersParamsWithContext creates a new GetDepositParametersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDepositParametersParamsWithContext(ctx context.Context) *GetDepositParametersParams {

	return &GetDepositParametersParams{

		Context: ctx,
	}
}

// NewGetDepositParametersParamsWithHTTPClient creates a new GetDepositParametersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDepositParametersParamsWithHTTPClient(client *http.Client) *GetDepositParametersParams {

	return &GetDepositParametersParams{
		HTTPClient: client,
	}
}

/*GetDepositParametersParams contains all the parameters to send to the API endpoint
for the get deposit parameters operation typically these are written to a http.Request
*/
type GetDepositParametersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deposit parameters params
func (o *GetDepositParametersParams) WithTimeout(timeout time.Duration) *GetDepositParametersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deposit parameters params
func (o *GetDepositParametersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deposit parameters params
func (o *GetDepositParametersParams) WithContext(ctx context.Context) *GetDepositParametersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deposit parameters params
func (o *GetDepositParametersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deposit parameters params
func (o *GetDepositParametersParams) WithHTTPClient(client *http.Client) *GetDepositParametersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deposit parameters params
func (o *GetDepositParametersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDepositParametersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
