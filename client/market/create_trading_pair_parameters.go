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

// NewCreateTradingPairParams creates a new CreateTradingPairParams object
// with the default values initialized.
func NewCreateTradingPairParams() *CreateTradingPairParams {
	var ()
	return &CreateTradingPairParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTradingPairParamsWithTimeout creates a new CreateTradingPairParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTradingPairParamsWithTimeout(timeout time.Duration) *CreateTradingPairParams {
	var ()
	return &CreateTradingPairParams{

		timeout: timeout,
	}
}

// NewCreateTradingPairParamsWithContext creates a new CreateTradingPairParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTradingPairParamsWithContext(ctx context.Context) *CreateTradingPairParams {
	var ()
	return &CreateTradingPairParams{

		Context: ctx,
	}
}

// NewCreateTradingPairParamsWithHTTPClient creates a new CreateTradingPairParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTradingPairParamsWithHTTPClient(client *http.Client) *CreateTradingPairParams {
	var ()
	return &CreateTradingPairParams{
		HTTPClient: client,
	}
}

/*CreateTradingPairParams contains all the parameters to send to the API endpoint
for the create trading pair operation typically these are written to a http.Request
*/
type CreateTradingPairParams struct {

	/*Info
	  Create trading-pair

	*/
	Info CreateTradingPairBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create trading pair params
func (o *CreateTradingPairParams) WithTimeout(timeout time.Duration) *CreateTradingPairParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create trading pair params
func (o *CreateTradingPairParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create trading pair params
func (o *CreateTradingPairParams) WithContext(ctx context.Context) *CreateTradingPairParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create trading pair params
func (o *CreateTradingPairParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create trading pair params
func (o *CreateTradingPairParams) WithHTTPClient(client *http.Client) *CreateTradingPairParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create trading pair params
func (o *CreateTradingPairParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the create trading pair params
func (o *CreateTradingPairParams) WithInfo(info CreateTradingPairBody) *CreateTradingPairParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the create trading pair params
func (o *CreateTradingPairParams) SetInfo(info CreateTradingPairBody) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTradingPairParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}