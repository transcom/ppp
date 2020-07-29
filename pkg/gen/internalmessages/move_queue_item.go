// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MoveQueueItem move queue item
// swagger:model MoveQueueItem
type MoveQueueItem struct {

	// actual move date
	// Format: date
	ActualMoveDate *strfmt.Date `json:"actual_move_date,omitempty"`

	// branch of service
	// Required: true
	BranchOfService *string `json:"branch_of_service"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// Customer Name
	// Required: true
	CustomerName *string `json:"customer_name"`

	// delivered date
	// Format: date-time
	DeliveredDate *strfmt.DateTime `json:"delivered_date,omitempty"`

	// Destination
	DestinationDutyStationName *string `json:"destination_duty_station_name,omitempty"`

	// Destination GBLOC
	DestinationGbloc *string `json:"destination_gbloc,omitempty"`

	// DoD ID #
	// Required: true
	// Max Length: 10
	// Min Length: 10
	// Pattern: ^\d{10}$
	Edipi *string `json:"edipi"`

	// GBL Number
	GblNumber *string `json:"gbl_number,omitempty"`

	// hhg status
	HhgStatus *string `json:"hhg_status,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// invoice approved date
	// Format: date-time
	InvoiceApprovedDate *strfmt.DateTime `json:"invoice_approved_date,omitempty"`

	// last modified date
	// Required: true
	// Format: date-time
	LastModifiedDate *strfmt.DateTime `json:"last_modified_date"`

	// locator
	// Required: true
	Locator *string `json:"locator"`

	// move date
	// Format: date
	MoveDate *strfmt.Date `json:"move_date,omitempty"`

	// Move Type
	// Required: true
	// Enum: [PCS - OCONUS PCS - CONUS PCS + TDY - OCONUS PCS + TDY - CONUS]
	OrdersType *string `json:"orders_type"`

	// Origin
	OriginDutyStationName *string `json:"origin_duty_station_name,omitempty"`

	// Origin GBLOC
	OriginGbloc *string `json:"origin_gbloc,omitempty"`

	// original move date
	// Format: date
	OriginalMoveDate *strfmt.Date `json:"original_move_date,omitempty"`

	// pm survey conducted date
	// Format: date-time
	PmSurveyConductedDate *strfmt.DateTime `json:"pm_survey_conducted_date,omitempty"`

	// ppm status
	PpmStatus *string `json:"ppm_status,omitempty"`

	// rank
	// Required: true
	Rank *ServiceMemberRank `json:"rank"`

	// status
	// Required: true
	Status *string `json:"status"`

	// submitted date
	// Format: date-time
	SubmittedDate *strfmt.DateTime `json:"submitted_date,omitempty"`

	// weight allotment
	WeightAllotment *WeightAllotment `json:"weight_allotment,omitempty"`
}

// Validate validates this move queue item
func (m *MoveQueueItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActualMoveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBranchOfService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveredDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdipi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceApprovedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrdersType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalMoveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmSurveyConductedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmittedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeightAllotment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveQueueItem) validateActualMoveDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ActualMoveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("actual_move_date", "body", "date", m.ActualMoveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateBranchOfService(formats strfmt.Registry) error {

	if err := validate.Required("branch_of_service", "body", m.BranchOfService); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateCustomerName(formats strfmt.Registry) error {

	if err := validate.Required("customer_name", "body", m.CustomerName); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateDeliveredDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DeliveredDate) { // not required
		return nil
	}

	if err := validate.FormatOf("delivered_date", "body", "date-time", m.DeliveredDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateEdipi(formats strfmt.Registry) error {

	if err := validate.Required("edipi", "body", m.Edipi); err != nil {
		return err
	}

	if err := validate.MinLength("edipi", "body", string(*m.Edipi), 10); err != nil {
		return err
	}

	if err := validate.MaxLength("edipi", "body", string(*m.Edipi), 10); err != nil {
		return err
	}

	if err := validate.Pattern("edipi", "body", string(*m.Edipi), `^\d{10}$`); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateInvoiceApprovedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.InvoiceApprovedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("invoice_approved_date", "body", "date-time", m.InvoiceApprovedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateLastModifiedDate(formats strfmt.Registry) error {

	if err := validate.Required("last_modified_date", "body", m.LastModifiedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("last_modified_date", "body", "date-time", m.LastModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateLocator(formats strfmt.Registry) error {

	if err := validate.Required("locator", "body", m.Locator); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateMoveDate(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("move_date", "body", "date", m.MoveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var moveQueueItemTypeOrdersTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PCS - OCONUS","PCS - CONUS","PCS + TDY - OCONUS","PCS + TDY - CONUS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		moveQueueItemTypeOrdersTypePropEnum = append(moveQueueItemTypeOrdersTypePropEnum, v)
	}
}

const (

	// MoveQueueItemOrdersTypePCSOCONUS captures enum value "PCS - OCONUS"
	MoveQueueItemOrdersTypePCSOCONUS string = "PCS - OCONUS"

	// MoveQueueItemOrdersTypePCSCONUS captures enum value "PCS - CONUS"
	MoveQueueItemOrdersTypePCSCONUS string = "PCS - CONUS"

	// MoveQueueItemOrdersTypePCSTDYOCONUS captures enum value "PCS + TDY - OCONUS"
	MoveQueueItemOrdersTypePCSTDYOCONUS string = "PCS + TDY - OCONUS"

	// MoveQueueItemOrdersTypePCSTDYCONUS captures enum value "PCS + TDY - CONUS"
	MoveQueueItemOrdersTypePCSTDYCONUS string = "PCS + TDY - CONUS"
)

// prop value enum
func (m *MoveQueueItem) validateOrdersTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, moveQueueItemTypeOrdersTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MoveQueueItem) validateOrdersType(formats strfmt.Registry) error {

	if err := validate.Required("orders_type", "body", m.OrdersType); err != nil {
		return err
	}

	// value enum
	if err := m.validateOrdersTypeEnum("orders_type", "body", *m.OrdersType); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateOriginalMoveDate(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginalMoveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("original_move_date", "body", "date", m.OriginalMoveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validatePmSurveyConductedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.PmSurveyConductedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("pm_survey_conducted_date", "body", "date-time", m.PmSurveyConductedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateRank(formats strfmt.Registry) error {

	if err := validate.Required("rank", "body", m.Rank); err != nil {
		return err
	}

	if m.Rank != nil {
		if err := m.Rank.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rank")
			}
			return err
		}
	}

	return nil
}

func (m *MoveQueueItem) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateSubmittedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmittedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("submitted_date", "body", "date-time", m.SubmittedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MoveQueueItem) validateWeightAllotment(formats strfmt.Registry) error {

	if swag.IsZero(m.WeightAllotment) { // not required
		return nil
	}

	if m.WeightAllotment != nil {
		if err := m.WeightAllotment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight_allotment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveQueueItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveQueueItem) UnmarshalBinary(b []byte) error {
	var res MoveQueueItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
