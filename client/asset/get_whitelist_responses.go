// Code generated by go-swagger; DO NOT EDIT.

package asset

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

// GetWhitelistReader is a Reader for the GetWhitelist structure.
type GetWhitelistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWhitelistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWhitelistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetWhitelistNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWhitelistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWhitelistOK creates a GetWhitelistOK with default headers values
func NewGetWhitelistOK() *GetWhitelistOK {
	return &GetWhitelistOK{}
}

/*GetWhitelistOK handles this case with default header values.

whitelist with provided symbol
*/
type GetWhitelistOK struct {
	Payload *GetWhitelistOKBody
}

func (o *GetWhitelistOK) Error() string {
	return fmt.Sprintf("[GET /asset/tokens/{symbol}/forbidden/whitelist][%d] getWhitelistOK  %+v", 200, o.Payload)
}

func (o *GetWhitelistOK) GetPayload() *GetWhitelistOKBody {
	return o.Payload
}

func (o *GetWhitelistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWhitelistOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWhitelistNotFound creates a GetWhitelistNotFound with default headers values
func NewGetWhitelistNotFound() *GetWhitelistNotFound {
	return &GetWhitelistNotFound{}
}

/*GetWhitelistNotFound handles this case with default header values.

Token not found
*/
type GetWhitelistNotFound struct {
}

func (o *GetWhitelistNotFound) Error() string {
	return fmt.Sprintf("[GET /asset/tokens/{symbol}/forbidden/whitelist][%d] getWhitelistNotFound ", 404)
}

func (o *GetWhitelistNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWhitelistInternalServerError creates a GetWhitelistInternalServerError with default headers values
func NewGetWhitelistInternalServerError() *GetWhitelistInternalServerError {
	return &GetWhitelistInternalServerError{}
}

/*GetWhitelistInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetWhitelistInternalServerError struct {
}

func (o *GetWhitelistInternalServerError) Error() string {
	return fmt.Sprintf("[GET /asset/tokens/{symbol}/forbidden/whitelist][%d] getWhitelistInternalServerError ", 500)
}

func (o *GetWhitelistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetWhitelistOKBody get whitelist o k body
swagger:model GetWhitelistOKBody
*/
type GetWhitelistOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result []models.Address `json:"result"`
}

// Validate validates this get whitelist o k body
func (o *GetWhitelistOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWhitelistOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	for i := 0; i < len(o.Result); i++ {

		if err := o.Result[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWhitelistOK" + "." + "result" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWhitelistOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWhitelistOKBody) UnmarshalBinary(b []byte) error {
	var res GetWhitelistOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
