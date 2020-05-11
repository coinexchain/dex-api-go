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

// Unlock unlock
//
// swagger:model Unlock
type Unlock struct {

	// address
	Address Address `json:"address,omitempty"`

	// coins
	Coins []*Coin `json:"coins"`

	// frozen coins
	FrozenCoins []*Coin `json:"frozen_coins"`

	// height
	Height int64 `json:"height,omitempty"`

	// locked coins
	LockedCoins []*LockedCoin `json:"locked_coins"`

	// unlocked
	Unlocked []*Coin `json:"unlocked"`
}

// Validate validates this unlock
func (m *Unlock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrozenCoins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLockedCoins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnlocked(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Unlock) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := m.Address.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("address")
		}
		return err
	}

	return nil
}

func (m *Unlock) validateCoins(formats strfmt.Registry) error {

	if swag.IsZero(m.Coins) { // not required
		return nil
	}

	for i := 0; i < len(m.Coins); i++ {
		if swag.IsZero(m.Coins[i]) { // not required
			continue
		}

		if m.Coins[i] != nil {
			if err := m.Coins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("coins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Unlock) validateFrozenCoins(formats strfmt.Registry) error {

	if swag.IsZero(m.FrozenCoins) { // not required
		return nil
	}

	for i := 0; i < len(m.FrozenCoins); i++ {
		if swag.IsZero(m.FrozenCoins[i]) { // not required
			continue
		}

		if m.FrozenCoins[i] != nil {
			if err := m.FrozenCoins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("frozen_coins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Unlock) validateLockedCoins(formats strfmt.Registry) error {

	if swag.IsZero(m.LockedCoins) { // not required
		return nil
	}

	for i := 0; i < len(m.LockedCoins); i++ {
		if swag.IsZero(m.LockedCoins[i]) { // not required
			continue
		}

		if m.LockedCoins[i] != nil {
			if err := m.LockedCoins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locked_coins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Unlock) validateUnlocked(formats strfmt.Registry) error {

	if swag.IsZero(m.Unlocked) { // not required
		return nil
	}

	for i := 0; i < len(m.Unlocked); i++ {
		if swag.IsZero(m.Unlocked[i]) { // not required
			continue
		}

		if m.Unlocked[i] != nil {
			if err := m.Unlocked[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unlocked" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Unlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Unlock) UnmarshalBinary(b []byte) error {
	var res Unlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}