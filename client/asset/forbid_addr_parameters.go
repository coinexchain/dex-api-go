// Code generated by go-swagger; DO NOT EDIT.

package asset

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

	"github.com/coinexchain/dex-api-go/models"
)

// NewForbidAddrParams creates a new ForbidAddrParams object
// with the default values initialized.
func NewForbidAddrParams() *ForbidAddrParams {
	var ()
	return &ForbidAddrParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewForbidAddrParamsWithTimeout creates a new ForbidAddrParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewForbidAddrParamsWithTimeout(timeout time.Duration) *ForbidAddrParams {
	var ()
	return &ForbidAddrParams{

		timeout: timeout,
	}
}

// NewForbidAddrParamsWithContext creates a new ForbidAddrParams object
// with the default values initialized, and the ability to set a context for a request
func NewForbidAddrParamsWithContext(ctx context.Context) *ForbidAddrParams {
	var ()
	return &ForbidAddrParams{

		Context: ctx,
	}
}

// NewForbidAddrParamsWithHTTPClient creates a new ForbidAddrParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewForbidAddrParamsWithHTTPClient(client *http.Client) *ForbidAddrParams {
	var ()
	return &ForbidAddrParams{
		HTTPClient: client,
	}
}

/*ForbidAddrParams contains all the parameters to send to the API endpoint
for the forbid addr operation typically these are written to a http.Request
*/
type ForbidAddrParams struct {

	/*Addresses
	  forbidden addresses

	*/
	Addresses *models.Addresses
	/*Symbol
	  token symbol

	*/
	Symbol string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the forbid addr params
func (o *ForbidAddrParams) WithTimeout(timeout time.Duration) *ForbidAddrParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the forbid addr params
func (o *ForbidAddrParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the forbid addr params
func (o *ForbidAddrParams) WithContext(ctx context.Context) *ForbidAddrParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the forbid addr params
func (o *ForbidAddrParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the forbid addr params
func (o *ForbidAddrParams) WithHTTPClient(client *http.Client) *ForbidAddrParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the forbid addr params
func (o *ForbidAddrParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddresses adds the addresses to the forbid addr params
func (o *ForbidAddrParams) WithAddresses(addresses *models.Addresses) *ForbidAddrParams {
	o.SetAddresses(addresses)
	return o
}

// SetAddresses adds the addresses to the forbid addr params
func (o *ForbidAddrParams) SetAddresses(addresses *models.Addresses) {
	o.Addresses = addresses
}

// WithSymbol adds the symbol to the forbid addr params
func (o *ForbidAddrParams) WithSymbol(symbol string) *ForbidAddrParams {
	o.SetSymbol(symbol)
	return o
}

// SetSymbol adds the symbol to the forbid addr params
func (o *ForbidAddrParams) SetSymbol(symbol string) {
	o.Symbol = symbol
}

// WriteToRequest writes these params to a swagger request
func (o *ForbidAddrParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Addresses != nil {
		if err := r.SetBodyParam(o.Addresses); err != nil {
			return err
		}
	}

	// path param symbol
	if err := r.SetPathParam("symbol", o.Symbol); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
