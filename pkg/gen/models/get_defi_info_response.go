// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetDefiInfoResponse GetDefiInfoResponse
//
// swagger:model GetDefiInfoResponse
type GetDefiInfoResponse struct {

	// Code 1: Success
	Code int32 `json:"code,omitempty"`

	// Response message
	Message string `json:"message,omitempty"`

	// result
	Result *GetDefiInfoResponseResult `json:"result,omitempty"`
}

// Validate validates this get defi info response
func (m *GetDefiInfoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDefiInfoResponse) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetDefiInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDefiInfoResponse) UnmarshalBinary(b []byte) error {
	var res GetDefiInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetDefiInfoResponseResult Response result
//
// swagger:model GetDefiInfoResponseResult
type GetDefiInfoResponseResult struct {

	// It describes whether the owner can spend the allowance that obtained by the contract. If so, this function could potentially be abused to steal user assets.
	// "1" means true;
	// "0" means false;
	// "-1" means unknown.
	ApprovalAbuse int32 `json:"approval_abuse,omitempty"`

	// It describes whether the contract has blacklist function that would block user from withdrawing their assets.
	// "1" means true;
	// "0" means false;
	// "-1" means unknown.
	Blacklist int32 `json:"blacklist,omitempty"`

	// Name of the contract.
	ContractName string `json:"contract_name,omitempty"`

	// It describes whether this contract is open source.
	// "1" means true;
	// "0" means false.
	IsOpenSource int32 `json:"is_open_source,omitempty"`

	// It describes whether this contract has a proxy contract.
	// "1" means true;
	// "0" means false;
	// "-1" means unknown.
	IsProxy int32 `json:"is_proxy,omitempty"`

	// owner
	Owner *GetDefiInfoResponseResultOwner `json:"owner,omitempty"`

	// It descirbes whether the contract owner can withdraw all the assets in the contract, without uses' permission.
	// "1" means true;
	// "0" means false;
	// "-1" means unknown.
	PrivilegeWithdraw int32 `json:"privilege_withdraw,omitempty"`

	// It describes whether this contract can self destruct.
	// "1" means true;
	// "0" means false;
	// "-1" means unknown.
	Selfdestruct int32 `json:"selfdestruct,omitempty"`

	// It describes whether the contract lacks withdrawal method. If it is missing, users will be unable to withdraw the assets they have putted in.
	// "1" means true;
	// "0" means false;
	// "-1" means unknown.
	WithdrawMissing int32 `json:"withdraw_missing,omitempty"`
}

// Validate validates this get defi info response result
func (m *GetDefiInfoResponseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDefiInfoResponseResult) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result" + "." + "owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetDefiInfoResponseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDefiInfoResponseResult) UnmarshalBinary(b []byte) error {
	var res GetDefiInfoResponseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetDefiInfoResponseResultOwner When there is no owner function, or the ownership is unreadable or private, it would return empty.
// "owner": {  }
//
// swagger:model GetDefiInfoResponseResultOwner
type GetDefiInfoResponseResultOwner struct {

	// owner address of the contract.
	// No return means unknown.
	OwnerAddress string `json:"owner_address,omitempty"`

	// the function name of ownership.
	// If there is no return, means unknown.
	OwnerName string `json:"owner_name,omitempty"`

	// blackhole" : the owner is a blackhole address.
	// "contract" : the owner is a contract.
	// "eoa" : the owner is a common address (eoa).
	// "multi-address": the owner is an array/list.
	// null: the address is not detected.
	// No return means unknown.
	OwnerType string `json:"owner_type,omitempty"`
}

// Validate validates this get defi info response result owner
func (m *GetDefiInfoResponseResultOwner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetDefiInfoResponseResultOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDefiInfoResponseResultOwner) UnmarshalBinary(b []byte) error {
	var res GetDefiInfoResponseResultOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
