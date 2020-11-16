// Code generated by go-swagger; DO NOT EDIT.

package move_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewListMoveTaskOrdersParams creates a new ListMoveTaskOrdersParams object
// no default values defined in spec.
func NewListMoveTaskOrdersParams() ListMoveTaskOrdersParams {

	return ListMoveTaskOrdersParams{}
}

// ListMoveTaskOrdersParams contains all the bound params for the list move task orders operation
// typically these are obtained from a http.Request
//
// swagger:parameters listMoveTaskOrders
type ListMoveTaskOrdersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of move order to use
	  Required: true
	  In: path
	*/
	MoveOrderID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListMoveTaskOrdersParams() beforehand.
func (o *ListMoveTaskOrdersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rMoveOrderID, rhkMoveOrderID, _ := route.Params.GetOK("moveOrderID")
	if err := o.bindMoveOrderID(rMoveOrderID, rhkMoveOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMoveOrderID binds and validates parameter MoveOrderID from path.
func (o *ListMoveTaskOrdersParams) bindMoveOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("moveOrderID", "path", "strfmt.UUID", raw)
	}
	o.MoveOrderID = *(value.(*strfmt.UUID))

	if err := o.validateMoveOrderID(formats); err != nil {
		return err
	}

	return nil
}

// validateMoveOrderID carries on validations for parameter MoveOrderID
func (o *ListMoveTaskOrdersParams) validateMoveOrderID(formats strfmt.Registry) error {

	if err := validate.FormatOf("moveOrderID", "path", "uuid", o.MoveOrderID.String(), formats); err != nil {
		return err
	}
	return nil
}
