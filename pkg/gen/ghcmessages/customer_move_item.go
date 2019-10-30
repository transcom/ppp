// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomerMoveItem customer move item
// swagger:model customerMoveItem
type CustomerMoveItem struct {

	// Branch of service / Agency
	BranchOfService string `json:"branch_of_service,omitempty"`

	// confirmation number
	ConfirmationNumber string `json:"confirmation_number,omitempty"`

	// when the access code was created
	// Format: datetime
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// customer id
	// Format: uuid
	CustomerID strfmt.UUID `json:"customer_id,omitempty"`

	// Customer Name
	CustomerName *string `json:"customer_name,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Origin Duty Station Name
	OriginDutyStationName *string `json:"origin_duty_station_name,omitempty"`
}

// Validate validates this customer move item
func (m *CustomerMoveItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerMoveItem) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "datetime", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomerMoveItem) validateCustomerID(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerID) { // not required
		return nil
	}

	if err := validate.FormatOf("customer_id", "body", "uuid", m.CustomerID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomerMoveItem) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerMoveItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerMoveItem) UnmarshalBinary(b []byte) error {
	var res CustomerMoveItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
