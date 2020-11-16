// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdatePaymentRequestStatusPayload update payment request status payload
//
// swagger:model UpdatePaymentRequestStatusPayload
type UpdatePaymentRequestStatusPayload struct {

	// e tag
	ETag string `json:"eTag,omitempty"`

	// rejection reason
	RejectionReason *string `json:"rejectionReason,omitempty"`

	// status
	Status PaymentRequestStatus `json:"status,omitempty"`
}

// Validate validates this update payment request status payload
func (m *UpdatePaymentRequestStatusPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatePaymentRequestStatusPayload) validateStatus(formats strfmt.Registry) error {

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
func (m *UpdatePaymentRequestStatusPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePaymentRequestStatusPayload) UnmarshalBinary(b []byte) error {
	var res UpdatePaymentRequestStatusPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
