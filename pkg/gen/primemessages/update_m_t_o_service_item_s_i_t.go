// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateMTOServiceItemSIT Subtype used to provide the departure date for origin or destination SIT. This is not creating a new service item but rather updating and existing service item.
//
//
// swagger:model UpdateMTOServiceItemSIT
type UpdateMTOServiceItemSIT struct {
	idField strfmt.UUID

	// Service code allowed for this model type.
	// Enum: [DDDSIT DOPSIT]
	ReServiceCode string `json:"reServiceCode,omitempty"`

	// Departure date for SIT. This is the end date of the SIT at either origin or destination.
	// Format: date
	SitDepartureDate strfmt.Date `json:"sitDepartureDate,omitempty"`

	// sit destination final address
	SitDestinationFinalAddress *Address `json:"sitDestinationFinalAddress,omitempty"`
}

// ID gets the id of this subtype
func (m *UpdateMTOServiceItemSIT) ID() strfmt.UUID {
	return m.idField
}

// SetID sets the id of this subtype
func (m *UpdateMTOServiceItemSIT) SetID(val strfmt.UUID) {
	m.idField = val
}

// ModelType gets the model type of this subtype
func (m *UpdateMTOServiceItemSIT) ModelType() UpdateMTOServiceItemModelType {
	return "UpdateMTOServiceItemSIT"
}

// SetModelType sets the model type of this subtype
func (m *UpdateMTOServiceItemSIT) SetModelType(val UpdateMTOServiceItemModelType) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *UpdateMTOServiceItemSIT) UnmarshalJSON(raw []byte) error {
	var data struct {

		// Service code allowed for this model type.
		// Enum: [DDDSIT DOPSIT]
		ReServiceCode string `json:"reServiceCode,omitempty"`

		// Departure date for SIT. This is the end date of the SIT at either origin or destination.
		// Format: date
		SitDepartureDate strfmt.Date `json:"sitDepartureDate,omitempty"`

		// sit destination final address
		SitDestinationFinalAddress *Address `json:"sitDestinationFinalAddress,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ID strfmt.UUID `json:"id,omitempty"`

		ModelType UpdateMTOServiceItemModelType `json:"modelType"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result UpdateMTOServiceItemSIT

	result.idField = base.ID

	if base.ModelType != result.ModelType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid modelType value: %q", base.ModelType)
	}

	result.ReServiceCode = data.ReServiceCode
	result.SitDepartureDate = data.SitDepartureDate
	result.SitDestinationFinalAddress = data.SitDestinationFinalAddress

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m UpdateMTOServiceItemSIT) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// Service code allowed for this model type.
		// Enum: [DDDSIT DOPSIT]
		ReServiceCode string `json:"reServiceCode,omitempty"`

		// Departure date for SIT. This is the end date of the SIT at either origin or destination.
		// Format: date
		SitDepartureDate strfmt.Date `json:"sitDepartureDate,omitempty"`

		// sit destination final address
		SitDestinationFinalAddress *Address `json:"sitDestinationFinalAddress,omitempty"`
	}{

		ReServiceCode: m.ReServiceCode,

		SitDepartureDate: m.SitDepartureDate,

		SitDestinationFinalAddress: m.SitDestinationFinalAddress,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ID strfmt.UUID `json:"id,omitempty"`

		ModelType UpdateMTOServiceItemModelType `json:"modelType"`
	}{

		ID: m.ID(),

		ModelType: m.ModelType(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this update m t o service item s i t
func (m *UpdateMTOServiceItemSIT) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReServiceCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSitDepartureDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSitDestinationFinalAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateMTOServiceItemSIT) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID().String(), formats); err != nil {
		return err
	}

	return nil
}

var updateMTOServiceItemSITTypeReServiceCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DDDSIT","DOPSIT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateMTOServiceItemSITTypeReServiceCodePropEnum = append(updateMTOServiceItemSITTypeReServiceCodePropEnum, v)
	}
}

// property enum
func (m *UpdateMTOServiceItemSIT) validateReServiceCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateMTOServiceItemSITTypeReServiceCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateMTOServiceItemSIT) validateReServiceCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ReServiceCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateReServiceCodeEnum("reServiceCode", "body", m.ReServiceCode); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMTOServiceItemSIT) validateSitDepartureDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SitDepartureDate) { // not required
		return nil
	}

	if err := validate.FormatOf("sitDepartureDate", "body", "date", m.SitDepartureDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMTOServiceItemSIT) validateSitDestinationFinalAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.SitDestinationFinalAddress) { // not required
		return nil
	}

	if m.SitDestinationFinalAddress != nil {
		if err := m.SitDestinationFinalAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sitDestinationFinalAddress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateMTOServiceItemSIT) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateMTOServiceItemSIT) UnmarshalBinary(b []byte) error {
	var res UpdateMTOServiceItemSIT
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
