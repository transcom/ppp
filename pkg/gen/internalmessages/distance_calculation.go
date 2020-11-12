// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DistanceCalculation distance calculation
//
// swagger:model DistanceCalculation
type DistanceCalculation struct {

	// destination address
	DestinationAddress *Address `json:"destination_address,omitempty"`

	// Distance
	//
	// unit is miles
	DistanceMiles int64 `json:"distance_miles,omitempty"`

	// origin address
	OriginAddress *Address `json:"origin_address,omitempty"`
}

// Validate validates this distance calculation
func (m *DistanceCalculation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistanceCalculation) validateDestinationAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationAddress) { // not required
		return nil
	}

	if m.DestinationAddress != nil {
		if err := m.DestinationAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination_address")
			}
			return err
		}
	}

	return nil
}

func (m *DistanceCalculation) validateOriginAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginAddress) { // not required
		return nil
	}

	if m.OriginAddress != nil {
		if err := m.OriginAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin_address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DistanceCalculation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistanceCalculation) UnmarshalBinary(b []byte) error {
	var res DistanceCalculation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
