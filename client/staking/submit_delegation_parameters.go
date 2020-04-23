// Code generated by go-swagger; DO NOT EDIT.

package staking

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

// NewSubmitDelegationParams creates a new SubmitDelegationParams object
// with the default values initialized.
func NewSubmitDelegationParams() *SubmitDelegationParams {
	var ()
	return &SubmitDelegationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitDelegationParamsWithTimeout creates a new SubmitDelegationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubmitDelegationParamsWithTimeout(timeout time.Duration) *SubmitDelegationParams {
	var ()
	return &SubmitDelegationParams{

		timeout: timeout,
	}
}

// NewSubmitDelegationParamsWithContext creates a new SubmitDelegationParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubmitDelegationParamsWithContext(ctx context.Context) *SubmitDelegationParams {
	var ()
	return &SubmitDelegationParams{

		Context: ctx,
	}
}

// NewSubmitDelegationParamsWithHTTPClient creates a new SubmitDelegationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubmitDelegationParamsWithHTTPClient(client *http.Client) *SubmitDelegationParams {
	var ()
	return &SubmitDelegationParams{
		HTTPClient: client,
	}
}

/*SubmitDelegationParams contains all the parameters to send to the API endpoint
for the submit delegation operation typically these are written to a http.Request
*/
type SubmitDelegationParams struct {

	/*Delegation
	  submit delegation to provided validator

	*/
	Delegation SubmitDelegationBody
	/*DelegatorAddr
	  Bech32 AccAddress of Delegator

	*/
	DelegatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the submit delegation params
func (o *SubmitDelegationParams) WithTimeout(timeout time.Duration) *SubmitDelegationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit delegation params
func (o *SubmitDelegationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit delegation params
func (o *SubmitDelegationParams) WithContext(ctx context.Context) *SubmitDelegationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit delegation params
func (o *SubmitDelegationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit delegation params
func (o *SubmitDelegationParams) WithHTTPClient(client *http.Client) *SubmitDelegationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit delegation params
func (o *SubmitDelegationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegation adds the delegation to the submit delegation params
func (o *SubmitDelegationParams) WithDelegation(delegation SubmitDelegationBody) *SubmitDelegationParams {
	o.SetDelegation(delegation)
	return o
}

// SetDelegation adds the delegation to the submit delegation params
func (o *SubmitDelegationParams) SetDelegation(delegation SubmitDelegationBody) {
	o.Delegation = delegation
}

// WithDelegatorAddr adds the delegatorAddr to the submit delegation params
func (o *SubmitDelegationParams) WithDelegatorAddr(delegatorAddr string) *SubmitDelegationParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the submit delegation params
func (o *SubmitDelegationParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitDelegationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Delegation); err != nil {
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
