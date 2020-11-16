// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOAgent m t o agent
//
// swagger:model MTOAgent
type MTOAgent struct {

	// agent type
	AgentType MTOAgentType `json:"agentType,omitempty"`

	// created at
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// e tag
	// Read Only: true
	ETag string `json:"eTag,omitempty"`

	// email
	// Pattern: ^([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})?$
	Email *string `json:"email,omitempty"`

	// first name
	FirstName *string `json:"firstName,omitempty"`

	// id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// last name
	LastName *string `json:"lastName,omitempty"`

	// mto shipment ID
	// Read Only: true
	// Format: uuid
	MtoShipmentID strfmt.UUID `json:"mtoShipmentID,omitempty"`

	// phone
	// Pattern: ^([2-9]\d{2}-\d{3}-\d{4})?$
	Phone *string `json:"phone,omitempty"`

	// updated at
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this m t o agent
func (m *MTOAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtoShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MTOAgent) validateAgentType(formats strfmt.Registry) error {

	if swag.IsZero(m.AgentType) { // not required
		return nil
	}

	if err := m.AgentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("agentType")
		}
		return err
	}

	return nil
}

func (m *MTOAgent) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOAgent) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.Pattern("email", "body", string(*m.Email), `^([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})?$`); err != nil {
		return err
	}

	return nil
}

func (m *MTOAgent) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOAgent) validateMtoShipmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.MtoShipmentID) { // not required
		return nil
	}

	if err := validate.FormatOf("mtoShipmentID", "body", "uuid", m.MtoShipmentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOAgent) validatePhone(formats strfmt.Registry) error {

	if swag.IsZero(m.Phone) { // not required
		return nil
	}

	if err := validate.Pattern("phone", "body", string(*m.Phone), `^([2-9]\d{2}-\d{3}-\d{4})?$`); err != nil {
		return err
	}

	return nil
}

func (m *MTOAgent) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOAgent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOAgent) UnmarshalBinary(b []byte) error {
	var res MTOAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
