// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8 ResponseWrapper«Map«string,string»»-2a740de0-be5a-4eb3-b6e4-a2cf0a2d8bf8
//
// swagger:model ResponseWrapper«Map«string,string»»-2a740de0-be5a-4eb3-b6e4-a2cf0a2d8bf8
type ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8 struct {

	// Code 1：Success
	Code int32 `json:"code,omitempty"`

	// Response message
	Message string `json:"message,omitempty"`

	// result
	Result *ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8Result `json:"result,omitempty"`
}

// Validate validates this response wrapper map string string 2a740de0 be5a 4eb3 b6e4 a2cf0a2d8bf8
func (m *ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8) validateResult(formats strfmt.Registry) error {

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
func (m *ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8Result Response result
//
// swagger:model ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8Result
type ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8Result struct {

	// It means whether the website is a phishing site.
	// "1" means true;
	// "0" means that we have not found malicious behavior of this website.
	PhishingSite int32 `json:"phishing_site,omitempty"`

	// website contract security
	WebsiteContractSecurity []string `json:"website_contract_security"`
}

// Validate validates this response wrapper map string string2a740de0 be5a4eb3 b6e4 a2cf0a2d8bf8 result
func (m *ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8Result) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8Result) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8Result) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperMapStringString2a740de0Be5a4eb3B6e4A2cf0a2d8bf8Result
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}