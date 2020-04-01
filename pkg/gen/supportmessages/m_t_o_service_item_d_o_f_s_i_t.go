// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOServiceItemDOFSIT Describes a domestic origin 1st day SIT service item subtype of a MTOServiceItem
// swagger:model MTOServiceItemDOFSIT
type MTOServiceItemDOFSIT struct {
	eTagField string

	idField strfmt.UUID

	moveTaskOrderIdField strfmt.UUID

	mtoShipmentIdField strfmt.UUID

	reServiceIdField strfmt.UUID

	reServiceNameField string

	// pickup postal code
	// Required: true
	// Pattern: ^(\d{5}([\-]\d{4})?)$
	PickupPostalCode *string `json:"pickupPostalCode"`

	// re service code
	ReServiceCode ReServiceCode `json:"reServiceCode,omitempty"`

	// reason
	// Required: true
	Reason *string `json:"reason"`
}

// ETag gets the e tag of this subtype
func (m *MTOServiceItemDOFSIT) ETag() string {
	return m.eTagField
}

// SetETag sets the e tag of this subtype
func (m *MTOServiceItemDOFSIT) SetETag(val string) {
	m.eTagField = val
}

// ID gets the id of this subtype
func (m *MTOServiceItemDOFSIT) ID() strfmt.UUID {
	return m.idField
}

// SetID sets the id of this subtype
func (m *MTOServiceItemDOFSIT) SetID(val strfmt.UUID) {
	m.idField = val
}

// ModelType gets the model type of this subtype
func (m *MTOServiceItemDOFSIT) ModelType() MTOServiceItemModelType {
	return "MTOServiceItemDOFSIT"
}

// SetModelType sets the model type of this subtype
func (m *MTOServiceItemDOFSIT) SetModelType(val MTOServiceItemModelType) {

}

// MoveTaskOrderID gets the move task order ID of this subtype
func (m *MTOServiceItemDOFSIT) MoveTaskOrderID() strfmt.UUID {
	return m.moveTaskOrderIdField
}

// SetMoveTaskOrderID sets the move task order ID of this subtype
func (m *MTOServiceItemDOFSIT) SetMoveTaskOrderID(val strfmt.UUID) {
	m.moveTaskOrderIdField = val
}

// MtoShipmentID gets the mto shipment ID of this subtype
func (m *MTOServiceItemDOFSIT) MtoShipmentID() strfmt.UUID {
	return m.mtoShipmentIdField
}

// SetMtoShipmentID sets the mto shipment ID of this subtype
func (m *MTOServiceItemDOFSIT) SetMtoShipmentID(val strfmt.UUID) {
	m.mtoShipmentIdField = val
}

// ReServiceID gets the re service ID of this subtype
func (m *MTOServiceItemDOFSIT) ReServiceID() strfmt.UUID {
	return m.reServiceIdField
}

// SetReServiceID sets the re service ID of this subtype
func (m *MTOServiceItemDOFSIT) SetReServiceID(val strfmt.UUID) {
	m.reServiceIdField = val
}

// ReServiceName gets the re service name of this subtype
func (m *MTOServiceItemDOFSIT) ReServiceName() string {
	return m.reServiceNameField
}

// SetReServiceName sets the re service name of this subtype
func (m *MTOServiceItemDOFSIT) SetReServiceName(val string) {
	m.reServiceNameField = val
}

// PickupPostalCode gets the pickup postal code of this subtype

// ReServiceCode gets the re service code of this subtype

// Reason gets the reason of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *MTOServiceItemDOFSIT) UnmarshalJSON(raw []byte) error {
	var data struct {

		// pickup postal code
		// Required: true
		// Pattern: ^(\d{5}([\-]\d{4})?)$
		PickupPostalCode *string `json:"pickupPostalCode"`

		// re service code
		ReServiceCode ReServiceCode `json:"reServiceCode,omitempty"`

		// reason
		// Required: true
		Reason *string `json:"reason"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ETag string `json:"eTag,omitempty"`

		ID strfmt.UUID `json:"id,omitempty"`

		ModelType MTOServiceItemModelType `json:"modelType"`

		MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

		MtoShipmentID strfmt.UUID `json:"mtoShipmentID,omitempty"`

		ReServiceID strfmt.UUID `json:"reServiceID,omitempty"`

		ReServiceName string `json:"reServiceName,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result MTOServiceItemDOFSIT

	result.eTagField = base.ETag

	result.idField = base.ID

	if base.ModelType != result.ModelType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid modelType value: %q", base.ModelType)
	}

	result.moveTaskOrderIdField = base.MoveTaskOrderID

	result.mtoShipmentIdField = base.MtoShipmentID

	result.reServiceIdField = base.ReServiceID

	result.reServiceNameField = base.ReServiceName

	result.PickupPostalCode = data.PickupPostalCode

	result.ReServiceCode = data.ReServiceCode

	result.Reason = data.Reason

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m MTOServiceItemDOFSIT) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// pickup postal code
		// Required: true
		// Pattern: ^(\d{5}([\-]\d{4})?)$
		PickupPostalCode *string `json:"pickupPostalCode"`

		// re service code
		ReServiceCode ReServiceCode `json:"reServiceCode,omitempty"`

		// reason
		// Required: true
		Reason *string `json:"reason"`
	}{

		PickupPostalCode: m.PickupPostalCode,

		ReServiceCode: m.ReServiceCode,

		Reason: m.Reason,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ETag string `json:"eTag,omitempty"`

		ID strfmt.UUID `json:"id,omitempty"`

		ModelType MTOServiceItemModelType `json:"modelType"`

		MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

		MtoShipmentID strfmt.UUID `json:"mtoShipmentID,omitempty"`

		ReServiceID strfmt.UUID `json:"reServiceID,omitempty"`

		ReServiceName string `json:"reServiceName,omitempty"`
	}{

		ETag: m.ETag(),

		ID: m.ID(),

		ModelType: m.ModelType(),

		MoveTaskOrderID: m.MoveTaskOrderID(),

		MtoShipmentID: m.MtoShipmentID(),

		ReServiceID: m.ReServiceID(),

		ReServiceName: m.ReServiceName(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this m t o service item d o f s i t
func (m *MTOServiceItemDOFSIT) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTaskOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtoShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReServiceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickupPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReServiceCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MTOServiceItemDOFSIT) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDOFSIT) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrderID()) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDOFSIT) validateMtoShipmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.MtoShipmentID()) { // not required
		return nil
	}

	if err := validate.FormatOf("mtoShipmentID", "body", "uuid", m.MtoShipmentID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDOFSIT) validateReServiceID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReServiceID()) { // not required
		return nil
	}

	if err := validate.FormatOf("reServiceID", "body", "uuid", m.ReServiceID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDOFSIT) validatePickupPostalCode(formats strfmt.Registry) error {

	if err := validate.Required("pickupPostalCode", "body", m.PickupPostalCode); err != nil {
		return err
	}

	if err := validate.Pattern("pickupPostalCode", "body", string(*m.PickupPostalCode), `^(\d{5}([\-]\d{4})?)$`); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDOFSIT) validateReServiceCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ReServiceCode) { // not required
		return nil
	}

	if err := m.ReServiceCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reServiceCode")
		}
		return err
	}

	return nil
}

func (m *MTOServiceItemDOFSIT) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOServiceItemDOFSIT) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOServiceItemDOFSIT) UnmarshalBinary(b []byte) error {
	var res MTOServiceItemDOFSIT
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
