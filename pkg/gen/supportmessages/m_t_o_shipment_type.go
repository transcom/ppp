// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// MTOShipmentType Shipment Type
//
// swagger:model MTOShipmentType
type MTOShipmentType string

const (

	// MTOShipmentTypeHHG captures enum value "HHG"
	MTOShipmentTypeHHG MTOShipmentType = "HHG"

	// MTOShipmentTypeINTERNATIONALHHG captures enum value "INTERNATIONAL_HHG"
	MTOShipmentTypeINTERNATIONALHHG MTOShipmentType = "INTERNATIONAL_HHG"

	// MTOShipmentTypeINTERNATIONALUB captures enum value "INTERNATIONAL_UB"
	MTOShipmentTypeINTERNATIONALUB MTOShipmentType = "INTERNATIONAL_UB"
)

// for schema
var mTOShipmentTypeEnum []interface{}

func init() {
	var res []MTOShipmentType
	if err := json.Unmarshal([]byte(`["HHG","INTERNATIONAL_HHG","INTERNATIONAL_UB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOShipmentTypeEnum = append(mTOShipmentTypeEnum, v)
	}
}

func (m MTOShipmentType) validateMTOShipmentTypeEnum(path, location string, value MTOShipmentType) error {
	if err := validate.EnumCase(path, location, value, mTOShipmentTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this m t o shipment type
func (m MTOShipmentType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMTOShipmentTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}