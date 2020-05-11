// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateOrderInfo create order info
//
// swagger:model CreateOrderInfo
type CreateOrderInfo struct {

	// Order feature fee, sato.CET as the unit
	FeatureFee int64 `json:"feature_fee,omitempty"`

	// Freeze sato.CET amount
	Freeze int64 `json:"freeze,omitempty"`

	// Order frozen fee, sato.CET as the unit
	FrozenFee int64 `json:"frozen_fee,omitempty"`

	// height
	Height int64 `json:"height,omitempty"`

	// order id
	OrderID string `json:"order_id,omitempty"`

	// Limited value 2
	OrderType int64 `json:"order_type,omitempty"`

	// price
	Price string `json:"price,omitempty"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`

	// sender
	Sender Address `json:"sender,omitempty"`

	// BUY:1/SELL:2
	Side int64 `json:"side,omitempty"`

	// GTE:3/IOC:4
	TimeInForce int32 `json:"time_in_force,omitempty"`

	// trading pair
	TradingPair string `json:"trading_pair,omitempty"`

	// The tx hash
	TxHash string `json:"tx_hash,omitempty"`
}

// Validate validates this create order info
func (m *CreateOrderInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateOrderInfo) validateSender(formats strfmt.Registry) error {

	if swag.IsZero(m.Sender) { // not required
		return nil
	}

	if err := m.Sender.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sender")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateOrderInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateOrderInfo) UnmarshalBinary(b []byte) error {
	var res CreateOrderInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}