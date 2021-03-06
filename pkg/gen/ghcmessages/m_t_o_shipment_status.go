// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// MTOShipmentStatus Shipment Status
//
// swagger:model MTOShipmentStatus
type MTOShipmentStatus string

const (

	// MTOShipmentStatusSUBMITTED captures enum value "SUBMITTED"
	MTOShipmentStatusSUBMITTED MTOShipmentStatus = "SUBMITTED"

	// MTOShipmentStatusREJECTED captures enum value "REJECTED"
	MTOShipmentStatusREJECTED MTOShipmentStatus = "REJECTED"

	// MTOShipmentStatusAPPROVED captures enum value "APPROVED"
	MTOShipmentStatusAPPROVED MTOShipmentStatus = "APPROVED"

	// MTOShipmentStatusCANCELLATIONREQUESTED captures enum value "CANCELLATION_REQUESTED"
	MTOShipmentStatusCANCELLATIONREQUESTED MTOShipmentStatus = "CANCELLATION_REQUESTED"
)

// for schema
var mTOShipmentStatusEnum []interface{}

func init() {
	var res []MTOShipmentStatus
	if err := json.Unmarshal([]byte(`["SUBMITTED","REJECTED","APPROVED","CANCELLATION_REQUESTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOShipmentStatusEnum = append(mTOShipmentStatusEnum, v)
	}
}

func (m MTOShipmentStatus) validateMTOShipmentStatusEnum(path, location string, value MTOShipmentStatus) error {
	if err := validate.EnumCase(path, location, value, mTOShipmentStatusEnum, true); err != nil {
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
