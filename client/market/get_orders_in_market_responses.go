// Code generated by go-swagger; DO NOT EDIT.

package market

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

// GetOrdersInMarketReader is a Reader for the GetOrdersInMarket structure.
type GetOrdersInMarketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrdersInMarketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrdersInMarketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrdersInMarketBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrdersInMarketInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrdersInMarketOK creates a GetOrdersInMarketOK with default headers values
func NewGetOrdersInMarketOK() *GetOrdersInMarketOK {
	return &GetOrdersInMarketOK{}
}

/*GetOrdersInMarketOK handles this case with default header values.

orders in the specified trading-pair
*/
type GetOrdersInMarketOK struct {
	Payload *GetOrdersInMarketOKBody
}

func (o *GetOrdersInMarketOK) Error() string {
	return fmt.Sprintf("[GET /market/orderbook/{stock}/{money}][%d] getOrdersInMarketOK  %+v", 200, o.Payload)
}

func (o *GetOrdersInMarketOK) GetPayload() *GetOrdersInMarketOKBody {
	return o.Payload
}

func (o *GetOrdersInMarketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrdersInMarketOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrdersInMarketBadRequest creates a GetOrdersInMarketBadRequest with default headers values
func NewGetOrdersInMarketBadRequest() *GetOrdersInMarketBadRequest {
	return &GetOrdersInMarketBadRequest{}
}

/*GetOrdersInMarketBadRequest handles this case with default header values.

Invalid symbol
*/
type GetOrdersInMarketBadRequest struct {
}

func (o *GetOrdersInMarketBadRequest) Error() string {
	return fmt.Sprintf("[GET /market/orderbook/{stock}/{money}][%d] getOrdersInMarketBadRequest ", 400)
}

func (o *GetOrdersInMarketBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrdersInMarketInternalServerError creates a GetOrdersInMarketInternalServerError with default headers values
func NewGetOrdersInMarketInternalServerError() *GetOrdersInMarketInternalServerError {
	return &GetOrdersInMarketInternalServerError{}
}

/*GetOrdersInMarketInternalServerError handles this case with default header values.

Server internal error
*/
type GetOrdersInMarketInternalServerError struct {
}

func (o *GetOrdersInMarketInternalServerError) Error() string {
	return fmt.Sprintf("[GET /market/orderbook/{stock}/{money}][%d] getOrdersInMarketInternalServerError ", 500)
}

func (o *GetOrdersInMarketInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetOrdersInMarketOKBody get orders in market o k body
swagger:model GetOrdersInMarketOKBody
*/
type GetOrdersInMarketOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result []*models.OrderInfo `json:"result"`
}

// Validate validates this get orders in market o k body
func (o *GetOrdersInMarketOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrdersInMarketOKBody) validateResult(formats strfmt.Registry) error {

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
					return ve.ValidateName("getOrdersInMarketOK" + "." + "result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrdersInMarketOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrdersInMarketOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrdersInMarketOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
