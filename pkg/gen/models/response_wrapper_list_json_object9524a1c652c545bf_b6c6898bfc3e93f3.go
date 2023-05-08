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

// ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3 ResponseWrapper«List«JSONObject»»-9524a1c6-52c5-45bf-b6c6-898bfc3e93f3
//
// swagger:model ResponseWrapper«List«JSONObject»»-9524a1c6-52c5-45bf-b6c6-898bfc3e93f3
type ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3 struct {

	// Code 1：Success
	Code int32 `json:"code,omitempty"`

	// Response message
	Message string `json:"message,omitempty"`

	// Response result
	Result []*ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3ResultItems0 `json:"result"`
}

// Validate validates this response wrapper list JSON object 9524a1c6 52c5 45bf b6c6 898bfc3e93f3
func (m *ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	for i := 0; i < len(m.Result); i++ {
		if swag.IsZero(m.Result[i]) { // not required
			continue
		}

		if m.Result[i] != nil {
			if err := m.Result[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3ResultItems0 response wrapper list JSON object9524a1c652c545bf b6c6898bfc3e93f3 result items0
//
// swagger:model ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3ResultItems0
type ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3ResultItems0 struct {

	// chain id
	ID string `json:"id,omitempty"`

	// chain name
	Name string `json:"name,omitempty"`
}

// Validate validates this response wrapper list JSON object9524a1c652c545bf b6c6898bfc3e93f3 result items0
func (m *ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3ResultItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3ResultItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3ResultItems0) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperListJSONObject9524a1c652c545bfB6c6898bfc3e93f3ResultItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}