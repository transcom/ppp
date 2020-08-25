// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// MTOShipmentStatus m t o shipment status
// swagger:model MTOShipmentStatus
type MTOShipmentStatus string

const (

	// MTOShipmentStatusDRAFT captures enum value "DRAFT"
	MTOShipmentStatusDRAFT MTOShipmentStatus = "DRAFT"

	// MTOShipmentStatusAPPROVED captures enum value "APPROVED"
	MTOShipmentStatusAPPROVED MTOShipmentStatus = "APPROVED"

	// MTOShipmentStatusSUBMITTED captures enum value "SUBMITTED"
	MTOShipmentStatusSUBMITTED MTOShipmentStatus = "SUBMITTED"

	// MTOShipmentStatusREJECTED captures enum value "REJECTED"
	MTOShipmentStatusREJECTED MTOShipmentStatus = "REJECTED"
)

// for schema
var mTOShipmentStatusEnum []interface{}

func init() {
	var res []MTOShipmentStatus
	if err := json.Unmarshal([]byte(`["DRAFT","APPROVED","SUBMITTED","REJECTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOShipmentStatusEnum = append(mTOShipmentStatusEnum, v)
	}
}

func (m MTOShipmentStatus) validateMTOShipmentStatusEnum(path, location string, value MTOShipmentStatus) error {
	if err := validate.Enum(path, location, value, mTOShipmentStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this m t o shipment status
func (m MTOShipmentStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMTOShipmentStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
