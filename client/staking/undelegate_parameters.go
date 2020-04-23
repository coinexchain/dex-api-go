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

// NewUndelegateParams creates a new UndelegateParams object
// with the default values initialized.
func NewUndelegateParams() *UndelegateParams {
	var ()
	return &UndelegateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUndelegateParamsWithTimeout creates a new UndelegateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUndelegateParamsWithTimeout(timeout time.Duration) *UndelegateParams {
	var ()
	return &UndelegateParams{

		timeout: timeout,
	}
}

// NewUndelegateParamsWithContext creates a new UndelegateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUndelegateParamsWithContext(ctx context.Context) *UndelegateParams {
	var ()
	return &UndelegateParams{

		Context: ctx,
	}
}

// NewUndelegateParamsWithHTTPClient creates a new UndelegateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUndelegateParamsWithHTTPClient(client *http.Client) *UndelegateParams {
	var ()
	return &UndelegateParams{
		HTTPClient: client,
	}
}

/*UndelegateParams contains all the parameters to send to the API endpoint
for the undelegate operation typically these are written to a http.Request
*/
type UndelegateParams struct {

	/*Delegation
	  The password of the account to remove from the KMS

	*/
	Delegation UndelegateBody
	/*DelegatorAddr
	  Bech32 AccAddress of Delegator

	*/
	DelegatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the undelegate params
func (o *UndelegateParams) WithTimeout(timeout time.Duration) *UndelegateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the undelegate params
func (o *UndelegateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the undelegate params
func (o *UndelegateParams) WithContext(ctx context.Context) *UndelegateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the undelegate params
func (o *UndelegateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the undelegate params
func (o *UndelegateParams) WithHTTPClient(client *http.Client) *UndelegateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the undelegate params
func (o *UndelegateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegation adds the delegation to the undelegate params
func (o *UndelegateParams) WithDelegation(delegation UndelegateBody) *UndelegateParams {
	o.SetDelegation(delegation)
	return o
}

// SetDelegation adds the delegation to the undelegate params
func (o *UndelegateParams) SetDelegation(delegation UndelegateBody) {
	o.Delegation = delegation
}

// WithDelegatorAddr adds the delegatorAddr to the undelegate params
func (o *UndelegateParams) WithDelegatorAddr(delegatorAddr string) *UndelegateParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the undelegate params
func (o *UndelegateParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *UndelegateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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