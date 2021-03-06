// Code generated by go-swagger; DO NOT EDIT.

package distribution

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

// NewSetWithdrawAddressParams creates a new SetWithdrawAddressParams object
// with the default values initialized.
func NewSetWithdrawAddressParams() *SetWithdrawAddressParams {
	var ()
	return &SetWithdrawAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetWithdrawAddressParamsWithTimeout creates a new SetWithdrawAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetWithdrawAddressParamsWithTimeout(timeout time.Duration) *SetWithdrawAddressParams {
	var ()
	return &SetWithdrawAddressParams{

		timeout: timeout,
	}
}

// NewSetWithdrawAddressParamsWithContext creates a new SetWithdrawAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetWithdrawAddressParamsWithContext(ctx context.Context) *SetWithdrawAddressParams {
	var ()
	return &SetWithdrawAddressParams{

		Context: ctx,
	}
}

// NewSetWithdrawAddressParamsWithHTTPClient creates a new SetWithdrawAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetWithdrawAddressParamsWithHTTPClient(client *http.Client) *SetWithdrawAddressParams {
	var ()
	return &SetWithdrawAddressParams{
		HTTPClient: client,
	}
}

/*SetWithdrawAddressParams contains all the parameters to send to the API endpoint
for the set withdraw address operation typically these are written to a http.Request
*/
type SetWithdrawAddressParams struct {

	/*WithdrawRequestBody*/
	WithdrawRequestBody SetWithdrawAddressBody
	/*DelegatorAddr
	  Bech32 AccAddress of Delegator

	*/
	DelegatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set withdraw address params
func (o *SetWithdrawAddressParams) WithTimeout(timeout time.Duration) *SetWithdrawAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set withdraw address params
func (o *SetWithdrawAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set withdraw address params
func (o *SetWithdrawAddressParams) WithContext(ctx context.Context) *SetWithdrawAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set withdraw address params
func (o *SetWithdrawAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set withdraw address params
func (o *SetWithdrawAddressParams) WithHTTPClient(client *http.Client) *SetWithdrawAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set withdraw address params
func (o *SetWithdrawAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWithdrawRequestBody adds the withdrawRequestBody to the set withdraw address params
func (o *SetWithdrawAddressParams) WithWithdrawRequestBody(withdrawRequestBody SetWithdrawAddressBody) *SetWithdrawAddressParams {
	o.SetWithdrawRequestBody(withdrawRequestBody)
	return o
}

// SetWithdrawRequestBody adds the withdrawRequestBody to the set withdraw address params
func (o *SetWithdrawAddressParams) SetWithdrawRequestBody(withdrawRequestBody SetWithdrawAddressBody) {
	o.WithdrawRequestBody = withdrawRequestBody
}

// WithDelegatorAddr adds the delegatorAddr to the set withdraw address params
func (o *SetWithdrawAddressParams) WithDelegatorAddr(delegatorAddr string) *SetWithdrawAddressParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the set withdraw address params
func (o *SetWithdrawAddressParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *SetWithdrawAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.WithdrawRequestBody); err != nil {
		return err
	}

	// path param delegatorAddr
	if err := r.SetPathParam("delegatorAddr", o.DelegatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
