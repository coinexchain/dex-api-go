// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DelegatorTotalRewards delegator total rewards
//
// swagger:model DelegatorTotalRewards
type DelegatorTotalRewards struct {

	// rewards
	// Required: true
	Rewards []*DelegationDelegatorReward `json:"rewards"`

	// total
	// Required: true
	Total []*Coin `json:"total"`
}

// Validate validates this delegator total rewards
func (m *DelegatorTotalRewards) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRewards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DelegatorTotalRewards) validateRewards(formats strfmt.Registry) error {

	if err := validate.Required("rewards", "body", m.Rewards); err != nil {
		return err
	}

	for i := 0; i < len(m.Rewards); i++ {
		if swag.IsZero(m.Rewards[i]) { // not required
			continue
		}

		if m.Rewards[i] != nil {
			if err := m.Rewards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rewards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DelegatorTotalRewards) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	for i := 0; i < len(m.Total); i++ {
		if swag.IsZero(m.Total[i]) { // not required
			continue
		}

		if m.Total[i] != nil {
			if err := m.Total[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("total" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DelegatorTotalRewards) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DelegatorTotalRewards) UnmarshalBinary(b []byte) error {
	var res DelegatorTotalRewards
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
