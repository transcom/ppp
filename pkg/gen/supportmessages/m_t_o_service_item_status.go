// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// MTOServiceItemStatus Describes all statuses for a MTOServiceItem
// swagger:model MTOServiceItemStatus
type MTOServiceItemStatus string

const (

	// MTOServiceItemStatusSUBMITTED captures enum value "SUBMITTED"
	MTOServiceItemStatusSUBMITTED MTOServiceItemStatus = "SUBMITTED"

	// MTOServiceItemStatusAPPROVED captures enum value "APPROVED"
	MTOServiceItemStatusAPPROVED MTOServiceItemStatus = "APPROVED"

	// MTOServiceItemStatusREJECTED captures enum value "REJECTED"
	MTOServiceItemStatusREJECTED MTOServiceItemStatus = "REJECTED"
)

// for schema
var mTOServiceItemStatusEnum []interface{}

func init() {
	var res []MTOServiceItemStatus
	if err := json.Unmarshal([]byte(`["SUBMITTED","APPROVED","REJECTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOServiceItemStatusEnum = append(mTOServiceItemStatusEnum, v)
	}
}

func (m MTOServiceItemStatus) validateMTOServiceItemStatusEnum(path, location string, value MTOServiceItemStatus) error {
	if err := validate.Enum(path, location, value, mTOServiceItemStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this m t o service item status
func (m MTOServiceItemStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMTOServiceItemStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
