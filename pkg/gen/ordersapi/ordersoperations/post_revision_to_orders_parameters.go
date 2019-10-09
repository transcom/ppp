// Code generated by go-swagger; DO NOT EDIT.

package ordersoperations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	ordersmessages "github.com/transcom/mymove/pkg/gen/ordersmessages"
)

// NewPostRevisionToOrdersParams creates a new PostRevisionToOrdersParams object
// no default values defined in spec.
func NewPostRevisionToOrdersParams() PostRevisionToOrdersParams {

	return PostRevisionToOrdersParams{}
}

// PostRevisionToOrdersParams contains all the bound params for the post revision to orders operation
// typically these are obtained from a http.Request
//
// swagger:parameters postRevisionToOrders
type PostRevisionToOrdersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Revision *ordersmessages.Revision
	/*UUID of the orders to amend
	  Required: true
	  In: path
	*/
	UUID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostRevisionToOrdersParams() beforehand.
func (o *PostRevisionToOrdersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body ordersmessages.Revision
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("revision", "body"))
			} else {
				res = append(res, errors.NewParseError("revision", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Revision = &body
			}
		}
	} else {
		res = append(res, errors.Required("revision", "body"))
	}
	rUUID, rhkUUID, _ := route.Params.GetOK("uuid")
	if err := o.bindUUID(rUUID, rhkUUID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUUID binds and validates parameter UUID from path.
func (o *PostRevisionToOrdersParams) bindUUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("uuid", "path", "strfmt.UUID", raw)
	}
	o.UUID = *(value.(*strfmt.UUID))

	if err := o.validateUUID(formats); err != nil {
		return err
	}

	return nil
}

// validateUUID carries on validations for parameter UUID
func (o *PostRevisionToOrdersParams) validateUUID(formats strfmt.Registry) error {

	if err := validate.FormatOf("uuid", "path", "uuid", o.UUID.String(), formats); err != nil {
		return err
	}
	return nil
}
