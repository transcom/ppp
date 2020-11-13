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

// QueuePaymentRequest queue payment request
//
// swagger:model QueuePaymentRequest
type QueuePaymentRequest struct {

	// Days since the payment request has been requested.  Decimal representation will allow more accurate sorting.
	Age int64 `json:"age,omitempty"`

	// customer
	Customer *Customer `json:"customer,omitempty"`

	// department indicator
	DepartmentIndicator DeptIndicator `json:"departmentIndicator,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// locator
	Locator string `json:"locator,omitempty"`

	// move ID
	// Format: uuid
	MoveID strfmt.UUID `json:"moveID,omitempty"`

	// origin g b l o c
	OriginGBLOC GBLOC `json:"originGBLOC,omitempty"`

	// status
	Status PaymentRequestStatus `json:"status,omitempty"`

	// submitted at
	// Format: date-time
	SubmittedAt strfmt.DateTime `json:"submittedAt,omitempty"`
}

// Validate validates this queue payment request
func (m *QueuePaymentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartmentIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginGBLOC(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmittedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueuePaymentRequest) validateCustomer(formats strfmt.Registry) error {

	if swag.IsZero(m.Customer) { // not required
		return nil
	}

	if m.Customer != nil {
		if err := m.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *QueuePaymentRequest) validateDepartmentIndicator(formats strfmt.Registry) error {

	if swag.IsZero(m.DepartmentIndicator) { // not required
		return nil
	}

	if err := m.DepartmentIndicator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("departmentIndicator")
		}
		return err
	}

	return nil
}

func (m *QueuePaymentRequest) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueuePaymentRequest) validateMoveID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveID", "body", "uuid", m.MoveID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueuePaymentRequest) validateOriginGBLOC(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginGBLOC) { // not required
		return nil
	}

	if err := m.OriginGBLOC.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("originGBLOC")
		}
		return err
	}

	return nil
}

func (m *QueuePaymentRequest) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *QueuePaymentRequest) validateSubmittedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmittedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("submittedAt", "body", "date-time", m.SubmittedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueuePaymentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueuePaymentRequest) UnmarshalBinary(b []byte) error {
	var res QueuePaymentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
