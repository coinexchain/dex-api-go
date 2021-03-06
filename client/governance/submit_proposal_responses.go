// Code generated by go-swagger; DO NOT EDIT.

package governance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/coinexchain/dex-api-go/models"
)

// SubmitProposalReader is a Reader for the SubmitProposal structure.
type SubmitProposalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitProposalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubmitProposalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubmitProposalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubmitProposalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubmitProposalOK creates a SubmitProposalOK with default headers values
func NewSubmitProposalOK() *SubmitProposalOK {
	return &SubmitProposalOK{}
}

/*SubmitProposalOK handles this case with default header values.

Tx was succesfully generated
*/
type SubmitProposalOK struct {
	Payload *models.StdTx
}

func (o *SubmitProposalOK) Error() string {
	return fmt.Sprintf("[POST /gov/proposals][%d] submitProposalOK  %+v", 200, o.Payload)
}

func (o *SubmitProposalOK) GetPayload() *models.StdTx {
	return o.Payload
}

func (o *SubmitProposalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StdTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitProposalBadRequest creates a SubmitProposalBadRequest with default headers values
func NewSubmitProposalBadRequest() *SubmitProposalBadRequest {
	return &SubmitProposalBadRequest{}
}

/*SubmitProposalBadRequest handles this case with default header values.

Invalid proposal body
*/
type SubmitProposalBadRequest struct {
}

func (o *SubmitProposalBadRequest) Error() string {
	return fmt.Sprintf("[POST /gov/proposals][%d] submitProposalBadRequest ", 400)
}

func (o *SubmitProposalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitProposalInternalServerError creates a SubmitProposalInternalServerError with default headers values
func NewSubmitProposalInternalServerError() *SubmitProposalInternalServerError {
	return &SubmitProposalInternalServerError{}
}

/*SubmitProposalInternalServerError handles this case with default header values.

Internal Server Error
*/
type SubmitProposalInternalServerError struct {
}

func (o *SubmitProposalInternalServerError) Error() string {
	return fmt.Sprintf("[POST /gov/proposals][%d] submitProposalInternalServerError ", 500)
}

func (o *SubmitProposalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*SubmitProposalBody submit proposal body
swagger:model SubmitProposalBody
*/
type SubmitProposalBody struct {

	// base req
	// Required: true
	BaseReq *models.BaseReq `json:"base_req"`

	// description
	// Required: true
	Description *string `json:"description"`

	// initial deposit
	// Required: true
	InitialDeposit []*models.Coin `json:"initial_deposit"`

	// proposal type
	// Required: true
	ProposalType *string `json:"proposal_type"`

	// proposer
	// Required: true
	Proposer models.Address `json:"proposer"`

	// title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this submit proposal body
func (o *SubmitProposalBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInitialDeposit(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProposalType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProposer(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SubmitProposalBody) validateBaseReq(formats strfmt.Registry) error {

	if err := validate.Required("post_proposal_body"+"."+"base_req", "body", o.BaseReq); err != nil {
		return err
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("post_proposal_body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *SubmitProposalBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("post_proposal_body"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *SubmitProposalBody) validateInitialDeposit(formats strfmt.Registry) error {

	if err := validate.Required("post_proposal_body"+"."+"initial_deposit", "body", o.InitialDeposit); err != nil {
		return err
	}

	for i := 0; i < len(o.InitialDeposit); i++ {
		if swag.IsZero(o.InitialDeposit[i]) { // not required
			continue
		}

		if o.InitialDeposit[i] != nil {
			if err := o.InitialDeposit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_proposal_body" + "." + "initial_deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SubmitProposalBody) validateProposalType(formats strfmt.Registry) error {

	if err := validate.Required("post_proposal_body"+"."+"proposal_type", "body", o.ProposalType); err != nil {
		return err
	}

	return nil
}

func (o *SubmitProposalBody) validateProposer(formats strfmt.Registry) error {

	if err := o.Proposer.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("post_proposal_body" + "." + "proposer")
		}
		return err
	}

	return nil
}

func (o *SubmitProposalBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("post_proposal_body"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SubmitProposalBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubmitProposalBody) UnmarshalBinary(b []byte) error {
	var res SubmitProposalBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
