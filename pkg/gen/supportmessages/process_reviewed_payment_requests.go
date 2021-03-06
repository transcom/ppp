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

// ProcessReviewedPaymentRequests process reviewed payment requests
//
// swagger:model ProcessReviewedPaymentRequests
type ProcessReviewedPaymentRequests struct {

	// payment request ID
	// Read Only: true
	// Format: uuid
	PaymentRequestID strfmt.UUID `json:"paymentRequestID,omitempty"`

	// send to syncada
	// Required: true
	SendToSyncada *bool `json:"sendToSyncada"`

	// status
	Status PaymentRequestStatus `json:"status,omitempty"`
}

// Validate validates this process reviewed payment requests
func (m *ProcessReviewedPaymentRequests) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSendToSyncada(formats); err != nil {
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

func (m *ProcessReviewedPaymentRequests) validatePaymentRequestID(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentRequestID) { // not required
		return nil
	}

	if err := validate.FormatOf("paymentRequestID", "body", "uuid", m.PaymentRequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProcessReviewedPaymentRequests) validateSendToSyncada(formats strfmt.Registry) error {

	if err := validate.Required("sendToSyncada", "body", m.SendToSyncada); err != nil {
		return err
	}

	return nil
}

func (m *ProcessReviewedPaymentRequests) validateStatus(formats strfmt.Registry) error {

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
func (m *ProcessReviewedPaymentRequests) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProcessReviewedPaymentRequests) UnmarshalBinary(b []byte) error {
	var res ProcessReviewedPaymentRequests
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
