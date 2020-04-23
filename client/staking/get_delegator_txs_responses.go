// Code generated by go-swagger; DO NOT EDIT.

package staking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/coinexchain/dex-api-go/models"
)

// GetDelegatorTxsReader is a Reader for the GetDelegatorTxs structure.
type GetDelegatorTxsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDelegatorTxsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDelegatorTxsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetDelegatorTxsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDelegatorTxsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDelegatorTxsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDelegatorTxsOK creates a GetDelegatorTxsOK with default headers values
func NewGetDelegatorTxsOK() *GetDelegatorTxsOK {
	return &GetDelegatorTxsOK{}
}

/*GetDelegatorTxsOK handles this case with default header values.

OK
*/
type GetDelegatorTxsOK struct {
	Payload []*models.PaginatedQueryTxs
}

func (o *GetDelegatorTxsOK) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/txs][%d] getDelegatorTxsOK  %+v", 200, o.Payload)
}

func (o *GetDelegatorTxsOK) GetPayload() []*models.PaginatedQueryTxs {
	return o.Payload
}

func (o *GetDelegatorTxsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDelegatorTxsNoContent creates a GetDelegatorTxsNoContent with default headers values
func NewGetDelegatorTxsNoContent() *GetDelegatorTxsNoContent {
	return &GetDelegatorTxsNoContent{}
}

/*GetDelegatorTxsNoContent handles this case with default header values.

No staking transaction about this delegator address
*/
type GetDelegatorTxsNoContent struct {
}

func (o *GetDelegatorTxsNoContent) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/txs][%d] getDelegatorTxsNoContent ", 204)
}

func (o *GetDelegatorTxsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDelegatorTxsBadRequest creates a GetDelegatorTxsBadRequest with default headers values
func NewGetDelegatorTxsBadRequest() *GetDelegatorTxsBadRequest {
	return &GetDelegatorTxsBadRequest{}
}

/*GetDelegatorTxsBadRequest handles this case with default header values.

Invalid delegator address
*/
type GetDelegatorTxsBadRequest struct {
}

func (o *GetDelegatorTxsBadRequest) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/txs][%d] getDelegatorTxsBadRequest ", 400)
}

func (o *GetDelegatorTxsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDelegatorTxsInternalServerError creates a GetDelegatorTxsInternalServerError with default headers values
func NewGetDelegatorTxsInternalServerError() *GetDelegatorTxsInternalServerError {
	return &GetDelegatorTxsInternalServerError{}
}

/*GetDelegatorTxsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDelegatorTxsInternalServerError struct {
}

func (o *GetDelegatorTxsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /staking/delegators/{delegatorAddr}/txs][%d] getDelegatorTxsInternalServerError ", 500)
}

func (o *GetDelegatorTxsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
