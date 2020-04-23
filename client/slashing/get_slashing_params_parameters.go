// Code generated by go-swagger; DO NOT EDIT.

package slashing

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

// NewGetSlashingParamsParams creates a new GetSlashingParamsParams object
// with the default values initialized.
func NewGetSlashingParamsParams() *GetSlashingParamsParams {

	return &GetSlashingParamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSlashingParamsParamsWithTimeout creates a new GetSlashingParamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSlashingParamsParamsWithTimeout(timeout time.Duration) *GetSlashingParamsParams {

	return &GetSlashingParamsParams{

		timeout: timeout,
	}
}

// NewGetSlashingParamsParamsWithContext creates a new GetSlashingParamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSlashingParamsParamsWithContext(ctx context.Context) *GetSlashingParamsParams {

	return &GetSlashingParamsParams{

		Context: ctx,
	}
}

// NewGetSlashingParamsParamsWithHTTPClient creates a new GetSlashingParamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSlashingParamsParamsWithHTTPClient(client *http.Client) *GetSlashingParamsParams {

	return &GetSlashingParamsParams{
		HTTPClient: client,
	}
}

/*GetSlashingParamsParams contains all the parameters to send to the API endpoint
for the get slashing params operation typically these are written to a http.Request
*/
type GetSlashingParamsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get slashing params params
func (o *GetSlashingParamsParams) WithTimeout(timeout time.Duration) *GetSlashingParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get slashing params params
func (o *GetSlashingParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get slashing params params
func (o *GetSlashingParamsParams) WithContext(ctx context.Context) *GetSlashingParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get slashing params params
func (o *GetSlashingParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get slashing params params
func (o *GetSlashingParamsParams) WithHTTPClient(client *http.Client) *GetSlashingParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get slashing params params
func (o *GetSlashingParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSlashingParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
