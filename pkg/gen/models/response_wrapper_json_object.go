// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseWrapperJSONObject ResponseWrapperJSONObject
//
// swagger:model ResponseWrapperJSONObject
type ResponseWrapperJSONObject struct {

	// Code 1: Success
	Code int32 `json:"code,omitempty"`

	// Response message
	Message string `json:"message,omitempty"`

	// Response result
	Result map[string]interface{} `json:"result,omitempty"`
}

// Validate validates this response wrapper JSON object
func (m *ResponseWrapperJSONObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperJSONObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperJSONObject) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperJSONObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
