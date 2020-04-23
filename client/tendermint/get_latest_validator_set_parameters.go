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

// NewGetLatestValidatorSetParams creates a new GetLatestValidatorSetParams object
// with the default values initialized.
func NewGetLatestValidatorSetParams() *GetLatestValidatorSetParams {

	return &GetLatestValidatorSetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLatestValidatorSetParamsWithTimeout creates a new GetLatestValidatorSetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLatestValidatorSetParamsWithTimeout(timeout time.Duration) *GetLatestValidatorSetParams {

	return &GetLatestValidatorSetParams{

		timeout: timeout,
	}
}

// NewGetLatestValidatorSetParamsWithContext creates a new GetLatestValidatorSetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLatestValidatorSetParamsWithContext(ctx context.Context) *GetLatestValidatorSetParams {

	return &GetLatestValidatorSetParams{

		Context: ctx,
	}
}

// NewGetLatestValidatorSetParamsWithHTTPClient creates a new GetLatestValidatorSetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLatestValidatorSetParamsWithHTTPClient(client *http.Client) *GetLatestValidatorSetParams {

	return &GetLatestValidatorSetParams{
		HTTPClient: client,
	}
}

/*GetLatestValidatorSetParams contains all the parameters to send to the API endpoint
for the get latest validator set operation typically these are written to a http.Request
*/
type GetLatestValidatorSetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get latest validator set params
func (o *GetLatestValidatorSetParams) WithTimeout(timeout time.Duration) *GetLatestValidatorSetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get latest validator set params
func (o *GetLatestValidatorSetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get latest validator set params
func (o *GetLatestValidatorSetParams) WithContext(ctx context.Context) *GetLatestValidatorSetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get latest validator set params
func (o *GetLatestValidatorSetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get latest validator set params
func (o *GetLatestValidatorSetParams) WithHTTPClient(client *http.Client) *GetLatestValidatorSetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get latest validator set params
func (o *GetLatestValidatorSetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLatestValidatorSetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
