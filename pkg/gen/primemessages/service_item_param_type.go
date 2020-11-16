// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ServiceItemParamType service item param type
//
// swagger:model ServiceItemParamType
type ServiceItemParamType string

const (

	// ServiceItemParamTypeSTRING captures enum value "STRING"
	ServiceItemParamTypeSTRING ServiceItemParamType = "STRING"

	// ServiceItemParamTypeDATE captures enum value "DATE"
	ServiceItemParamTypeDATE ServiceItemParamType = "DATE"

	// ServiceItemParamTypeINTEGER captures enum value "INTEGER"
	ServiceItemParamTypeINTEGER ServiceItemParamType = "INTEGER"

	// ServiceItemParamTypeDECIMAL captures enum value "DECIMAL"
	ServiceItemParamTypeDECIMAL ServiceItemParamType = "DECIMAL"

	// ServiceItemParamTypeTIMESTAMP captures enum value "TIMESTAMP"
	ServiceItemParamTypeTIMESTAMP ServiceItemParamType = "TIMESTAMP"

	// ServiceItemParamTypePaymentServiceItemUUID captures enum value "PaymentServiceItemUUID"
	ServiceItemParamTypePaymentServiceItemUUID ServiceItemParamType = "PaymentServiceItemUUID"
)

// for schema
var serviceItemParamTypeEnum []interface{}

func init() {
	var res []ServiceItemParamType
	if err := json.Unmarshal([]byte(`["STRING","DATE","INTEGER","DECIMAL","TIMESTAMP","PaymentServiceItemUUID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceItemParamTypeEnum = append(serviceItemParamTypeEnum, v)
	}
}

func (m ServiceItemParamType) validateServiceItemParamTypeEnum(path, location string, value ServiceItemParamType) error {
	if err := validate.EnumCase(path, location, value, serviceItemParamTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this service item param type
func (m ServiceItemParamType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateServiceItemParamTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
