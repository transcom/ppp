// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOShipment m t o shipment
// swagger:model MTOShipment
type MTOShipment struct {

	// agents
	Agents MTOAgents `json:"agents,omitempty"`

	// created at
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// customer remarks
	// Read Only: true
	CustomerRemarks *string `json:"customerRemarks,omitempty"`

	// destination address
	DestinationAddress *Address `json:"destinationAddress,omitempty"`

	// e tag
	ETag string `json:"eTag,omitempty"`

	// id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// move task order ID
	// Read Only: true
	// Format: uuid
	MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

	// pickup address
	PickupAddress *Address `json:"pickupAddress,omitempty"`

	// requested delivery date
	// Read Only: true
	// Format: date
	RequestedDeliveryDate strfmt.Date `json:"requestedDeliveryDate,omitempty"`

	// requested pickup date
	// Read Only: true
	// Format: date
	RequestedPickupDate strfmt.Date `json:"requestedPickupDate,omitempty"`

	// secondary delivery address
	SecondaryDeliveryAddress *Address `json:"secondaryDeliveryAddress,omitempty"`

	// secondary pickup address
	SecondaryPickupAddress *Address `json:"secondaryPickupAddress,omitempty"`

	// shipment type
	ShipmentType MTOShipmentType `json:"shipmentType,omitempty"`

	// status
	Status MTOShipmentStatus `json:"status,omitempty"`

	// updated at
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this m t o shipment
func (m *MTOShipment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTaskOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickupAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedPickupDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryDeliveryAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryPickupAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *MTOShipment) validateAgents(formats strfmt.Registry) error {

	if swag.IsZero(m.Agents) { // not required
		return nil
	}

	if err := m.Agents.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("agents")
		}
		return err
	}

	return nil
}

func (m *MTOShipment) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateDestinationAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationAddress) { // not required
		return nil
	}

	if m.DestinationAddress != nil {
		if err := m.DestinationAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MTOShipment) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validatePickupAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.PickupAddress) { // not required
		return nil
	}

	if m.PickupAddress != nil {
		if err := m.PickupAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pickupAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MTOShipment) validateRequestedDeliveryDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedDeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedDeliveryDate", "body", "date", m.RequestedDeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateRequestedPickupDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedPickupDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedPickupDate", "body", "date", m.RequestedPickupDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateSecondaryDeliveryAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryDeliveryAddress) { // not required
		return nil
	}

	if m.SecondaryDeliveryAddress != nil {
		if err := m.SecondaryDeliveryAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryDeliveryAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MTOShipment) validateSecondaryPickupAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryPickupAddress) { // not required
		return nil
	}

	if m.SecondaryPickupAddress != nil {
		if err := m.SecondaryPickupAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryPickupAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MTOShipment) validateShipmentType(formats strfmt.Registry) error {

	if swag.IsZero(m.ShipmentType) { // not required
		return nil
	}

	if err := m.ShipmentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shipmentType")
		}
		return err
	}

	return nil
}

func (m *MTOShipment) validateStatus(formats strfmt.Registry) error {

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

func (m *MTOShipment) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOShipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOShipment) UnmarshalBinary(b []byte) error {
	var res MTOShipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
