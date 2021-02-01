// Code generated by go-swagger; DO NOT EDIT.

package move

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewUpdateMoveParams creates a new UpdateMoveParams object
// no default values defined in spec.
func NewUpdateMoveParams() UpdateMoveParams {

	return UpdateMoveParams{}
}

// UpdateMoveParams contains all the bound params for the update move operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateMove
type UpdateMoveParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Move information
	  Required: true
	  In: body
	*/
	Move UpdateMoveBody
	/*
	  Required: true
	  In: path
	*/
	MoveID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateMoveParams() beforehand.
func (o *UpdateMoveParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body UpdateMoveBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("move", "body", ""))
			} else {
				res = append(res, errors.NewParseError("move", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Move = body
			}
		}
	} else {
		res = append(res, errors.Required("move", "body", ""))
	}
	rMoveID, rhkMoveID, _ := route.Params.GetOK("moveID")
	if err := o.bindMoveID(rMoveID, rhkMoveID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMoveID binds and validates parameter MoveID from path.
func (o *UpdateMoveParams) bindMoveID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("moveID", "path", "strfmt.UUID", raw)
	}
	o.MoveID = *(value.(*strfmt.UUID))

	if err := o.validateMoveID(formats); err != nil {
		return err
	}

	return nil
}

// validateMoveID carries on validations for parameter MoveID
func (o *UpdateMoveParams) validateMoveID(formats strfmt.Registry) error {

	if err := validate.FormatOf("moveID", "path", "uuid", o.MoveID.String(), formats); err != nil {
		return err
	}
	return nil
}