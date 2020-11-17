// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// TShirtSize Size
//
// swagger:model TShirtSize
type TShirtSize string

const (

	// TShirtSizeS captures enum value "S"
	TShirtSizeS TShirtSize = "S"

	// TShirtSizeM captures enum value "M"
	TShirtSizeM TShirtSize = "M"

	// TShirtSizeL captures enum value "L"
	TShirtSizeL TShirtSize = "L"
)

// for schema
var tShirtSizeEnum []interface{}

func init() {
	var res []TShirtSize
	if err := json.Unmarshal([]byte(`["S","M","L"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tShirtSizeEnum = append(tShirtSizeEnum, v)
	}
}

func (m TShirtSize) validateTShirtSizeEnum(path, location string, value TShirtSize) error {
	if err := validate.EnumCase(path, location, value, tShirtSizeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this t shirt size
func (m TShirtSize) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTShirtSizeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
