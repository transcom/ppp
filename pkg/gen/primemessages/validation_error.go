// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ValidationError validation error
// swagger:model ValidationError
type ValidationError struct {
	ClientError

	ValidationErrorAllOf1

	// invalid fields
	// Required: true
	InvalidFields map[string]string `json:"invalidFields"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ValidationError) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ClientError
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ClientError = aO0

	// AO1
	var aO1 ValidationErrorAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ValidationErrorAllOf1 = aO1

	// now for regular properties
	var propsValidationError struct {
		InvalidFields map[string]string `json:"invalidFields"`
	}
	if err := swag.ReadJSON(raw, &propsValidationError); err != nil {
		return err
	}
	m.InvalidFields = propsValidationError.InvalidFields

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ValidationError) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ClientError)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ValidationErrorAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	// now for regular properties
	var propsValidationError struct {
		InvalidFields map[string]string `json:"invalidFields"`
	}
	propsValidationError.InvalidFields = m.InvalidFields

	jsonDataPropsValidationError, errValidationError := swag.WriteJSON(propsValidationError)
	if errValidationError != nil {
		return nil, errValidationError
	}
	_parts = append(_parts, jsonDataPropsValidationError)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this validation error
func (m *ValidationError) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ClientError
	if err := m.ClientError.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ValidationErrorAllOf1

	if err := m.validateInvalidFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ValidationError) validateInvalidFields(formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *ValidationError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ValidationError) UnmarshalBinary(b []byte) error {
	var res ValidationError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ValidationErrorAllOf1 validation error all of1
// swagger:model ValidationErrorAllOf1
type ValidationErrorAllOf1 interface{}
