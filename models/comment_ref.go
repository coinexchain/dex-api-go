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

// CommentRef comment ref
//
// swagger:model CommentRef
type CommentRef struct {

	// attitudes
	// Required: true
	Attitudes []int64 `json:"attitudes"`

	// id
	// Required: true
	ID *string `json:"id"`

	// reward amount
	// Required: true
	RewardAmount *string `json:"reward_amount"`

	// reward target
	// Required: true
	RewardTarget *string `json:"reward_target"`

	// reward token
	// Required: true
	RewardToken *string `json:"reward_token"`
}

// Validate validates this comment ref
func (m *CommentRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttitudes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewardAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewardTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewardToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommentRef) validateAttitudes(formats strfmt.Registry) error {

	if err := validate.Required("attitudes", "body", m.Attitudes); err != nil {
		return err
	}

	return nil
}

func (m *CommentRef) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *CommentRef) validateRewardAmount(formats strfmt.Registry) error {

	if err := validate.Required("reward_amount", "body", m.RewardAmount); err != nil {
		return err
	}

	return nil
}

func (m *CommentRef) validateRewardTarget(formats strfmt.Registry) error {

	if err := validate.Required("reward_target", "body", m.RewardTarget); err != nil {
		return err
	}

	return nil
}

func (m *CommentRef) validateRewardToken(formats strfmt.Registry) error {

	if err := validate.Required("reward_token", "body", m.RewardToken); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommentRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommentRef) UnmarshalBinary(b []byte) error {
	var res CommentRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
