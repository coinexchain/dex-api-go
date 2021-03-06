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

// NewGetUndelegationsOfValidatorParams creates a new GetUndelegationsOfValidatorParams object
// with the default values initialized.
func NewGetUndelegationsOfValidatorParams() *GetUndelegationsOfValidatorParams {
	var ()
	return &GetUndelegationsOfValidatorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUndelegationsOfValidatorParamsWithTimeout creates a new GetUndelegationsOfValidatorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUndelegationsOfValidatorParamsWithTimeout(timeout time.Duration) *GetUndelegationsOfValidatorParams {
	var ()
	return &GetUndelegationsOfValidatorParams{

		timeout: timeout,
	}
}

// NewGetUndelegationsOfValidatorParamsWithContext creates a new GetUndelegationsOfValidatorParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUndelegationsOfValidatorParamsWithContext(ctx context.Context) *GetUndelegationsOfValidatorParams {
	var ()
	return &GetUndelegationsOfValidatorParams{

		Context: ctx,
	}
}

// NewGetUndelegationsOfValidatorParamsWithHTTPClient creates a new GetUndelegationsOfValidatorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUndelegationsOfValidatorParamsWithHTTPClient(client *http.Client) *GetUndelegationsOfValidatorParams {
	var ()
	return &GetUndelegationsOfValidatorParams{
		HTTPClient: client,
	}
}

/*GetUndelegationsOfValidatorParams contains all the parameters to send to the API endpoint
for the get undelegations of validator operation typically these are written to a http.Request
*/
type GetUndelegationsOfValidatorParams struct {

	/*ValidatorAddr
	  Bech32 OperatorAddress of validator

	*/
	ValidatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get undelegations of validator params
func (o *GetUndelegationsOfValidatorParams) WithTimeout(timeout time.Duration) *GetUndelegationsOfValidatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get undelegations of validator params
func (o *GetUndelegationsOfValidatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get undelegations of validator params
func (o *GetUndelegationsOfValidatorParams) WithContext(ctx context.Context) *GetUndelegationsOfValidatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get undelegations of validator params
func (o *GetUndelegationsOfValidatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get undelegations of validator params
func (o *GetUndelegationsOfValidatorParams) WithHTTPClient(client *http.Client) *GetUndelegationsOfValidatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get undelegations of validator params
func (o *GetUndelegationsOfValidatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidatorAddr adds the validatorAddr to the get undelegations of validator params
func (o *GetUndelegationsOfValidatorParams) WithValidatorAddr(validatorAddr string) *GetUndelegationsOfValidatorParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the get undelegations of validator params
func (o *GetUndelegationsOfValidatorParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *GetUndelegationsOfValidatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param validatorAddr
	if err := r.SetPathParam("validatorAddr", o.ValidatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
