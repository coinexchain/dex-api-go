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

// BlockHeader block header
//
// swagger:model BlockHeader
type BlockHeader struct {

	// app hash
	// Required: true
	AppHash Hash `json:"app_hash"`

	// chain id
	// Required: true
	ChainID *string `json:"chain_id"`

	// consensus hash
	// Required: true
	ConsensusHash Hash `json:"consensus_hash"`

	// data hash
	// Required: true
	DataHash Hash `json:"data_hash"`

	// evidence hash
	// Required: true
	EvidenceHash Hash `json:"evidence_hash"`

	// height
	// Required: true
	Height *float64 `json:"height"`

	// last block id
	// Required: true
	LastBlockID *BlockID `json:"last_block_id"`

	// last commit hash
	// Required: true
	LastCommitHash Hash `json:"last_commit_hash"`

	// last results hash
	// Required: true
	LastResultsHash Hash `json:"last_results_hash"`

	// next validators hash
	// Required: true
	NextValidatorsHash Hash `json:"next_validators_hash"`

	// num txs
	// Required: true
	NumTxs *float64 `json:"num_txs"`

	// proposer address
	// Required: true
	ProposerAddress Address `json:"proposer_address"`

	// time
	// Required: true
	Time *string `json:"time"`

	// total txs
	// Required: true
	TotalTxs *float64 `json:"total_txs"`

	// validators hash
	// Required: true
	ValidatorsHash Hash `json:"validators_hash"`

	// version
	// Required: true
	Version *BlockHeaderVersion `json:"version"`
}

// Validate validates this block header
func (m *BlockHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChainID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsensusHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvidenceHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastBlockID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastCommitHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastResultsHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextValidatorsHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumTxs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProposerAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalTxs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidatorsHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockHeader) validateAppHash(formats strfmt.Registry) error {

	if err := m.AppHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("app_hash")
		}
		return err
	}

	return nil
}

func (m *BlockHeader) validateChainID(formats strfmt.Registry) error {

	if err := validate.Required("chain_id", "body", m.ChainID); err != nil {
		return err
	}

	return nil
}

func (m *BlockHeader) validateConsensusHash(formats strfmt.Registry) error {

	if err := m.ConsensusHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("consensus_hash")
		}
		return err
	}

	return nil
}

func (m *BlockHeader) validateDataHash(formats strfmt.Registry) error {

	if err := m.DataHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data_hash")
		}
		return err
	}

	return nil
}

func (m *BlockHeader) validateEvidenceHash(formats strfmt.Registry) error {

	if err := m.EvidenceHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("evidence_hash")
		}
		return err
	}

	return nil
}

func (m *BlockHeader) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

func (m *BlockHeader) validateLastBlockID(formats strfmt.Registry) error {

	if err := validate.Required("last_block_id", "body", m.LastBlockID); err != nil {
		return err
	}

	if m.LastBlockID != nil {
		if err := m.LastBlockID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_block_id")
			}
			return err
		}
	}

	return nil
}

func (m *BlockHeader) validateLastCommitHash(formats strfmt.Registry) error {

	if err := m.LastCommitHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("last_commit_hash")
		}
		return err
	}

	return nil
}

func (m *BlockHeader) validateLastResultsHash(formats strfmt.Registry) error {

	if err := m.LastResultsHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("last_results_hash")
		}
		return err
	}

	return nil
}

func (m *BlockHeader) validateNextValidatorsHash(formats strfmt.Registry) error {

	if err := m.NextValidatorsHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("next_validators_hash")
		}
		return err
	}

	return nil
}

func (m *BlockHeader) validateNumTxs(formats strfmt.Registry) error {

	if err := validate.Required("num_txs", "body", m.NumTxs); err != nil {
		return err
	}

	return nil
}

func (m *BlockHeader) validateProposerAddress(formats strfmt.Registry) error {

	if err := m.ProposerAddress.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("proposer_address")
		}
		return err
	}

	return nil
}

func (m *BlockHeader) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

func (m *BlockHeader) validateTotalTxs(formats strfmt.Registry) error {

	if err := validate.Required("total_txs", "body", m.TotalTxs); err != nil {
		return err
	}

	return nil
}

func (m *BlockHeader) validateValidatorsHash(formats strfmt.Registry) error {

	if err := m.ValidatorsHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("validators_hash")
		}
		return err
	}

	return nil
}

func (m *BlockHeader) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlockHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockHeader) UnmarshalBinary(b []byte) error {
	var res BlockHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockHeaderVersion block header version
//
// swagger:model BlockHeaderVersion
type BlockHeaderVersion struct {

	// app
	App string `json:"app,omitempty"`

	// block
	Block string `json:"block,omitempty"`
}

// Validate validates this block header version
func (m *BlockHeaderVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BlockHeaderVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockHeaderVersion) UnmarshalBinary(b []byte) error {
	var res BlockHeaderVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
