// Code generated by go-swagger; DO NOT EDIT.

package asset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetAssetParamsReader is a Reader for the GetAssetParams structure.
type GetAssetParamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssetParamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssetParamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAssetParamsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAssetParamsOK creates a GetAssetParamsOK with default headers values
func NewGetAssetParamsOK() *GetAssetParamsOK {
	return &GetAssetParamsOK{}
}

/*GetAssetParamsOK handles this case with default header values.

OK
*/
type GetAssetParamsOK struct {
	Payload *GetAssetParamsOKBody
}

func (o *GetAssetParamsOK) Error() string {
	return fmt.Sprintf("[GET /asset/parameters][%d] getAssetParamsOK  %+v", 200, o.Payload)
}

func (o *GetAssetParamsOK) GetPayload() *GetAssetParamsOKBody {
	return o.Payload
}

func (o *GetAssetParamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAssetParamsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssetParamsInternalServerError creates a GetAssetParamsInternalServerError with default headers values
func NewGetAssetParamsInternalServerError() *GetAssetParamsInternalServerError {
	return &GetAssetParamsInternalServerError{}
}

/*GetAssetParamsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAssetParamsInternalServerError struct {
}

func (o *GetAssetParamsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /asset/parameters][%d] getAssetParamsInternalServerError ", 500)
}

func (o *GetAssetParamsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetAssetParamsOKBody get asset params o k body
swagger:model GetAssetParamsOKBody
*/
type GetAssetParamsOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result *GetAssetParamsOKBodyResult `json:"result,omitempty"`
}

// Validate validates this get asset params o k body
func (o *GetAssetParamsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAssetParamsOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAssetParamsOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAssetParamsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAssetParamsOKBody) UnmarshalBinary(b []byte) error {
	var res GetAssetParamsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAssetParamsOKBodyResult get asset params o k body result
swagger:model GetAssetParamsOKBodyResult
*/
type GetAssetParamsOKBodyResult struct {

	// issue 3char token fee
	Issue3charTokenFee string `json:"issue_3char_token_fee,omitempty"`

	// issue 4char token fee
	Issue4charTokenFee string `json:"issue_4char_token_fee,omitempty"`

	// issue 5char token fee
	Issue5charTokenFee string `json:"issue_5char_token_fee,omitempty"`

	// issue 6char token fee
	Issue6charTokenFee string `json:"issue_6char_token_fee,omitempty"`

	// issue rare token fee
	IssueRareTokenFee string `json:"issue_rare_token_fee,omitempty"`

	// issue token fee
	IssueTokenFee string `json:"issue_token_fee,omitempty"`
}

// Validate validates this get asset params o k body result
func (o *GetAssetParamsOKBodyResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAssetParamsOKBodyResult) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAssetParamsOKBodyResult) UnmarshalBinary(b []byte) error {
	var res GetAssetParamsOKBodyResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
