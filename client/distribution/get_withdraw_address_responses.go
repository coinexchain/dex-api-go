// Code generated by go-swagger; DO NOT EDIT.

package distribution

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

// GetWithdrawAddressReader is a Reader for the GetWithdrawAddress structure.
type GetWithdrawAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWithdrawAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWithdrawAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWithdrawAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWithdrawAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWithdrawAddressOK creates a GetWithdrawAddressOK with default headers values
func NewGetWithdrawAddressOK() *GetWithdrawAddressOK {
	return &GetWithdrawAddressOK{}
}

/*GetWithdrawAddressOK handles this case with default header values.

OK
*/
type GetWithdrawAddressOK struct {
	Payload *GetWithdrawAddressOKBody
}

func (o *GetWithdrawAddressOK) Error() string {
	return fmt.Sprintf("[GET /distribution/delegators/{delegatorAddr}/withdraw_address][%d] getWithdrawAddressOK  %+v", 200, o.Payload)
}

func (o *GetWithdrawAddressOK) GetPayload() *GetWithdrawAddressOKBody {
	return o.Payload
}

func (o *GetWithdrawAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWithdrawAddressOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWithdrawAddressBadRequest creates a GetWithdrawAddressBadRequest with default headers values
func NewGetWithdrawAddressBadRequest() *GetWithdrawAddressBadRequest {
	return &GetWithdrawAddressBadRequest{}
}

/*GetWithdrawAddressBadRequest handles this case with default header values.

Invalid delegator address
*/
type GetWithdrawAddressBadRequest struct {
}

func (o *GetWithdrawAddressBadRequest) Error() string {
	return fmt.Sprintf("[GET /distribution/delegators/{delegatorAddr}/withdraw_address][%d] getWithdrawAddressBadRequest ", 400)
}

func (o *GetWithdrawAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWithdrawAddressInternalServerError creates a GetWithdrawAddressInternalServerError with default headers values
func NewGetWithdrawAddressInternalServerError() *GetWithdrawAddressInternalServerError {
	return &GetWithdrawAddressInternalServerError{}
}

/*GetWithdrawAddressInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetWithdrawAddressInternalServerError struct {
}

func (o *GetWithdrawAddressInternalServerError) Error() string {
	return fmt.Sprintf("[GET /distribution/delegators/{delegatorAddr}/withdraw_address][%d] getWithdrawAddressInternalServerError ", 500)
}

func (o *GetWithdrawAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetWithdrawAddressOKBody get withdraw address o k body
swagger:model GetWithdrawAddressOKBody
*/
type GetWithdrawAddressOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result models.Address `json:"result,omitempty"`
}

// Validate validates this get withdraw address o k body
func (o *GetWithdrawAddressOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWithdrawAddressOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if err := o.Result.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getWithdrawAddressOK" + "." + "result")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWithdrawAddressOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWithdrawAddressOKBody) UnmarshalBinary(b []byte) error {
	var res GetWithdrawAddressOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
