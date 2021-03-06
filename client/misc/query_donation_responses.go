// Code generated by go-swagger; DO NOT EDIT.

package misc

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

// QueryDonationReader is a Reader for the QueryDonation structure.
type QueryDonationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryDonationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryDonationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewQueryDonationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQueryDonationOK creates a QueryDonationOK with default headers values
func NewQueryDonationOK() *QueryDonationOK {
	return &QueryDonationOK{}
}

/*QueryDonationOK handles this case with default header values.

OK
*/
type QueryDonationOK struct {
	Payload *QueryDonationOKBody
}

func (o *QueryDonationOK) Error() string {
	return fmt.Sprintf("[GET /misc/donations][%d] queryDonationOK  %+v", 200, o.Payload)
}

func (o *QueryDonationOK) GetPayload() *QueryDonationOKBody {
	return o.Payload
}

func (o *QueryDonationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(QueryDonationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDonationInternalServerError creates a QueryDonationInternalServerError with default headers values
func NewQueryDonationInternalServerError() *QueryDonationInternalServerError {
	return &QueryDonationInternalServerError{}
}

/*QueryDonationInternalServerError handles this case with default header values.

Server internal error
*/
type QueryDonationInternalServerError struct {
}

func (o *QueryDonationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /misc/donations][%d] queryDonationInternalServerError ", 500)
}

func (o *QueryDonationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*QueryDonationOKBody query donation o k body
swagger:model QueryDonationOKBody
*/
type QueryDonationOKBody struct {

	// data
	Data []*models.Donation `json:"data"`

	// timesid
	Timesid []int64 `json:"timesid"`
}

// Validate validates this query donation o k body
func (o *QueryDonationOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QueryDonationOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queryDonationOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *QueryDonationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QueryDonationOKBody) UnmarshalBinary(b []byte) error {
	var res QueryDonationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
