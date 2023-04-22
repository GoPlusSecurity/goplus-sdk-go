// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContractApproveResponse ContractApproveResponse
//
// swagger:model ContractApproveResponse
type ContractApproveResponse struct {

	// It describes the approved contract name.
	ContractName string `json:"contract_name,omitempty"`

	// It describes the creator address of the contract.(Notice:When the address is not a contract ("is_contract"=0), it will return "null".)
	CreatorAddress string `json:"creator_address,omitempty"`

	// It describes the deployed time of the contract.
	// The value is presented as a timestamp.
	// Example: "deployed_time": 1626578345(Notice:When the address is not a contract ("is_contract"=0), it will return "null".)
	DeployedTime int64 `json:"deployed_time,omitempty"`

	// It describes whether the address is a suspected malicious contract.
	// "1" means true;
	// "0" means that we have not found malicious behavior of this address.(Notice:Return "0" does not mean it is safe. Maybe we just haven't found its malicious behavior.)
	DoubtList int32 `json:"doubt_list,omitempty"`

	// It describes whether the address is a contract.
	// "1" means true;
	// "0" means false.
	IsContract int32 `json:"is_contract,omitempty"`

	// It describes whether this contract is open source.
	// "1" means true;
	// "0" means false.(Notice:When the address is not a contract ("is_contract"=0), it will return "null".)
	IsOpenSource int32 `json:"is_open_source,omitempty"`

	// Whether the spender is a proxy contract.
	IsProxy int32 `json:"is_proxy,omitempty"`

	// It describes specific malicious behaviors.
	// "honeypot_related_address" means that the address is related to honeypot tokens or has created scam tokens.
	// "phishing_activities" means that this address has implemented phishing activities.
	// "blackmail_activities" means that this address has implemented blackmail activities.
	// "stealing_attack" means that this address has implemented stealing attacks.
	// "fake_kyc" means that this address is involved in fake KYC.
	// "malicious_mining_activities" means that this address is involved in malicious mining activities.
	// "darkweb_transactions" means that this address is involved in darkweb transactions.
	// "cybercrime" means that this address is involved in cybercrime.
	// "money_laundering" means that this address is involved in money laundering.
	// "financial_crime" means that this address is involved in financial crime.
	// “blacklist_doubt” means that the address is suspected of malicious behavior and is therefore blacklisted.(Notice:Returning an empty array means that no malicious behavior was found at that address.)
	MaliciousBehavior []string `json:"malicious_behavior"`

	// It describes which dapp uses the contract.
	// Example:"tag": "Compound"
	Tag string `json:"tag,omitempty"`

	// It describes whether the address is a famous and trustworthy one.
	// "1" means true;
	// "0" means that we have not included this address in the trusted list;(Notice:Return "0" does not mean the address is not trustworthy. Maybe we just haven't included it yet.)
	TrustList int32 `json:"trust_list,omitempty"`
}

// Validate validates this contract approve response
func (m *ContractApproveResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContractApproveResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractApproveResponse) UnmarshalBinary(b []byte) error {
	var res ContractApproveResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}