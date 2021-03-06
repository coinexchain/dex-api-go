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
	"github.com/go-openapi/swag"
)

// NewGetValidatorsParams creates a new GetValidatorsParams object
// with the default values initialized.
func NewGetValidatorsParams() *GetValidatorsParams {
	var ()
	return &GetValidatorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetValidatorsParamsWithTimeout creates a new GetValidatorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetValidatorsParamsWithTimeout(timeout time.Duration) *GetValidatorsParams {
	var ()
	return &GetValidatorsParams{

		timeout: timeout,
	}
}

// NewGetValidatorsParamsWithContext creates a new GetValidatorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetValidatorsParamsWithContext(ctx context.Context) *GetValidatorsParams {
	var ()
	return &GetValidatorsParams{

		Context: ctx,
	}
}

// NewGetValidatorsParamsWithHTTPClient creates a new GetValidatorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetValidatorsParamsWithHTTPClient(client *http.Client) *GetValidatorsParams {
	var ()
	return &GetValidatorsParams{
		HTTPClient: client,
	}
}

/*GetValidatorsParams contains all the parameters to send to the API endpoint
for the get validators operation typically these are written to a http.Request
*/
type GetValidatorsParams struct {

	/*Limit
	  The maximum number of items per page.

	*/
	Limit *int64
	/*Page
	  The page number.

	*/
	Page *int64
	/*Status
	  The validator bond status. Must be either 'bonded', 'unbonded', or 'unbonding'.

	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get validators params
func (o *GetValidatorsParams) WithTimeout(timeout time.Duration) *GetValidatorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get validators params
func (o *GetValidatorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get validators params
func (o *GetValidatorsParams) WithContext(ctx context.Context) *GetValidatorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get validators params
func (o *GetValidatorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get validators params
func (o *GetValidatorsParams) WithHTTPClient(client *http.Client) *GetValidatorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get validators params
func (o *GetValidatorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get validators params
func (o *GetValidatorsParams) WithLimit(limit *int64) *GetValidatorsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get validators params
func (o *GetValidatorsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get validators params
func (o *GetValidatorsParams) WithPage(page *int64) *GetValidatorsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get validators params
func (o *GetValidatorsParams) SetPage(page *int64) {
	o.Page = page
}

// WithStatus adds the status to the get validators params
func (o *GetValidatorsParams) WithStatus(status *string) *GetValidatorsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get validators params
func (o *GetValidatorsParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetValidatorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
