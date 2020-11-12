// Code generated by go-swagger; DO NOT EDIT.

package office_users

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

	"github.com/transcom/mymove/pkg/gen/adminmessages"
)

// NewUpdateOfficeUserParams creates a new UpdateOfficeUserParams object
// no default values defined in spec.
func NewUpdateOfficeUserParams() UpdateOfficeUserParams {

	return UpdateOfficeUserParams{}
}

// UpdateOfficeUserParams contains all the bound params for the update office user operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateOfficeUser
type UpdateOfficeUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Office user information
	  Required: true
	  In: body
	*/
	OfficeUser *adminmessages.OfficeUserUpdatePayload
	/*
	  Required: true
	  In: path
	*/
	OfficeUserID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateOfficeUserParams() beforehand.
func (o *UpdateOfficeUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body adminmessages.OfficeUserUpdatePayload
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("officeUser", "body", ""))
			} else {
				res = append(res, errors.NewParseError("officeUser", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.OfficeUser = &body
			}
		}
	} else {
		res = append(res, errors.Required("officeUser", "body", ""))
	}
	rOfficeUserID, rhkOfficeUserID, _ := route.Params.GetOK("officeUserId")
	if err := o.bindOfficeUserID(rOfficeUserID, rhkOfficeUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindOfficeUserID binds and validates parameter OfficeUserID from path.
func (o *UpdateOfficeUserParams) bindOfficeUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("officeUserId", "path", "strfmt.UUID", raw)
	}
	o.OfficeUserID = *(value.(*strfmt.UUID))

	if err := o.validateOfficeUserID(formats); err != nil {
		return err
	}

	return nil
}

// validateOfficeUserID carries on validations for parameter OfficeUserID
func (o *UpdateOfficeUserParams) validateOfficeUserID(formats strfmt.Registry) error {

	if err := validate.FormatOf("officeUserId", "path", "uuid", o.OfficeUserID.String(), formats); err != nil {
		return err
	}
	return nil
}
