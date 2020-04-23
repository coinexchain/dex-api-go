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

// NewQueryTradingPairsParams creates a new QueryTradingPairsParams object
// with the default values initialized.
func NewQueryTradingPairsParams() *QueryTradingPairsParams {

	return &QueryTradingPairsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryTradingPairsParamsWithTimeout creates a new QueryTradingPairsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryTradingPairsParamsWithTimeout(timeout time.Duration) *QueryTradingPairsParams {

	return &QueryTradingPairsParams{

		timeout: timeout,
	}
}

// NewQueryTradingPairsParamsWithContext creates a new QueryTradingPairsParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryTradingPairsParamsWithContext(ctx context.Context) *QueryTradingPairsParams {

	return &QueryTradingPairsParams{

		Context: ctx,
	}
}

// NewQueryTradingPairsParamsWithHTTPClient creates a new QueryTradingPairsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryTradingPairsParamsWithHTTPClient(client *http.Client) *QueryTradingPairsParams {

	return &QueryTradingPairsParams{
		HTTPClient: client,
	}
}

/*QueryTradingPairsParams contains all the parameters to send to the API endpoint
for the query trading pairs operation typically these are written to a http.Request
*/
type QueryTradingPairsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query trading pairs params
func (o *QueryTradingPairsParams) WithTimeout(timeout time.Duration) *QueryTradingPairsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query trading pairs params
func (o *QueryTradingPairsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query trading pairs params
func (o *QueryTradingPairsParams) WithContext(ctx context.Context) *QueryTradingPairsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query trading pairs params
func (o *QueryTradingPairsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query trading pairs params
func (o *QueryTradingPairsParams) WithHTTPClient(client *http.Client) *QueryTradingPairsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query trading pairs params
func (o *QueryTradingPairsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *QueryTradingPairsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
