// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseWrapperListJSONObject ResponseWrapperListJSONObject
//
// swagger:model ResponseWrapperListJSONObject
type ResponseWrapperListJSONObject struct {

	// Code 1: Success
	Code int32 `json:"code,omitempty"`

	// Response message
	Message string `json:"message,omitempty"`

	// Response result
	Result []JSONObject `json:"result"`
}

// Validate validates this response wrapper list JSON object
func (m *ResponseWrapperListJSONObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseWrapperListJSONObject) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	for i := 0; i < len(m.Result); i++ {

		if err := m.Result[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperListJSONObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperListJSONObject) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperListJSONObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
