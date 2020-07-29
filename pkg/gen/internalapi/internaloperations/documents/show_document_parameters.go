// Code generated by go-swagger; DO NOT EDIT.

package documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShowDocumentParams creates a new ShowDocumentParams object
// no default values defined in spec.
func NewShowDocumentParams() ShowDocumentParams {

	return ShowDocumentParams{}
}

// ShowDocumentParams contains all the bound params for the show document operation
// typically these are obtained from a http.Request
//
// swagger:parameters showDocument
type ShowDocumentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*UUID of the document to return
	  Required: true
	  In: path
	*/
	DocumentID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewShowDocumentParams() beforehand.
func (o *ShowDocumentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDocumentID, rhkDocumentID, _ := route.Params.GetOK("documentId")
	if err := o.bindDocumentID(rDocumentID, rhkDocumentID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDocumentID binds and validates parameter DocumentID from path.
func (o *ShowDocumentParams) bindDocumentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("documentId", "path", "strfmt.UUID", raw)
	}
	o.DocumentID = *(value.(*strfmt.UUID))

	if err := o.validateDocumentID(formats); err != nil {
		return err
	}

	return nil
}

// validateDocumentID carries on validations for parameter DocumentID
func (o *ShowDocumentParams) validateDocumentID(formats strfmt.Registry) error {

	if err := validate.FormatOf("documentId", "path", "uuid", o.DocumentID.String(), formats); err != nil {
		return err
	}
	return nil
}
