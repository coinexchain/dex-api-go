// Code generated by go-swagger; DO NOT EDIT.

package tendermint

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

// NewGetLatestBlockParams creates a new GetLatestBlockParams object
// with the default values initialized.
func NewGetLatestBlockParams() *GetLatestBlockParams {

	return &GetLatestBlockParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLatestBlockParamsWithTimeout creates a new GetLatestBlockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLatestBlockParamsWithTimeout(timeout time.Duration) *GetLatestBlockParams {

	return &GetLatestBlockParams{

		timeout: timeout,
	}
}

// NewGetLatestBlockParamsWithContext creates a new GetLatestBlockParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLatestBlockParamsWithContext(ctx context.Context) *GetLatestBlockParams {

	return &GetLatestBlockParams{

		Context: ctx,
	}
}

// NewGetLatestBlockParamsWithHTTPClient creates a new GetLatestBlockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLatestBlockParamsWithHTTPClient(client *http.Client) *GetLatestBlockParams {

	return &GetLatestBlockParams{
		HTTPClient: client,
	}
}

/*GetLatestBlockParams contains all the parameters to send to the API endpoint
for the get latest block operation typically these are written to a http.Request
*/
type GetLatestBlockParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get latest block params
func (o *GetLatestBlockParams) WithTimeout(timeout time.Duration) *GetLatestBlockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get latest block params
func (o *GetLatestBlockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get latest block params
func (o *GetLatestBlockParams) WithContext(ctx context.Context) *GetLatestBlockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get latest block params
func (o *GetLatestBlockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get latest block params
func (o *GetLatestBlockParams) WithHTTPClient(client *http.Client) *GetLatestBlockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get latest block params
func (o *GetLatestBlockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLatestBlockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}