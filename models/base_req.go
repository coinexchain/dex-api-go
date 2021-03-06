// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BaseReq base req
//
// swagger:model BaseReq
type BaseReq struct {

	// account number
	AccountNumber string `json:"account_number,omitempty"`

	// chain id
	ChainID string `json:"chain_id,omitempty"`

	// fees
	Fees []*Coin `json:"fees"`

	// Sender address or Keybase name to generate a transaction
	From string `json:"from,omitempty"`

	// gas
	Gas string `json:"gas,omitempty"`

	// gas adjustment
	GasAdjustment string `json:"gas_adjustment,omitempty"`

	// memo
	Memo string `json:"memo,omitempty"`

	// sequence
	Sequence string `json:"sequence,omitempty"`

	// Estimate gas for a transaction (cannot be used in conjunction with generate_only)
	Simulate bool `json:"simulate,omitempty"`
}

// Validate validates this base req
func (m *BaseReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseReq) validateFees(formats strfmt.Registry) error {

	if swag.IsZero(m.Fees) { // not required
		return nil
	}

	for i := 0; i < len(m.Fees); i++ {
		if swag.IsZero(m.Fees[i]) { // not required
			continue
		}

		if m.Fees[i] != nil {
			if err := m.Fees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BaseReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseReq) UnmarshalBinary(b []byte) error {
	var res BaseReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
