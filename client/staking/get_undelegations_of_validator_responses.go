// Code generated by go-swagger; DO NOT EDIT.

package staking

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

	"github.com/coinexchain/dex-api-go/models"
)

// GetUndelegationsOfValidatorReader is a Reader for the GetUndelegationsOfValidator structure.
type GetUndelegationsOfValidatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUndelegationsOfValidatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUndelegationsOfValidatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUndelegationsOfValidatorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUndelegationsOfValidatorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUndelegationsOfValidatorOK creates a GetUndelegationsOfValidatorOK with default headers values
func NewGetUndelegationsOfValidatorOK() *GetUndelegationsOfValidatorOK {
	return &GetUndelegationsOfValidatorOK{}
}

/*GetUndelegationsOfValidatorOK handles this case with default header values.

OK
*/
type GetUndelegationsOfValidatorOK struct {
	Payload *GetUndelegationsOfValidatorOKBody
}

func (o *GetUndelegationsOfValidatorOK) Error() string {
	return fmt.Sprintf("[GET /staking/validators/{validatorAddr}/unbonding_delegations][%d] getUndelegationsOfValidatorOK  %+v", 200, o.Payload)
}

func (o *GetUndelegationsOfValidatorOK) GetPayload() *GetUndelegationsOfValidatorOKBody {
	return o.Payload
}

func (o *GetUndelegationsOfValidatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUndelegationsOfValidatorOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUndelegationsOfValidatorBadRequest creates a GetUndelegationsOfValidatorBadRequest with default headers values
func NewGetUndelegationsOfValidatorBadRequest() *GetUndelegationsOfValidatorBadRequest {
	return &GetUndelegationsOfValidatorBadRequest{}
}

/*GetUndelegationsOfValidatorBadRequest handles this case with default header values.

Invalid validator address
*/
type GetUndelegationsOfValidatorBadRequest struct {
}

func (o *GetUndelegationsOfValidatorBadRequest) Error() string {
	return fmt.Sprintf("[GET /staking/validators/{validatorAddr}/unbonding_delegations][%d] getUndelegationsOfValidatorBadRequest ", 400)
}

func (o *GetUndelegationsOfValidatorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUndelegationsOfValidatorInternalServerError creates a GetUndelegationsOfValidatorInternalServerError with default headers values
func NewGetUndelegationsOfValidatorInternalServerError() *GetUndelegationsOfValidatorInternalServerError {
	return &GetUndelegationsOfValidatorInternalServerError{}
}

/*GetUndelegationsOfValidatorInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetUndelegationsOfValidatorInternalServerError struct {
}

func (o *GetUndelegationsOfValidatorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /staking/validators/{validatorAddr}/unbonding_delegations][%d] getUndelegationsOfValidatorInternalServerError ", 500)
}

func (o *GetUndelegationsOfValidatorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetUndelegationsOfValidatorOKBody get undelegations of validator o k body
swagger:model GetUndelegationsOfValidatorOKBody
*/
type GetUndelegationsOfValidatorOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result []*models.UnbondingDelegationPair `json:"result"`
}

// Validate validates this get undelegations of validator o k body
func (o *GetUndelegationsOfValidatorOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUndelegationsOfValidatorOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	for i := 0; i < len(o.Result); i++ {
		if swag.IsZero(o.Result[i]) { // not required
			continue
		}

		if o.Result[i] != nil {
			if err := o.Result[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUndelegationsOfValidatorOK" + "." + "result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUndelegationsOfValidatorOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUndelegationsOfValidatorOKBody) UnmarshalBinary(b []byte) error {
	var res GetUndelegationsOfValidatorOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
