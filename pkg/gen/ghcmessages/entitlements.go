// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Entitlements entitlements
// swagger:model entitlements
type Entitlements struct {

	// dependents authorized
	DependentsAuthorized bool `json:"dependentsAuthorized,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// non temporary storage
	NonTemporaryStorage *bool `json:"nonTemporaryStorage,omitempty"`

	// privately owned vehicle
	PrivatelyOwnedVehicle *bool `json:"privatelyOwnedVehicle,omitempty"`

	// pro gear weight
	ProGearWeight int64 `json:"proGearWeight,omitempty"`

	// pro gear weight spouse
	ProGearWeightSpouse int64 `json:"proGearWeightSpouse,omitempty"`

	// storage in transit
	StorageInTransit int64 `json:"storageInTransit,omitempty"`

	// total dependents
	TotalDependents int64 `json:"totalDependents,omitempty"`

	// total weight self
	TotalWeightSelf int64 `json:"totalWeightSelf,omitempty"`
}

// Validate validates this entitlements
func (m *Entitlements) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Entitlements) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Entitlements) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Entitlements) UnmarshalBinary(b []byte) error {
	var res Entitlements
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
