// Code generated by go-swagger; DO NOT EDIT.

package market

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

// NewGetOrdersParams creates a new GetOrdersParams object
// with the default values initialized.
func NewGetOrdersParams() *GetOrdersParams {
	var ()
	return &GetOrdersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrdersParamsWithTimeout creates a new GetOrdersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrdersParamsWithTimeout(timeout time.Duration) *GetOrdersParams {
	var ()
	return &GetOrdersParams{

		timeout: timeout,
	}
}

// NewGetOrdersParamsWithContext creates a new GetOrdersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrdersParamsWithContext(ctx context.Context) *GetOrdersParams {
	var ()
	return &GetOrdersParams{

		Context: ctx,
	}
}

// NewGetOrdersParamsWithHTTPClient creates a new GetOrdersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrdersParamsWithHTTPClient(client *http.Client) *GetOrdersParams {
	var ()
	return &GetOrdersParams{
		HTTPClient: client,
	}
}

/*GetOrdersParams contains all the parameters to send to the API endpoint
for the get orders operation typically these are written to a http.Request
*/
type GetOrdersParams struct {

	/*Address
	  The user address

	*/
	Address string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get orders params
func (o *GetOrdersParams) WithTimeout(timeout time.Duration) *GetOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orders params
func (o *GetOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orders params
func (o *GetOrdersParams) WithContext(ctx context.Context) *GetOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orders params
func (o *GetOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orders params
func (o *GetOrdersParams) WithHTTPClient(client *http.Client) *GetOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orders params
func (o *GetOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the get orders params
func (o *GetOrdersParams) WithAddress(address string) *GetOrdersParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the get orders params
func (o *GetOrdersParams) SetAddress(address string) {
	o.Address = address
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
