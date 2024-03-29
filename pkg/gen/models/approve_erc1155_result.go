// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApproveErc1155Result ApproveErc1155Result
//
// swagger:model ApproveErc1155Result
type ApproveErc1155Result struct {

	// address_info
	AddressInfo *ApproveAddressInfo `json:"address_info,omitempty"`

	// Spender Address
	ApprovedContract string `json:"approved_contract,omitempty"`

	// Latest approval time
	ApprovedTime int64 `json:"approved_time,omitempty"`

	// Latest approval hash
	Hash string `json:"hash,omitempty"`

	// Initial approval hash
	InitialApprovalHash string `json:"initial_approval_hash,omitempty"`

	// Initial approval time
	InitialApprovalTime int64 `json:"initial_approval_time,omitempty"`
}

// Validate validates this approve erc1155 result
func (m *ApproveErc1155Result) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApproveErc1155Result) validateAddressInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressInfo) { // not required
		return nil
	}

	if m.AddressInfo != nil {
		if err := m.AddressInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApproveErc1155Result) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApproveErc1155Result) UnmarshalBinary(b []byte) error {
	var res ApproveErc1155Result
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
