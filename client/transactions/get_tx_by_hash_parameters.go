// Code generated by go-swagger; DO NOT EDIT.

package transactions

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

// NewGetTxByHashParams creates a new GetTxByHashParams object
// with the default values initialized.
func NewGetTxByHashParams() *GetTxByHashParams {
	var ()
	return &GetTxByHashParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTxByHashParamsWithTimeout creates a new GetTxByHashParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTxByHashParamsWithTimeout(timeout time.Duration) *GetTxByHashParams {
	var ()
	return &GetTxByHashParams{

		timeout: timeout,
	}
}

// NewGetTxByHashParamsWithContext creates a new GetTxByHashParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTxByHashParamsWithContext(ctx context.Context) *GetTxByHashParams {
	var ()
	return &GetTxByHashParams{

		Context: ctx,
	}
}

// NewGetTxByHashParamsWithHTTPClient creates a new GetTxByHashParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTxByHashParamsWithHTTPClient(client *http.Client) *GetTxByHashParams {
	var ()
	return &GetTxByHashParams{
		HTTPClient: client,
	}
}

/*GetTxByHashParams contains all the parameters to send to the API endpoint
for the get tx by hash operation typically these are written to a http.Request
*/
type GetTxByHashParams struct {

	/*Hash
	  Tx hash

	*/
	Hash string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tx by hash params
func (o *GetTxByHashParams) WithTimeout(timeout time.Duration) *GetTxByHashParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tx by hash params
func (o *GetTxByHashParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tx by hash params
func (o *GetTxByHashParams) WithContext(ctx context.Context) *GetTxByHashParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tx by hash params
func (o *GetTxByHashParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tx by hash params
func (o *GetTxByHashParams) WithHTTPClient(client *http.Client) *GetTxByHashParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tx by hash params
func (o *GetTxByHashParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHash adds the hash to the get tx by hash params
func (o *GetTxByHashParams) WithHash(hash string) *GetTxByHashParams {
	o.SetHash(hash)
	return o
}

// SetHash adds the hash to the get tx by hash params
func (o *GetTxByHashParams) SetHash(hash string) {
	o.Hash = hash
}

// WriteToRequest writes these params to a swagger request
func (o *GetTxByHashParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param hash
	if err := r.SetPathParam("hash", o.Hash); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
