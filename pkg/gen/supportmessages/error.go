// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Error error
//
// swagger:model Error
type Error struct {

	// detail
	// Required: true
	Detail *string `json:"detail"`

	// instance
	// Format: uuid
	Instance strfmt.UUID `json:"instance,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this error
func (m *Error) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Error) validateDetail(formats strfmt.Registry) error {

	if err := validate.Required("detail", "body", m.Detail); err != nil {
		return err
	}

	return nil
}

func (m *Error) validateInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.Instance) { // not required
		return nil
	}

	if err := validate.FormatOf("instance", "body", "uuid", m.Instance.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Error) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Error) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error) UnmarshalBinary(b []byte) error {
	var res Error
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
