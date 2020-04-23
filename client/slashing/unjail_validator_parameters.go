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

// NewUnjailValidatorParams creates a new UnjailValidatorParams object
// with the default values initialized.
func NewUnjailValidatorParams() *UnjailValidatorParams {
	var ()
	return &UnjailValidatorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnjailValidatorParamsWithTimeout creates a new UnjailValidatorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnjailValidatorParamsWithTimeout(timeout time.Duration) *UnjailValidatorParams {
	var ()
	return &UnjailValidatorParams{

		timeout: timeout,
	}
}

// NewUnjailValidatorParamsWithContext creates a new UnjailValidatorParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnjailValidatorParamsWithContext(ctx context.Context) *UnjailValidatorParams {
	var ()
	return &UnjailValidatorParams{

		Context: ctx,
	}
}

// NewUnjailValidatorParamsWithHTTPClient creates a new UnjailValidatorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnjailValidatorParamsWithHTTPClient(client *http.Client) *UnjailValidatorParams {
	var ()
	return &UnjailValidatorParams{
		HTTPClient: client,
	}
}

/*UnjailValidatorParams contains all the parameters to send to the API endpoint
for the unjail validator operation typically these are written to a http.Request
*/
type UnjailValidatorParams struct {

	/*UnjailBody*/
	UnjailBody UnjailValidatorBody
	/*ValidatorAddr
	  Bech32 validator address

	*/
	ValidatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the unjail validator params
func (o *UnjailValidatorParams) WithTimeout(timeout time.Duration) *UnjailValidatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unjail validator params
func (o *UnjailValidatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unjail validator params
func (o *UnjailValidatorParams) WithContext(ctx context.Context) *UnjailValidatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unjail validator params
func (o *UnjailValidatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unjail validator params
func (o *UnjailValidatorParams) WithHTTPClient(client *http.Client) *UnjailValidatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unjail validator params
func (o *UnjailValidatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUnjailBody adds the unjailBody to the unjail validator params
func (o *UnjailValidatorParams) WithUnjailBody(unjailBody UnjailValidatorBody) *UnjailValidatorParams {
	o.SetUnjailBody(unjailBody)
	return o
}

// SetUnjailBody adds the unjailBody to the unjail validator params
func (o *UnjailValidatorParams) SetUnjailBody(unjailBody UnjailValidatorBody) {
	o.UnjailBody = unjailBody
}

// WithValidatorAddr adds the validatorAddr to the unjail validator params
func (o *UnjailValidatorParams) WithValidatorAddr(validatorAddr string) *UnjailValidatorParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the unjail validator params
func (o *UnjailValidatorParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *UnjailValidatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.UnjailBody); err != nil {
		return err
	}

	// path param validatorAddr
	if err := r.SetPathParam("validatorAddr", o.ValidatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
