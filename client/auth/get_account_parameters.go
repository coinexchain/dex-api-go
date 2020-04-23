// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewGetAccountParams creates a new GetAccountParams object
// with the default values initialized.
func NewGetAccountParams() *GetAccountParams {
	var ()
	return &GetAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountParamsWithTimeout creates a new GetAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountParamsWithTimeout(timeout time.Duration) *GetAccountParams {
	var ()
	return &GetAccountParams{

		timeout: timeout,
	}
}

// NewGetAccountParamsWithContext creates a new GetAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountParamsWithContext(ctx context.Context) *GetAccountParams {
	var ()
	return &GetAccountParams{

		Context: ctx,
	}
}

// NewGetAccountParamsWithHTTPClient creates a new GetAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountParamsWithHTTPClient(client *http.Client) *GetAccountParams {
	var ()
	return &GetAccountParams{
		HTTPClient: client,
	}
}

/*GetAccountParams contains all the parameters to send to the API endpoint
for the get account operation typically these are written to a http.Request
*/
type GetAccountParams struct {

	/*Address
	  Account address

	*/
	Address string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account params
func (o *GetAccountParams) WithTimeout(timeout time.Duration) *GetAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account params
func (o *GetAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account params
func (o *GetAccountParams) WithContext(ctx context.Context) *GetAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account params
func (o *GetAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account params
func (o *GetAccountParams) WithHTTPClient(client *http.Client) *GetAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account params
func (o *GetAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the get account params
func (o *GetAccountParams) WithAddress(address string) *GetAccountParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the get account params
func (o *GetAccountParams) SetAddress(address string) {
	o.Address = address
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
