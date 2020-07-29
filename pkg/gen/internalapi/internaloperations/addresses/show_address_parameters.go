// Code generated by go-swagger; DO NOT EDIT.

package addresses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShowAddressParams creates a new ShowAddressParams object
// no default values defined in spec.
func NewShowAddressParams() ShowAddressParams {

	return ShowAddressParams{}
}

// ShowAddressParams contains all the bound params for the show address operation
// typically these are obtained from a http.Request
//
// swagger:parameters showAddress
type ShowAddressParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*UUID of the address to return
	  Required: true
	  In: path
	*/
	AddressID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewShowAddressParams() beforehand.
func (o *ShowAddressParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAddressID, rhkAddressID, _ := route.Params.GetOK("addressId")
	if err := o.bindAddressID(rAddressID, rhkAddressID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAddressID binds and validates parameter AddressID from path.
func (o *ShowAddressParams) bindAddressID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("addressId", "path", "strfmt.UUID", raw)
	}
	o.AddressID = *(value.(*strfmt.UUID))

	if err := o.validateAddressID(formats); err != nil {
		return err
	}

	return nil
}

// validateAddressID carries on validations for parameter AddressID
func (o *ShowAddressParams) validateAddressID(formats strfmt.Registry) error {

	if err := validate.FormatOf("addressId", "path", "uuid", o.AddressID.String(), formats); err != nil {
		return err
	}
	return nil
}
