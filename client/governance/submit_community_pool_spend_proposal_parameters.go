// Code generated by go-swagger; DO NOT EDIT.

package governance

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

// NewSubmitCommunityPoolSpendProposalParams creates a new SubmitCommunityPoolSpendProposalParams object
// with the default values initialized.
func NewSubmitCommunityPoolSpendProposalParams() *SubmitCommunityPoolSpendProposalParams {
	var ()
	return &SubmitCommunityPoolSpendProposalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitCommunityPoolSpendProposalParamsWithTimeout creates a new SubmitCommunityPoolSpendProposalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubmitCommunityPoolSpendProposalParamsWithTimeout(timeout time.Duration) *SubmitCommunityPoolSpendProposalParams {
	var ()
	return &SubmitCommunityPoolSpendProposalParams{

		timeout: timeout,
	}
}

// NewSubmitCommunityPoolSpendProposalParamsWithContext creates a new SubmitCommunityPoolSpendProposalParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubmitCommunityPoolSpendProposalParamsWithContext(ctx context.Context) *SubmitCommunityPoolSpendProposalParams {
	var ()
	return &SubmitCommunityPoolSpendProposalParams{

		Context: ctx,
	}
}

// NewSubmitCommunityPoolSpendProposalParamsWithHTTPClient creates a new SubmitCommunityPoolSpendProposalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubmitCommunityPoolSpendProposalParamsWithHTTPClient(client *http.Client) *SubmitCommunityPoolSpendProposalParams {
	var ()
	return &SubmitCommunityPoolSpendProposalParams{
		HTTPClient: client,
	}
}

/*SubmitCommunityPoolSpendProposalParams contains all the parameters to send to the API endpoint
for the submit community pool spend proposal operation typically these are written to a http.Request
*/
type SubmitCommunityPoolSpendProposalParams struct {

	/*PostProposalBody
	  The community pool spend proposal body contains coin amount to be spent

	*/
	PostProposalBody SubmitCommunityPoolSpendProposalBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the submit community pool spend proposal params
func (o *SubmitCommunityPoolSpendProposalParams) WithTimeout(timeout time.Duration) *SubmitCommunityPoolSpendProposalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit community pool spend proposal params
func (o *SubmitCommunityPoolSpendProposalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit community pool spend proposal params
func (o *SubmitCommunityPoolSpendProposalParams) WithContext(ctx context.Context) *SubmitCommunityPoolSpendProposalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit community pool spend proposal params
func (o *SubmitCommunityPoolSpendProposalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit community pool spend proposal params
func (o *SubmitCommunityPoolSpendProposalParams) WithHTTPClient(client *http.Client) *SubmitCommunityPoolSpendProposalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit community pool spend proposal params
func (o *SubmitCommunityPoolSpendProposalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostProposalBody adds the postProposalBody to the submit community pool spend proposal params
func (o *SubmitCommunityPoolSpendProposalParams) WithPostProposalBody(postProposalBody SubmitCommunityPoolSpendProposalBody) *SubmitCommunityPoolSpendProposalParams {
	o.SetPostProposalBody(postProposalBody)
	return o
}

// SetPostProposalBody adds the postProposalBody to the submit community pool spend proposal params
func (o *SubmitCommunityPoolSpendProposalParams) SetPostProposalBody(postProposalBody SubmitCommunityPoolSpendProposalBody) {
	o.PostProposalBody = postProposalBody
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitCommunityPoolSpendProposalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.PostProposalBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}