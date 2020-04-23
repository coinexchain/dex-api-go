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

// NewSubmitRedelegationParams creates a new SubmitRedelegationParams object
// with the default values initialized.
func NewSubmitRedelegationParams() *SubmitRedelegationParams {
	var ()
	return &SubmitRedelegationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitRedelegationParamsWithTimeout creates a new SubmitRedelegationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubmitRedelegationParamsWithTimeout(timeout time.Duration) *SubmitRedelegationParams {
	var ()
	return &SubmitRedelegationParams{

		timeout: timeout,
	}
}

// NewSubmitRedelegationParamsWithContext creates a new SubmitRedelegationParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubmitRedelegationParamsWithContext(ctx context.Context) *SubmitRedelegationParams {
	var ()
	return &SubmitRedelegationParams{

		Context: ctx,
	}
}

// NewSubmitRedelegationParamsWithHTTPClient creates a new SubmitRedelegationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubmitRedelegationParamsWithHTTPClient(client *http.Client) *SubmitRedelegationParams {
	var ()
	return &SubmitRedelegationParams{
		HTTPClient: client,
	}
}

/*SubmitRedelegationParams contains all the parameters to send to the API endpoint
for the submit redelegation operation typically these are written to a http.Request
*/
type SubmitRedelegationParams struct {

	/*Delegation
	  The sender and tx information

	*/
	Delegation SubmitRedelegationBody
	/*DelegatorAddr
	  Bech32 AccAddress of Delegator

	*/
	DelegatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the submit redelegation params
func (o *SubmitRedelegationParams) WithTimeout(timeout time.Duration) *SubmitRedelegationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit redelegation params
func (o *SubmitRedelegationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit redelegation params
func (o *SubmitRedelegationParams) WithContext(ctx context.Context) *SubmitRedelegationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit redelegation params
func (o *SubmitRedelegationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit redelegation params
func (o *SubmitRedelegationParams) WithHTTPClient(client *http.Client) *SubmitRedelegationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit redelegation params
func (o *SubmitRedelegationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegation adds the delegation to the submit redelegation params
func (o *SubmitRedelegationParams) WithDelegation(delegation SubmitRedelegationBody) *SubmitRedelegationParams {
	o.SetDelegation(delegation)
	return o
}

// SetDelegation adds the delegation to the submit redelegation params
func (o *SubmitRedelegationParams) SetDelegation(delegation SubmitRedelegationBody) {
	o.Delegation = delegation
}

// WithDelegatorAddr adds the delegatorAddr to the submit redelegation params
func (o *SubmitRedelegationParams) WithDelegatorAddr(delegatorAddr string) *SubmitRedelegationParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the submit redelegation params
func (o *SubmitRedelegationParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitRedelegationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
