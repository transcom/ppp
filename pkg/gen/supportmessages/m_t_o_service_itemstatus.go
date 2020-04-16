// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOServiceItemstatus m t o service itemstatus
// swagger:model MTOServiceItemstatus
type MTOServiceItemstatus struct {

	// status
	// Enum: [APPROVED SUBMITTED REJECTED]
	Status string `json:"status,omitempty"`
}

// Validate validates this m t o service itemstatus
func (m *MTOServiceItemstatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mTOServiceItemstatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APPROVED","SUBMITTED","REJECTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOServiceItemstatusTypeStatusPropEnum = append(mTOServiceItemstatusTypeStatusPropEnum, v)
	}
}

const (

	// MTOServiceItemstatusStatusAPPROVED captures enum value "APPROVED"
	MTOServiceItemstatusStatusAPPROVED string = "APPROVED"

	// MTOServiceItemstatusStatusSUBMITTED captures enum value "SUBMITTED"
	MTOServiceItemstatusStatusSUBMITTED string = "SUBMITTED"

	// MTOServiceItemstatusStatusREJECTED captures enum value "REJECTED"
	MTOServiceItemstatusStatusREJECTED string = "REJECTED"
)

// prop value enum
func (m *MTOServiceItemstatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, mTOServiceItemstatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MTOServiceItemstatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOServiceItemstatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOServiceItemstatus) UnmarshalBinary(b []byte) error {
	var res MTOServiceItemstatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}