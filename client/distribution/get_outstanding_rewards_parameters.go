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

// NewGetOutstandingRewardsParams creates a new GetOutstandingRewardsParams object
// with the default values initialized.
func NewGetOutstandingRewardsParams() *GetOutstandingRewardsParams {
	var ()
	return &GetOutstandingRewardsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutstandingRewardsParamsWithTimeout creates a new GetOutstandingRewardsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutstandingRewardsParamsWithTimeout(timeout time.Duration) *GetOutstandingRewardsParams {
	var ()
	return &GetOutstandingRewardsParams{

		timeout: timeout,
	}
}

// NewGetOutstandingRewardsParamsWithContext creates a new GetOutstandingRewardsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutstandingRewardsParamsWithContext(ctx context.Context) *GetOutstandingRewardsParams {
	var ()
	return &GetOutstandingRewardsParams{

		Context: ctx,
	}
}

// NewGetOutstandingRewardsParamsWithHTTPClient creates a new GetOutstandingRewardsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutstandingRewardsParamsWithHTTPClient(client *http.Client) *GetOutstandingRewardsParams {
	var ()
	return &GetOutstandingRewardsParams{
		HTTPClient: client,
	}
}

/*GetOutstandingRewardsParams contains all the parameters to send to the API endpoint
for the get outstanding rewards operation typically these are written to a http.Request
*/
type GetOutstandingRewardsParams struct {

	/*ValidatorAddr
	  Bech32 OperatorAddress of validator

	*/
	ValidatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get outstanding rewards params
func (o *GetOutstandingRewardsParams) WithTimeout(timeout time.Duration) *GetOutstandingRewardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outstanding rewards params
func (o *GetOutstandingRewardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outstanding rewards params
func (o *GetOutstandingRewardsParams) WithContext(ctx context.Context) *GetOutstandingRewardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outstanding rewards params
func (o *GetOutstandingRewardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outstanding rewards params
func (o *GetOutstandingRewardsParams) WithHTTPClient(client *http.Client) *GetOutstandingRewardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outstanding rewards params
func (o *GetOutstandingRewardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidatorAddr adds the validatorAddr to the get outstanding rewards params
func (o *GetOutstandingRewardsParams) WithValidatorAddr(validatorAddr string) *GetOutstandingRewardsParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the get outstanding rewards params
func (o *GetOutstandingRewardsParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutstandingRewardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
