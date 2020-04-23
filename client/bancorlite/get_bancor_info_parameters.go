// Code generated by go-swagger; DO NOT EDIT.

package bancorlite

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

// NewGetBancorInfoParams creates a new GetBancorInfoParams object
// with the default values initialized.
func NewGetBancorInfoParams() *GetBancorInfoParams {
	var ()
	return &GetBancorInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBancorInfoParamsWithTimeout creates a new GetBancorInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBancorInfoParamsWithTimeout(timeout time.Duration) *GetBancorInfoParams {
	var ()
	return &GetBancorInfoParams{

		timeout: timeout,
	}
}

// NewGetBancorInfoParamsWithContext creates a new GetBancorInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBancorInfoParamsWithContext(ctx context.Context) *GetBancorInfoParams {
	var ()
	return &GetBancorInfoParams{

		Context: ctx,
	}
}

// NewGetBancorInfoParamsWithHTTPClient creates a new GetBancorInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBancorInfoParamsWithHTTPClient(client *http.Client) *GetBancorInfoParams {
	var ()
	return &GetBancorInfoParams{
		HTTPClient: client,
	}
}

/*GetBancorInfoParams contains all the parameters to send to the API endpoint
for the get bancor info operation typically these are written to a http.Request
*/
type GetBancorInfoParams struct {

	/*Symbol
	  stock and money pair

	*/
	Symbol string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bancor info params
func (o *GetBancorInfoParams) WithTimeout(timeout time.Duration) *GetBancorInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bancor info params
func (o *GetBancorInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bancor info params
func (o *GetBancorInfoParams) WithContext(ctx context.Context) *GetBancorInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bancor info params
func (o *GetBancorInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bancor info params
func (o *GetBancorInfoParams) WithHTTPClient(client *http.Client) *GetBancorInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bancor info params
func (o *GetBancorInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSymbol adds the symbol to the get bancor info params
func (o *GetBancorInfoParams) WithSymbol(symbol string) *GetBancorInfoParams {
	o.SetSymbol(symbol)
	return o
}

// SetSymbol adds the symbol to the get bancor info params
func (o *GetBancorInfoParams) SetSymbol(symbol string) {
	o.Symbol = symbol
}

// WriteToRequest writes these params to a swagger request
func (o *GetBancorInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param symbol
	if err := r.SetPathParam("symbol", o.Symbol); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}