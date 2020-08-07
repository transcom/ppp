// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMTOServiceItemParams creates a new GetMTOServiceItemParams object
// no default values defined in spec.
func NewGetMTOServiceItemParams() GetMTOServiceItemParams {

	return GetMTOServiceItemParams{}
}

// GetMTOServiceItemParams contains all the bound params for the get m t o service item operation
// typically these are obtained from a http.Request
//
// swagger:parameters getMTOServiceItem
type GetMTOServiceItemParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of move order to use
	  Required: true
	  In: path
	*/
	MoveTaskOrderID string
	/*ID of line item to use
	  Required: true
	  In: path
	*/
	MtoServiceItemID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetMTOServiceItemParams() beforehand.
func (o *GetMTOServiceItemParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rMoveTaskOrderID, rhkMoveTaskOrderID, _ := route.Params.GetOK("moveTaskOrderID")
	if err := o.bindMoveTaskOrderID(rMoveTaskOrderID, rhkMoveTaskOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	rMtoServiceItemID, rhkMtoServiceItemID, _ := route.Params.GetOK("mtoServiceItemID")
	if err := o.bindMtoServiceItemID(rMtoServiceItemID, rhkMtoServiceItemID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMoveTaskOrderID binds and validates parameter MoveTaskOrderID from path.
func (o *GetMTOServiceItemParams) bindMoveTaskOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.MoveTaskOrderID = raw

	return nil
}

// bindMtoServiceItemID binds and validates parameter MtoServiceItemID from path.
func (o *GetMTOServiceItemParams) bindMtoServiceItemID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.MtoServiceItemID = raw

	return nil
}
