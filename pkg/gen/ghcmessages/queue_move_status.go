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

// QueueMoveStatus queue move status
//
// swagger:model QueueMoveStatus
type QueueMoveStatus string

const (

	// QueueMoveStatusNewMove captures enum value "New move"
	QueueMoveStatusNewMove QueueMoveStatus = "New move"

	// QueueMoveStatusMoveApproved captures enum value "Move approved"
	QueueMoveStatusMoveApproved QueueMoveStatus = "Move approved"

	// QueueMoveStatusApprovalsRequested captures enum value "Approvals requested"
	QueueMoveStatusApprovalsRequested QueueMoveStatus = "Approvals requested"
)

// for schema
var queueMoveStatusEnum []interface{}

func init() {
	var res []QueueMoveStatus
	if err := json.Unmarshal([]byte(`["New move","Move approved","Approvals requested"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queueMoveStatusEnum = append(queueMoveStatusEnum, v)
	}
}

func (m QueueMoveStatus) validateQueueMoveStatusEnum(path, location string, value QueueMoveStatus) error {
	if err := validate.EnumCase(path, location, value, queueMoveStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this queue move status
func (m QueueMoveStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateQueueMoveStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
