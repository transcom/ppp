// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentServiceItem payment service item
// swagger:model PaymentServiceItem
type PaymentServiceItem struct {

	// e tag
	// Read Only: true
	ETag string `json:"eTag,omitempty"`

	// id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// mto service item ID
	// Format: uuid
	MtoServiceItemID strfmt.UUID `json:"mtoServiceItemID,omitempty"`

	// payment request ID
	// Format: uuid
	PaymentRequestID strfmt.UUID `json:"paymentRequestID,omitempty"`

	// payment service item params
	PaymentServiceItemParams PaymentServiceItemParams `json:"paymentServiceItemParams,omitempty"`

	// Price of the service item in cents
	PriceCents *int64 `json:"priceCents,omitempty"`

	// rejection reason
	RejectionReason *string `json:"rejectionReason,omitempty"`

	// status
	Status PaymentServiceItemStatus `json:"status,omitempty"`
}

// Validate validates this payment service item
func (m *PaymentServiceItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtoServiceItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentServiceItemParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentServiceItem) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentServiceItem) validateMtoServiceItemID(formats strfmt.Registry) error {

	if swag.IsZero(m.MtoServiceItemID) { // not required
		return nil
	}

	if err := validate.FormatOf("mtoServiceItemID", "body", "uuid", m.MtoServiceItemID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentServiceItem) validatePaymentRequestID(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentRequestID) { // not required
		return nil
	}

	if err := validate.FormatOf("paymentRequestID", "body", "uuid", m.PaymentRequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentServiceItem) validatePaymentServiceItemParams(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentServiceItemParams) { // not required
		return nil
	}

	if err := m.PaymentServiceItemParams.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("paymentServiceItemParams")
		}
		return err
	}

	return nil
}

func (m *PaymentServiceItem) validateStatus(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PaymentServiceItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentServiceItem) UnmarshalBinary(b []byte) error {
	var res PaymentServiceItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
