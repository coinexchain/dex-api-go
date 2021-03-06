// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UnbondingEntries unbonding entries
//
// swagger:model UnbondingEntries
type UnbondingEntries struct {

	// balance
	// Required: true
	Balance *string `json:"balance"`

	// completion time
	// Required: true
	CompletionTime *string `json:"completion_time"`

	// creation height
	// Required: true
	CreationHeight *string `json:"creation_height"`

	// initial balance
	// Required: true
	InitialBalance *string `json:"initial_balance"`
}

// Validate validates this unbonding entries
func (m *UnbondingEntries) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitialBalance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnbondingEntries) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

func (m *UnbondingEntries) validateCompletionTime(formats strfmt.Registry) error {

	if err := validate.Required("completion_time", "body", m.CompletionTime); err != nil {
		return err
	}

	return nil
}

func (m *UnbondingEntries) validateCreationHeight(formats strfmt.Registry) error {

	if err := validate.Required("creation_height", "body", m.CreationHeight); err != nil {
		return err
	}

	return nil
}

func (m *UnbondingEntries) validateInitialBalance(formats strfmt.Registry) error {

	if err := validate.Required("initial_balance", "body", m.InitialBalance); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnbondingEntries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnbondingEntries) UnmarshalBinary(b []byte) error {
	var res UnbondingEntries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
