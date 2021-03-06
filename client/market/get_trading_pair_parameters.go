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

// NewGetTradingPairParams creates a new GetTradingPairParams object
// with the default values initialized.
func NewGetTradingPairParams() *GetTradingPairParams {
	var ()
	return &GetTradingPairParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTradingPairParamsWithTimeout creates a new GetTradingPairParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTradingPairParamsWithTimeout(timeout time.Duration) *GetTradingPairParams {
	var ()
	return &GetTradingPairParams{

		timeout: timeout,
	}
}

// NewGetTradingPairParamsWithContext creates a new GetTradingPairParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTradingPairParamsWithContext(ctx context.Context) *GetTradingPairParams {
	var ()
	return &GetTradingPairParams{

		Context: ctx,
	}
}

// NewGetTradingPairParamsWithHTTPClient creates a new GetTradingPairParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTradingPairParamsWithHTTPClient(client *http.Client) *GetTradingPairParams {
	var ()
	return &GetTradingPairParams{
		HTTPClient: client,
	}
}

/*GetTradingPairParams contains all the parameters to send to the API endpoint
for the get trading pair operation typically these are written to a http.Request
*/
type GetTradingPairParams struct {

	/*Money
	  money symbol

	*/
	Money string
	/*Stock
	  stock symbol

	*/
	Stock string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get trading pair params
func (o *GetTradingPairParams) WithTimeout(timeout time.Duration) *GetTradingPairParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get trading pair params
func (o *GetTradingPairParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get trading pair params
func (o *GetTradingPairParams) WithContext(ctx context.Context) *GetTradingPairParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get trading pair params
func (o *GetTradingPairParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get trading pair params
func (o *GetTradingPairParams) WithHTTPClient(client *http.Client) *GetTradingPairParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get trading pair params
func (o *GetTradingPairParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMoney adds the money to the get trading pair params
func (o *GetTradingPairParams) WithMoney(money string) *GetTradingPairParams {
	o.SetMoney(money)
	return o
}

// SetMoney adds the money to the get trading pair params
func (o *GetTradingPairParams) SetMoney(money string) {
	o.Money = money
}

// WithStock adds the stock to the get trading pair params
func (o *GetTradingPairParams) WithStock(stock string) *GetTradingPairParams {
	o.SetStock(stock)
	return o
}

// SetStock adds the stock to the get trading pair params
func (o *GetTradingPairParams) SetStock(stock string) {
	o.Stock = stock
}

// WriteToRequest writes these params to a swagger request
func (o *GetTradingPairParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param money
	if err := r.SetPathParam("money", o.Money); err != nil {
		return err
	}

	// path param stock
	if err := r.SetPathParam("stock", o.Stock); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
