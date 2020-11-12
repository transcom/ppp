// Code generated by go-swagger; DO NOT EDIT.

package upload

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetUploadParams creates a new GetUploadParams object
// no default values defined in spec.
func NewGetUploadParams() GetUploadParams {

	return GetUploadParams{}
}

// GetUploadParams contains all the bound params for the get upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters getUpload
type GetUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	UploadID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetUploadParams() beforehand.
func (o *GetUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rUploadID, rhkUploadID, _ := route.Params.GetOK("uploadId")
	if err := o.bindUploadID(rUploadID, rhkUploadID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUploadID binds and validates parameter UploadID from path.
func (o *GetUploadParams) bindUploadID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("uploadId", "path", "strfmt.UUID", raw)
	}
	o.UploadID = *(value.(*strfmt.UUID))

	if err := o.validateUploadID(formats); err != nil {
		return err
	}

	return nil
}

// validateUploadID carries on validations for parameter UploadID
func (o *GetUploadParams) validateUploadID(formats strfmt.Registry) error {

	if err := validate.FormatOf("uploadId", "path", "uuid", o.UploadID.String(), formats); err != nil {
		return err
	}
	return nil
}
