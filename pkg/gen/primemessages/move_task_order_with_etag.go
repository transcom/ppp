// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MoveTaskOrderWithEtag move task order with etag
// swagger:model MoveTaskOrderWithEtag
type MoveTaskOrderWithEtag struct {
	MoveTaskOrder

	MoveTaskOrderWithEtagAllOf1

	// e tag
	ETag string `json:"eTag,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MoveTaskOrderWithEtag) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoveTaskOrder
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoveTaskOrder = aO0

	// AO1
	var aO1 MoveTaskOrderWithEtagAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.MoveTaskOrderWithEtagAllOf1 = aO1

	// now for regular properties
	var propsMoveTaskOrderWithEtag struct {
		ETag string `json:"eTag,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsMoveTaskOrderWithEtag); err != nil {
		return err
	}
	m.ETag = propsMoveTaskOrderWithEtag.ETag

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MoveTaskOrderWithEtag) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoveTaskOrder)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.MoveTaskOrderWithEtagAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	// now for regular properties
	var propsMoveTaskOrderWithEtag struct {
		ETag string `json:"eTag,omitempty"`
	}
	propsMoveTaskOrderWithEtag.ETag = m.ETag

	jsonDataPropsMoveTaskOrderWithEtag, errMoveTaskOrderWithEtag := swag.WriteJSON(propsMoveTaskOrderWithEtag)
	if errMoveTaskOrderWithEtag != nil {
		return nil, errMoveTaskOrderWithEtag
	}
	_parts = append(_parts, jsonDataPropsMoveTaskOrderWithEtag)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this move task order with etag
func (m *MoveTaskOrderWithEtag) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoveTaskOrder
	if err := m.MoveTaskOrder.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with MoveTaskOrderWithEtagAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MoveTaskOrderWithEtag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveTaskOrderWithEtag) UnmarshalBinary(b []byte) error {
	var res MoveTaskOrderWithEtag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MoveTaskOrderWithEtagAllOf1 move task order with etag all of1
// swagger:model MoveTaskOrderWithEtagAllOf1
type MoveTaskOrderWithEtagAllOf1 interface{}
