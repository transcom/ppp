// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Orders orders
//
// swagger:model Orders
type Orders struct {

	// authorized weight
	AuthorizedWeight *int64 `json:"authorizedWeight,omitempty"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// department indicator
	DepartmentIndicator *DeptIndicator `json:"department_indicator,omitempty"`

	// grade
	Grade *string `json:"grade,omitempty"`

	// Are dependents included in your orders?
	// Required: true
	HasDependents *bool `json:"has_dependents"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Date issued
	//
	// The date and time that these orders were cut.
	// Required: true
	// Format: date
	IssueDate *strfmt.Date `json:"issue_date"`

	// moves
	Moves IndexMovesPayload `json:"moves,omitempty"`

	// new duty station
	// Required: true
	NewDutyStation *DutyStationPayload `json:"new_duty_station"`

	// Orders Number
	OrdersNumber *string `json:"orders_number,omitempty"`

	// orders type
	// Required: true
	OrdersType OrdersType `json:"orders_type"`

	// orders type detail
	OrdersTypeDetail *OrdersTypeDetail `json:"orders_type_detail,omitempty"`

	// origin duty station
	OriginDutyStation *DutyStationPayload `json:"origin_duty_station,omitempty"`

	// Report by
	//
	// Report By Date
	// Required: true
	// Format: date
	ReportByDate *strfmt.Date `json:"report_by_date"`

	// SAC
	Sac *string `json:"sac,omitempty"`

	// service member id
	// Required: true
	// Format: uuid
	ServiceMemberID *strfmt.UUID `json:"service_member_id"`

	// Do you have a spouse who will need to move items related to their occupation (also known as spouse pro-gear)?
	// Required: true
	SpouseHasProGear *bool `json:"spouse_has_pro_gear"`

	// status
	Status OrdersStatus `json:"status,omitempty"`

	// TAC
	Tac *string `json:"tac,omitempty"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	// uploaded orders
	// Required: true
	UploadedOrders *DocumentPayload `json:"uploaded_orders"`
}

// Validate validates this orders
func (m *Orders) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartmentIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasDependents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoves(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrdersType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrdersTypeDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginDutyStation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportByDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceMemberID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpouseHasProGear(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadedOrders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Orders) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Orders) validateDepartmentIndicator(formats strfmt.Registry) error {

	if swag.IsZero(m.DepartmentIndicator) { // not required
		return nil
	}

	if m.DepartmentIndicator != nil {
		if err := m.DepartmentIndicator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("department_indicator")
			}
			return err
		}
	}

	return nil
}

func (m *Orders) validateHasDependents(formats strfmt.Registry) error {

	if err := validate.Required("has_dependents", "body", m.HasDependents); err != nil {
		return err
	}

	return nil
}

func (m *Orders) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Orders) validateIssueDate(formats strfmt.Registry) error {

	if err := validate.Required("issue_date", "body", m.IssueDate); err != nil {
		return err
	}

	if err := validate.FormatOf("issue_date", "body", "date", m.IssueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Orders) validateMoves(formats strfmt.Registry) error {

	if swag.IsZero(m.Moves) { // not required
		return nil
	}

	if err := m.Moves.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("moves")
		}
		return err
	}

	return nil
}

func (m *Orders) validateNewDutyStation(formats strfmt.Registry) error {

	if err := validate.Required("new_duty_station", "body", m.NewDutyStation); err != nil {
		return err
	}

	if m.NewDutyStation != nil {
		if err := m.NewDutyStation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("new_duty_station")
			}
			return err
		}
	}

	return nil
}

func (m *Orders) validateOrdersType(formats strfmt.Registry) error {

	if err := m.OrdersType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orders_type")
		}
		return err
	}

	return nil
}

func (m *Orders) validateOrdersTypeDetail(formats strfmt.Registry) error {

	if swag.IsZero(m.OrdersTypeDetail) { // not required
		return nil
	}

	if m.OrdersTypeDetail != nil {
		if err := m.OrdersTypeDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orders_type_detail")
			}
			return err
		}
	}

	return nil
}

func (m *Orders) validateOriginDutyStation(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginDutyStation) { // not required
		return nil
	}

	if m.OriginDutyStation != nil {
		if err := m.OriginDutyStation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin_duty_station")
			}
			return err
		}
	}

	return nil
}

func (m *Orders) validateReportByDate(formats strfmt.Registry) error {

	if err := validate.Required("report_by_date", "body", m.ReportByDate); err != nil {
		return err
	}

	if err := validate.FormatOf("report_by_date", "body", "date", m.ReportByDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Orders) validateServiceMemberID(formats strfmt.Registry) error {

	if err := validate.Required("service_member_id", "body", m.ServiceMemberID); err != nil {
		return err
	}

	if err := validate.FormatOf("service_member_id", "body", "uuid", m.ServiceMemberID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Orders) validateSpouseHasProGear(formats strfmt.Registry) error {

	if err := validate.Required("spouse_has_pro_gear", "body", m.SpouseHasProGear); err != nil {
		return err
	}

	return nil
}

func (m *Orders) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *Orders) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Orders) validateUploadedOrders(formats strfmt.Registry) error {

	if err := validate.Required("uploaded_orders", "body", m.UploadedOrders); err != nil {
		return err
	}

	if m.UploadedOrders != nil {
		if err := m.UploadedOrders.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploaded_orders")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Orders) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Orders) UnmarshalBinary(b []byte) error {
	var res Orders
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
