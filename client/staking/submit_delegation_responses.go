// Code generated by go-swagger; DO NOT EDIT.

package staking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/coinexchain/dex-api-go/models"
)

// SubmitDelegationReader is a Reader for the SubmitDelegation structure.
type SubmitDelegationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitDelegationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubmitDelegationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSubmitDelegationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSubmitDelegationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubmitDelegationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubmitDelegationOK creates a SubmitDelegationOK with default headers values
func NewSubmitDelegationOK() *SubmitDelegationOK {
	return &SubmitDelegationOK{}
}

/*SubmitDelegationOK handles this case with default header values.

OK
*/
type SubmitDelegationOK struct {
	Payload *models.BroadcastTxCommitResult
}

func (o *SubmitDelegationOK) Error() string {
	return fmt.Sprintf("[POST /staking/delegators/{delegatorAddr}/delegations][%d] submitDelegationOK  %+v", 200, o.Payload)
}

func (o *SubmitDelegationOK) GetPayload() *models.BroadcastTxCommitResult {
	return o.Payload
}

func (o *SubmitDelegationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BroadcastTxCommitResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitDelegationBadRequest creates a SubmitDelegationBadRequest with default headers values
func NewSubmitDelegationBadRequest() *SubmitDelegationBadRequest {
	return &SubmitDelegationBadRequest{}
}

/*SubmitDelegationBadRequest handles this case with default header values.

Invalid delegator address or delegation request body
*/
type SubmitDelegationBadRequest struct {
}

func (o *SubmitDelegationBadRequest) Error() string {
	return fmt.Sprintf("[POST /staking/delegators/{delegatorAddr}/delegations][%d] submitDelegationBadRequest ", 400)
}

func (o *SubmitDelegationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitDelegationUnauthorized creates a SubmitDelegationUnauthorized with default headers values
func NewSubmitDelegationUnauthorized() *SubmitDelegationUnauthorized {
	return &SubmitDelegationUnauthorized{}
}

/*SubmitDelegationUnauthorized handles this case with default header values.

Key password is wrong
*/
type SubmitDelegationUnauthorized struct {
}

func (o *SubmitDelegationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /staking/delegators/{delegatorAddr}/delegations][%d] submitDelegationUnauthorized ", 401)
}

func (o *SubmitDelegationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitDelegationInternalServerError creates a SubmitDelegationInternalServerError with default headers values
func NewSubmitDelegationInternalServerError() *SubmitDelegationInternalServerError {
	return &SubmitDelegationInternalServerError{}
}

/*SubmitDelegationInternalServerError handles this case with default header values.

Internal Server Error
*/
type SubmitDelegationInternalServerError struct {
}

func (o *SubmitDelegationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /staking/delegators/{delegatorAddr}/delegations][%d] submitDelegationInternalServerError ", 500)
}

func (o *SubmitDelegationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*SubmitDelegationBody submit delegation body
swagger:model SubmitDelegationBody
*/
type SubmitDelegationBody struct {

	// amount
	Amount *models.Coin `json:"amount,omitempty"`

	// base req
	BaseReq *models.BaseReq `json:"base_req,omitempty"`

	// delegator address
	DelegatorAddress models.Address `json:"delegator_address,omitempty"`

	// validator address
	ValidatorAddress models.ValidatorAddress `json:"validator_address,omitempty"`
}

// Validate validates this submit delegation body
func (o *SubmitDelegationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDelegatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValidatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SubmitDelegationBody) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(o.Amount) { // not required
		return nil
	}

	if o.Amount != nil {
		if err := o.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delegation" + "." + "amount")
			}
			return err
		}
	}

	return nil
}

func (o *SubmitDelegationBody) validateBaseReq(formats strfmt.Registry) error {

	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delegation" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *SubmitDelegationBody) validateDelegatorAddress(formats strfmt.Registry) error {

	if swag.IsZero(o.DelegatorAddress) { // not required
		return nil
	}

	if err := o.DelegatorAddress.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("delegation" + "." + "delegator_address")
		}
		return err
	}

	return nil
}

func (o *SubmitDelegationBody) validateValidatorAddress(formats strfmt.Registry) error {

	if swag.IsZero(o.ValidatorAddress) { // not required
		return nil
	}

	if err := o.ValidatorAddress.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("delegation" + "." + "validator_address")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SubmitDelegationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubmitDelegationBody) UnmarshalBinary(b []byte) error {
	var res SubmitDelegationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
