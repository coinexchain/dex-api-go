// Code generated by go-swagger; DO NOT EDIT.

package bancor

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
	"github.com/go-openapi/swag"
)

// NewQueryBancorDealParams creates a new QueryBancorDealParams object
// with the default values initialized.
func NewQueryBancorDealParams() *QueryBancorDealParams {
	var ()
	return &QueryBancorDealParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryBancorDealParamsWithTimeout creates a new QueryBancorDealParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryBancorDealParamsWithTimeout(timeout time.Duration) *QueryBancorDealParams {
	var ()
	return &QueryBancorDealParams{

		timeout: timeout,
	}
}

// NewQueryBancorDealParamsWithContext creates a new QueryBancorDealParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryBancorDealParamsWithContext(ctx context.Context) *QueryBancorDealParams {
	var ()
	return &QueryBancorDealParams{

		Context: ctx,
	}
}

// NewQueryBancorDealParamsWithHTTPClient creates a new QueryBancorDealParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryBancorDealParamsWithHTTPClient(client *http.Client) *QueryBancorDealParams {
	var ()
	return &QueryBancorDealParams{
		HTTPClient: client,
	}
}

/*QueryBancorDealParams contains all the parameters to send to the API endpoint
for the query bancor deal operation typically these are written to a http.Request
*/
type QueryBancorDealParams struct {

	/*Count
	  Querier deal count limited to 1024

	*/
	Count int32
	/*Market
	  B:stock/money

	*/
	Market string
	/*Sid
	  Sequence id

	*/
	Sid int64
	/*Time
	  Unix timestamp

	*/
	Time int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query bancor deal params
func (o *QueryBancorDealParams) WithTimeout(timeout time.Duration) *QueryBancorDealParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query bancor deal params
func (o *QueryBancorDealParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query bancor deal params
func (o *QueryBancorDealParams) WithContext(ctx context.Context) *QueryBancorDealParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query bancor deal params
func (o *QueryBancorDealParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query bancor deal params
func (o *QueryBancorDealParams) WithHTTPClient(client *http.Client) *QueryBancorDealParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query bancor deal params
func (o *QueryBancorDealParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the query bancor deal params
func (o *QueryBancorDealParams) WithCount(count int32) *QueryBancorDealParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the query bancor deal params
func (o *QueryBancorDealParams) SetCount(count int32) {
	o.Count = count
}

// WithMarket adds the market to the query bancor deal params
func (o *QueryBancorDealParams) WithMarket(market string) *QueryBancorDealParams {
	o.SetMarket(market)
	return o
}

// SetMarket adds the market to the query bancor deal params
func (o *QueryBancorDealParams) SetMarket(market string) {
	o.Market = market
}

// WithSid adds the sid to the query bancor deal params
func (o *QueryBancorDealParams) WithSid(sid int64) *QueryBancorDealParams {
	o.SetSid(sid)
	return o
}

// SetSid adds the sid to the query bancor deal params
func (o *QueryBancorDealParams) SetSid(sid int64) {
	o.Sid = sid
}

// WithTime adds the time to the query bancor deal params
func (o *QueryBancorDealParams) WithTime(time int64) *QueryBancorDealParams {
	o.SetTime(time)
	return o
}

// SetTime adds the time to the query bancor deal params
func (o *QueryBancorDealParams) SetTime(time int64) {
	o.Time = time
}

// WriteToRequest writes these params to a swagger request
func (o *QueryBancorDealParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param count
	qrCount := o.Count
	qCount := swag.FormatInt32(qrCount)
	if qCount != "" {
		if err := r.SetQueryParam("count", qCount); err != nil {
			return err
		}
	}

	// query param market
	qrMarket := o.Market
	qMarket := qrMarket
	if qMarket != "" {
		if err := r.SetQueryParam("market", qMarket); err != nil {
			return err
		}
	}

	// query param sid
	qrSid := o.Sid
	qSid := swag.FormatInt64(qrSid)
	if qSid != "" {
		if err := r.SetQueryParam("sid", qSid); err != nil {
			return err
		}
	}

	// query param time
	qrTime := o.Time
	qTime := swag.FormatInt64(qrTime)
	if qTime != "" {
		if err := r.SetQueryParam("time", qTime); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
