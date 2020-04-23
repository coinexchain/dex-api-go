// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/coinexchain/dex-api-go/models"
)

// SearchTxReader is a Reader for the SearchTx structure.
type SearchTxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchTxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchTxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchTxBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchTxInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchTxOK creates a SearchTxOK with default headers values
func NewSearchTxOK() *SearchTxOK {
	return &SearchTxOK{}
}

/*SearchTxOK handles this case with default header values.

All txs matching the provided events
*/
type SearchTxOK struct {
	Payload *models.PaginatedQueryTxs
}

func (o *SearchTxOK) Error() string {
	return fmt.Sprintf("[GET /txs][%d] searchTxOK  %+v", 200, o.Payload)
}

func (o *SearchTxOK) GetPayload() *models.PaginatedQueryTxs {
	return o.Payload
}

func (o *SearchTxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedQueryTxs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchTxBadRequest creates a SearchTxBadRequest with default headers values
func NewSearchTxBadRequest() *SearchTxBadRequest {
	return &SearchTxBadRequest{}
}

/*SearchTxBadRequest handles this case with default header values.

Invalid search events
*/
type SearchTxBadRequest struct {
}

func (o *SearchTxBadRequest) Error() string {
	return fmt.Sprintf("[GET /txs][%d] searchTxBadRequest ", 400)
}

func (o *SearchTxBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSearchTxInternalServerError creates a SearchTxInternalServerError with default headers values
func NewSearchTxInternalServerError() *SearchTxInternalServerError {
	return &SearchTxInternalServerError{}
}

/*SearchTxInternalServerError handles this case with default header values.

Internal Server Error
*/
type SearchTxInternalServerError struct {
}

func (o *SearchTxInternalServerError) Error() string {
	return fmt.Sprintf("[GET /txs][%d] searchTxInternalServerError ", 500)
}

func (o *SearchTxInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}