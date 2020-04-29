// Code generated by go-swagger; DO NOT EDIT.

package primemessages

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

// MTOServiceItemDDFSIT Describes a domestic destination 1st day SIT service item subtype of a MTOServiceItem
// swagger:model MTOServiceItemDDFSIT
type MTOServiceItemDDFSIT struct {
	eTagField string

	idField strfmt.UUID

	moveTaskOrderIdField strfmt.UUID

	mtoShipmentIdField strfmt.UUID

	reServiceIdField strfmt.UUID

	reServiceNameField string

	rejectionReasonField *string

	statusField MTOServiceItemStatus

	// first available delivery date1
	// Required: true
	// Format: date
	FirstAvailableDeliveryDate1 *strfmt.Date `json:"firstAvailableDeliveryDate1"`

	// first available delivery date2
	// Required: true
	// Format: date
	FirstAvailableDeliveryDate2 *strfmt.Date `json:"firstAvailableDeliveryDate2"`

	// re service code
	ReServiceCode ReServiceCode `json:"reServiceCode,omitempty"`

	// time military1
	// Required: true
	TimeMilitary1 *string `json:"timeMilitary1"`

	// time military2
	// Required: true
	TimeMilitary2 *string `json:"timeMilitary2"`

	// type
	Type CustomerContactType `json:"type,omitempty"`
}

// ETag gets the e tag of this subtype
func (m *MTOServiceItemDDFSIT) ETag() string {
	return m.eTagField
}

// SetETag sets the e tag of this subtype
func (m *MTOServiceItemDDFSIT) SetETag(val string) {
	m.eTagField = val
}

// ID gets the id of this subtype
func (m *MTOServiceItemDDFSIT) ID() strfmt.UUID {
	return m.idField
}

// SetID sets the id of this subtype
func (m *MTOServiceItemDDFSIT) SetID(val strfmt.UUID) {
	m.idField = val
}

// ModelType gets the model type of this subtype
func (m *MTOServiceItemDDFSIT) ModelType() MTOServiceItemModelType {
	return "MTOServiceItemDDFSIT"
}

// SetModelType sets the model type of this subtype
func (m *MTOServiceItemDDFSIT) SetModelType(val MTOServiceItemModelType) {

}

// MoveTaskOrderID gets the move task order ID of this subtype
func (m *MTOServiceItemDDFSIT) MoveTaskOrderID() strfmt.UUID {
	return m.moveTaskOrderIdField
}

// SetMoveTaskOrderID sets the move task order ID of this subtype
func (m *MTOServiceItemDDFSIT) SetMoveTaskOrderID(val strfmt.UUID) {
	m.moveTaskOrderIdField = val
}

// MtoShipmentID gets the mto shipment ID of this subtype
func (m *MTOServiceItemDDFSIT) MtoShipmentID() strfmt.UUID {
	return m.mtoShipmentIdField
}

// SetMtoShipmentID sets the mto shipment ID of this subtype
func (m *MTOServiceItemDDFSIT) SetMtoShipmentID(val strfmt.UUID) {
	m.mtoShipmentIdField = val
}

// ReServiceID gets the re service ID of this subtype
func (m *MTOServiceItemDDFSIT) ReServiceID() strfmt.UUID {
	return m.reServiceIdField
}

// SetReServiceID sets the re service ID of this subtype
func (m *MTOServiceItemDDFSIT) SetReServiceID(val strfmt.UUID) {
	m.reServiceIdField = val
}

// ReServiceName gets the re service name of this subtype
func (m *MTOServiceItemDDFSIT) ReServiceName() string {
	return m.reServiceNameField
}

// SetReServiceName sets the re service name of this subtype
func (m *MTOServiceItemDDFSIT) SetReServiceName(val string) {
	m.reServiceNameField = val
}

// RejectionReason gets the rejection reason of this subtype
func (m *MTOServiceItemDDFSIT) RejectionReason() *string {
	return m.rejectionReasonField
}

// SetRejectionReason sets the rejection reason of this subtype
func (m *MTOServiceItemDDFSIT) SetRejectionReason(val *string) {
	m.rejectionReasonField = val
}

// Status gets the status of this subtype
func (m *MTOServiceItemDDFSIT) Status() MTOServiceItemStatus {
	return m.statusField
}

// SetStatus sets the status of this subtype
func (m *MTOServiceItemDDFSIT) SetStatus(val MTOServiceItemStatus) {
	m.statusField = val
}

// FirstAvailableDeliveryDate1 gets the first available delivery date1 of this subtype

// FirstAvailableDeliveryDate2 gets the first available delivery date2 of this subtype

// ReServiceCode gets the re service code of this subtype

// TimeMilitary1 gets the time military1 of this subtype

// TimeMilitary2 gets the time military2 of this subtype

