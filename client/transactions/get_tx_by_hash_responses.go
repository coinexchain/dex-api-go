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

// GetTxByHashReader is a Reader for the GetTxByHash structure.
type GetTxByHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTxByHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTxByHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTxByHashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTxByHashInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTxByHashOK creates a GetTxByHashOK with default headers values
func NewGetTxByHashOK() *GetTxByHashOK {
	return &GetTxByHashOK{}
}

/*GetTxByHashOK handles this case with default header values.

Tx with the provided hash
*/
type GetTxByHashOK struct {
	Payload *models.BroadcastTxCommitResult
}

func (o *GetTxByHashOK) Error() string {
	return fmt.Sprintf("[GET /txs/{hash}][%d] getTxByHashOK  %+v", 200, o.Payload)
}

func (o *GetTxByHashOK) GetPayload() *models.BroadcastTxCommitResult {
	return o.Payload
}

func (o *GetTxByHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BroadcastTxCommitResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTxByHashNotFound creates a GetTxByHashNotFound with default headers values
func NewGetTxByHashNotFound() *GetTxByHashNotFound {
	return &GetTxByHashNotFound{}
}

/*GetTxByHashNotFound handles this case with default header values.

Tx not found
*/
type GetTxByHashNotFound struct {
}

func (o *GetTxByHashNotFound) Error() string {
	return fmt.Sprintf("[GET /txs/{hash}][%d] getTxByHashNotFound ", 404)
}

func (o *GetTxByHashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTxByHashInternalServerError creates a GetTxByHashInternalServerError with default headers values
func NewGetTxByHashInternalServerError() *GetTxByHashInternalServerError {
	return &GetTxByHashInternalServerError{}
}

/*GetTxByHashInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetTxByHashInternalServerError struct {
}

func (o *GetTxByHashInternalServerError) Error() string {
	return fmt.Sprintf("[GET /txs/{hash}][%d] getTxByHashInternalServerError ", 500)
}

func (o *GetTxByHashInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
