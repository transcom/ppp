// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupContact backup contact
//
// swagger:model BackupContact
type BackupContact struct {

	// email
	// Pattern: ^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
	Email *string `json:"email,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// phone
	// Pattern: ^[2-9]\d{2}-\d{3}-\d{4}$
	Phone *string `json:"phone,omitempty"`
}

// Validate validates this backup contact
func (m *BackupContact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupContact) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.Pattern("email", "body", string(*m.Email), `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`); err != nil {
		return err
	}

	return nil
}

func (m *BackupContact) validatePhone(formats strfmt.Registry) error {

	if swag.IsZero(m.Phone) { // not required
		return nil
	}

	if err := validate.Pattern("phone", "body", string(*m.Phone), `^[2-9]\d{2}-\d{3}-\d{4}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupContact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupContact) UnmarshalBinary(b []byte) error {
	var res BackupContact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