// Type gets the type of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *MTOServiceItemDDFSIT) UnmarshalJSON(raw []byte) error {
	var data struct {

		// first available delivery date1
		// Required: true
		// Format: date
		FirstAvailableDeliveryDate1 *strfmt.Date `json:"firstAvailableDeliveryDate1"`

		// first available delivery date2
		// Required: true
		// Format: date
		FirstAvailableDeliveryDate2 *strfmt.Date `json:"firstAvailableDeliveryDate2"`

		// re service code
		ReServiceCode ReServiceCode `json:"reServiceCode,omitempty"`

		// time military1
		// Required: true
		TimeMilitary1 *string `json:"timeMilitary1"`

		// time military2
		// Required: true
		TimeMilitary2 *string `json:"timeMilitary2"`

		// type
		Type CustomerContactType `json:"type,omitempty"`
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

		RejectionReason *string `json:"rejectionReason,omitempty"`

		Status MTOServiceItemStatus `json:"status,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result MTOServiceItemDDFSIT

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

	result.rejectionReasonField = base.RejectionReason

	result.statusField = base.Status

	result.FirstAvailableDeliveryDate1 = data.FirstAvailableDeliveryDate1

	result.FirstAvailableDeliveryDate2 = data.FirstAvailableDeliveryDate2

	result.ReServiceCode = data.ReServiceCode

	result.TimeMilitary1 = data.TimeMilitary1

	result.TimeMilitary2 = data.TimeMilitary2

	result.Type = data.Type

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m MTOServiceItemDDFSIT) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// first available delivery date1
		// Required: true
		// Format: date
		FirstAvailableDeliveryDate1 *strfmt.Date `json:"firstAvailableDeliveryDate1"`

		// first available delivery date2
		// Required: true
		// Format: date
		FirstAvailableDeliveryDate2 *strfmt.Date `json:"firstAvailableDeliveryDate2"`

		// re service code
		ReServiceCode ReServiceCode `json:"reServiceCode,omitempty"`

		// time military1
		// Required: true
		TimeMilitary1 *string `json:"timeMilitary1"`

		// time military2
		// Required: true
		TimeMilitary2 *string `json:"timeMilitary2"`

		// type
		Type CustomerContactType `json:"type,omitempty"`
	}{

		FirstAvailableDeliveryDate1: m.FirstAvailableDeliveryDate1,

		FirstAvailableDeliveryDate2: m.FirstAvailableDeliveryDate2,

		ReServiceCode: m.ReServiceCode,

		TimeMilitary1: m.TimeMilitary1,

		TimeMilitary2: m.TimeMilitary2,

		Type: m.Type,
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

		RejectionReason *string `json:"rejectionReason,omitempty"`

		Status MTOServiceItemStatus `json:"status,omitempty"`
	}{

		ETag: m.ETag(),

		ID: m.ID(),

		ModelType: m.ModelType(),

		MoveTaskOrderID: m.MoveTaskOrderID(),

		MtoShipmentID: m.MtoShipmentID(),

		ReServiceID: m.ReServiceID(),

		ReServiceName: m.ReServiceName(),

		RejectionReason: m.RejectionReason(),

		Status: m.Status(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this m t o service item d d f s i t
func (m *MTOServiceItemDDFSIT) Validate(formats strfmt.Registry) error {
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

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstAvailableDeliveryDate1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstAvailableDeliveryDate2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReServiceCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeMilitary1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeMilitary2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MTOServiceItemDDFSIT) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID()) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDDFSIT) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrderID()) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDDFSIT) validateMtoShipmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.MtoShipmentID()) { // not required
		return nil
	}

	if err := validate.FormatOf("mtoShipmentID", "body", "uuid", m.MtoShipmentID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDDFSIT) validateReServiceID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReServiceID()) { // not required
		return nil
	}

	if err := validate.FormatOf("reServiceID", "body", "uuid", m.ReServiceID().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDDFSIT) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status()) { // not required
		return nil
	}

	if err := m.Status().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *MTOServiceItemDDFSIT) validateFirstAvailableDeliveryDate1(formats strfmt.Registry) error {

	if err := validate.Required("firstAvailableDeliveryDate1", "body", m.FirstAvailableDeliveryDate1); err != nil {
		return err
	}

	if err := validate.FormatOf("firstAvailableDeliveryDate1", "body", "date", m.FirstAvailableDeliveryDate1.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDDFSIT) validateFirstAvailableDeliveryDate2(formats strfmt.Registry) error {

	if err := validate.Required("firstAvailableDeliveryDate2", "body", m.FirstAvailableDeliveryDate2); err != nil {
		return err
	}

	if err := validate.FormatOf("firstAvailableDeliveryDate2", "body", "date", m.FirstAvailableDeliveryDate2.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDDFSIT) validateReServiceCode(formats strfmt.Registry) error {

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

func (m *MTOServiceItemDDFSIT) validateTimeMilitary1(formats strfmt.Registry) error {

	if err := validate.Required("timeMilitary1", "body", m.TimeMilitary1); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDDFSIT) validateTimeMilitary2(formats strfmt.Registry) error {

	if err := validate.Required("timeMilitary2", "body", m.TimeMilitary2); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItemDDFSIT) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOServiceItemDDFSIT) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOServiceItemDDFSIT) UnmarshalBinary(b []byte) error {
	var res MTOServiceItemDDFSIT
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
