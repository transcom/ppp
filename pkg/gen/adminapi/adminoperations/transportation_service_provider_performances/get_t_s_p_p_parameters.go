// Code generated by go-swagger; DO NOT EDIT.

package transportation_service_provider_performances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetTSPPParams creates a new GetTSPPParams object
// no default values defined in spec.
func NewGetTSPPParams() GetTSPPParams {

	return GetTSPPParams{}
}

// GetTSPPParams contains all the bound params for the get t s p p operation
// typically these are obtained from a http.Request
//
// swagger:parameters getTSPP
type GetTSPPParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	TsppID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetTSPPParams() beforehand.
func (o *GetTSPPParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rTsppID, rhkTsppID, _ := route.Params.GetOK("tsppId")
	if err := o.bindTsppID(rTsppID, rhkTsppID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTsppID binds and validates parameter TsppID from path.
func (o *GetTSPPParams) bindTsppID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("tsppId", "path", "strfmt.UUID", raw)
	}
	o.TsppID = *(value.(*strfmt.UUID))

	if err := o.validateTsppID(formats); err != nil {
		return err
	}

	return nil
}

// validateTsppID carries on validations for parameter TsppID
func (o *GetTSPPParams) validateTsppID(formats strfmt.Registry) error {

	if err := validate.FormatOf("tsppId", "path", "uuid", o.TsppID.String(), formats); err != nil {
		return err
	}
	return nil
}
