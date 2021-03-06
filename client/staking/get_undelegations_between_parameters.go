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

// NewGetUndelegationsBetweenParams creates a new GetUndelegationsBetweenParams object
// with the default values initialized.
func NewGetUndelegationsBetweenParams() *GetUndelegationsBetweenParams {
	var ()
	return &GetUndelegationsBetweenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUndelegationsBetweenParamsWithTimeout creates a new GetUndelegationsBetweenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUndelegationsBetweenParamsWithTimeout(timeout time.Duration) *GetUndelegationsBetweenParams {
	var ()
	return &GetUndelegationsBetweenParams{

		timeout: timeout,
	}
}

// NewGetUndelegationsBetweenParamsWithContext creates a new GetUndelegationsBetweenParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUndelegationsBetweenParamsWithContext(ctx context.Context) *GetUndelegationsBetweenParams {
	var ()
	return &GetUndelegationsBetweenParams{

		Context: ctx,
	}
}

// NewGetUndelegationsBetweenParamsWithHTTPClient creates a new GetUndelegationsBetweenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUndelegationsBetweenParamsWithHTTPClient(client *http.Client) *GetUndelegationsBetweenParams {
	var ()
	return &GetUndelegationsBetweenParams{
		HTTPClient: client,
	}
}

/*GetUndelegationsBetweenParams contains all the parameters to send to the API endpoint
for the get undelegations between operation typically these are written to a http.Request
*/
type GetUndelegationsBetweenParams struct {

	/*DelegatorAddr
	  Bech32 AccAddress of Delegator

	*/
	DelegatorAddr string
	/*ValidatorAddr
	  Bech32 OperatorAddress of validator

	*/
	ValidatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get undelegations between params
func (o *GetUndelegationsBetweenParams) WithTimeout(timeout time.Duration) *GetUndelegationsBetweenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get undelegations between params
func (o *GetUndelegationsBetweenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get undelegations between params
func (o *GetUndelegationsBetweenParams) WithContext(ctx context.Context) *GetUndelegationsBetweenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get undelegations between params
func (o *GetUndelegationsBetweenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get undelegations between params
func (o *GetUndelegationsBetweenParams) WithHTTPClient(client *http.Client) *GetUndelegationsBetweenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get undelegations between params
func (o *GetUndelegationsBetweenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegatorAddr adds the delegatorAddr to the get undelegations between params
func (o *GetUndelegationsBetweenParams) WithDelegatorAddr(delegatorAddr string) *GetUndelegationsBetweenParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the get undelegations between params
func (o *GetUndelegationsBetweenParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WithValidatorAddr adds the validatorAddr to the get undelegations between params
func (o *GetUndelegationsBetweenParams) WithValidatorAddr(validatorAddr string) *GetUndelegationsBetweenParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the get undelegations between params
func (o *GetUndelegationsBetweenParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *GetUndelegationsBetweenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param delegatorAddr
	if err := r.SetPathParam("delegatorAddr", o.DelegatorAddr); err != nil {
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
