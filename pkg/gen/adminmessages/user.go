// Code generated by go-swagger; DO NOT EDIT.

package adminmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// User user
// swagger:model User
type User struct {

	// current admin session Id
	// Required: true
	CurrentAdminSessionID *string `json:"currentAdminSessionId"`

	// current mil session Id
	// Required: true
	CurrentMilSessionID *string `json:"currentMilSessionId"`

	// current office session Id
	// Required: true
	CurrentOfficeSessionID *string `json:"currentOfficeSessionId"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// login gov email
	// Required: true
	// Pattern: ^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
	LoginGovEmail *string `json:"loginGovEmail"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentAdminSessionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentMilSessionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentOfficeSessionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoginGovEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateCurrentAdminSessionID(formats strfmt.Registry) error {

	if err := validate.Required("currentAdminSessionId", "body", m.CurrentAdminSessionID); err != nil {
		return err
	}

	return nil
}

func (m *User) validateCurrentMilSessionID(formats strfmt.Registry) error {

	if err := validate.Required("currentMilSessionId", "body", m.CurrentMilSessionID); err != nil {
		return err
	}

	return nil
}

func (m *User) validateCurrentOfficeSessionID(formats strfmt.Registry) error {

	if err := validate.Required("currentOfficeSessionId", "body", m.CurrentOfficeSessionID); err != nil {
		return err
	}

	return nil
}

func (m *User) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateLoginGovEmail(formats strfmt.Registry) error {

	if err := validate.Required("loginGovEmail", "body", m.LoginGovEmail); err != nil {
		return err
	}

	if err := validate.Pattern("loginGovEmail", "body", string(*m.LoginGovEmail), `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
