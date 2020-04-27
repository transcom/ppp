// Code generated by go-swagger; DO NOT EDIT.

package supportmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOServiceItem m t o service item
// swagger:model MTOServiceItem
type MTOServiceItem struct {

	// approved at
	// Format: date
	ApprovedAt strfmt.Date `json:"approvedAt,omitempty"`

	// created at
	// Format: date
	CreatedAt strfmt.Date `json:"createdAt,omitempty"`

	// deleted at
	// Format: date
	DeletedAt strfmt.Date `json:"deletedAt,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// e tag
	ETag string `json:"eTag,omitempty"`

	// fee type
	// Enum: [COUNSELING CRATING TRUCKING SHUTTLE]
	FeeType string `json:"feeType,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// move task order ID
	// Required: true
	// Format: uuid
	MoveTaskOrderID *strfmt.UUID `json:"moveTaskOrderID"`

	// mto shipment ID
	// Required: true
	// Format: uuid
	MtoShipmentID *strfmt.UUID `json:"mtoShipmentID"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`

	// rate
	Rate int64 `json:"rate,omitempty"`

	// re service code
	// Required: true
	ReServiceCode *string `json:"reServiceCode"`

	// re service ID
	// Required: true
	// Format: uuid
	ReServiceID *strfmt.UUID `json:"reServiceID"`

	// re service name
	// Required: true
	ReServiceName *string `json:"reServiceName"`

	// reason
	Reason string `json:"reason,omitempty"`

	// rejected at
	// Format: date
	RejectedAt strfmt.Date `json:"rejectedAt,omitempty"`

	// status
	// Enum: [SUBMITTED APPROVED REJECTED]
	Status string `json:"status,omitempty"`

	// submitted at
	// Format: date
	SubmittedAt strfmt.Date `json:"submittedAt,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`

	// updated at
	// Format: datetime
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this m t o service item
func (m *MTOServiceItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApprovedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTaskOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtoShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReServiceCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReServiceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReServiceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejectedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmittedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MTOServiceItem) validateApprovedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ApprovedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("approvedAt", "body", "date", m.ApprovedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateDeletedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.DeletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("deletedAt", "body", "date", m.DeletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var mTOServiceItemTypeFeeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COUNSELING","CRATING","TRUCKING","SHUTTLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOServiceItemTypeFeeTypePropEnum = append(mTOServiceItemTypeFeeTypePropEnum, v)
	}
}

const (

	// MTOServiceItemFeeTypeCOUNSELING captures enum value "COUNSELING"
	MTOServiceItemFeeTypeCOUNSELING string = "COUNSELING"

	// MTOServiceItemFeeTypeCRATING captures enum value "CRATING"
	MTOServiceItemFeeTypeCRATING string = "CRATING"

	// MTOServiceItemFeeTypeTRUCKING captures enum value "TRUCKING"
	MTOServiceItemFeeTypeTRUCKING string = "TRUCKING"

	// MTOServiceItemFeeTypeSHUTTLE captures enum value "SHUTTLE"
	MTOServiceItemFeeTypeSHUTTLE string = "SHUTTLE"
)

// prop value enum
func (m *MTOServiceItem) validateFeeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, mTOServiceItemTypeFeeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MTOServiceItem) validateFeeType(formats strfmt.Registry) error {

	if swag.IsZero(m.FeeType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFeeTypeEnum("feeType", "body", m.FeeType); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if err := validate.Required("moveTaskOrderID", "body", m.MoveTaskOrderID); err != nil {
		return err
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateMtoShipmentID(formats strfmt.Registry) error {

	if err := validate.Required("mtoShipmentID", "body", m.MtoShipmentID); err != nil {
		return err
	}

	if err := validate.FormatOf("mtoShipmentID", "body", "uuid", m.MtoShipmentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateReServiceCode(formats strfmt.Registry) error {

	if err := validate.Required("reServiceCode", "body", m.ReServiceCode); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateReServiceID(formats strfmt.Registry) error {

	if err := validate.Required("reServiceID", "body", m.ReServiceID); err != nil {
		return err
	}

	if err := validate.FormatOf("reServiceID", "body", "uuid", m.ReServiceID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateReServiceName(formats strfmt.Registry) error {

	if err := validate.Required("reServiceName", "body", m.ReServiceName); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateRejectedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.RejectedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("rejectedAt", "body", "date", m.RejectedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var mTOServiceItemTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUBMITTED","APPROVED","REJECTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOServiceItemTypeStatusPropEnum = append(mTOServiceItemTypeStatusPropEnum, v)
	}
}

const (

	// MTOServiceItemStatusSUBMITTED captures enum value "SUBMITTED"
	MTOServiceItemStatusSUBMITTED string = "SUBMITTED"

	// MTOServiceItemStatusAPPROVED captures enum value "APPROVED"
	MTOServiceItemStatusAPPROVED string = "APPROVED"

	// MTOServiceItemStatusREJECTED captures enum value "REJECTED"
	MTOServiceItemStatusREJECTED string = "REJECTED"
)

// prop value enum
func (m *MTOServiceItem) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, mTOServiceItemTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MTOServiceItem) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateSubmittedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmittedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("submittedAt", "body", "date", m.SubmittedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOServiceItem) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "datetime", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOServiceItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOServiceItem) UnmarshalBinary(b []byte) error {
	var res MTOServiceItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
