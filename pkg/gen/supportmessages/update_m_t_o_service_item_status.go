// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateMTOServiceItemStatus update m t o service item status
//
// swagger:model UpdateMTOServiceItemStatus
type UpdateMTOServiceItemStatus struct {

	// Reason the service item was rejected by the TOO""
	RejectionReason *string `json:"rejectionReason,omitempty"`

	// status
	Status MTOServiceItemStatus `json:"status,omitempty"`
}

// Validate validates this update m t o service item status
func (m *UpdateMTOServiceItemStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateMTOServiceItemStatus) validateStatus(formats strfmt.Registry) error {

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
func (m *UpdateMTOServiceItemStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateMTOServiceItemStatus) UnmarshalBinary(b []byte) error {
	var res UpdateMTOServiceItemStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
