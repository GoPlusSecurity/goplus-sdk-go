// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Contracts Contracts
//
// swagger:model Contracts
type Contracts struct {

	// It describes the dAap's contract address.
	ContractAddress string `json:"contract_address,omitempty"`

	// It describes the creator address of the contract.
	CreatorAddress string `json:"creator_address,omitempty"`

	// It describes the deployed time of the contract.The value is presented as a timestamp.
	// Example: "deployed_time": 1626578345
	DeploymentTime int64 `json:"deployment_time,omitempty"`

	// It describes whether this contract is open source.
	// "1" means true;
	// "0" means false.
	IsOpenSource int32 `json:"is_open_source,omitempty"`

	// It describes specific malicious behaviors of the contract.
	MaliciousBehavior []interface{} `json:"malicious_behavior"`

	// It describes whether the address is a suspected malicious contract.
	// "1" means true;
	// "0" means that we have not found malicious behavior of this contract.(Notice:"malicious_contract" return "0" does not mean the address is completely safe. Maybe we just haven't found its malicious behavior.)
	MaliciousContract int32 `json:"malicious_contract,omitempty"`

	// It describes whether the creator is a suspected malicious address.
	// "1" means true;
	// "0" means that we have not found malicious behavior of this address.(Notice:"malicious_creator" return "0" does not mean the address is completely safe. Maybe we just haven't found its malicious behavior.)
	MaliciousCreator int32 `json:"malicious_creator,omitempty"`

	// It describes specific malicious behaviors of the contract creator.
	MaliciousCreatorBehavior []interface{} `json:"malicious_creator_behavior"`
}

// Validate validates this contracts
func (m *Contracts) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Contracts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Contracts) UnmarshalBinary(b []byte) error {
	var res Contracts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
