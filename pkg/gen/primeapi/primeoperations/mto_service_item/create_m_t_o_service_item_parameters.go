// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// NewCreateMTOServiceItemParams creates a new CreateMTOServiceItemParams object
// no default values defined in spec.
func NewCreateMTOServiceItemParams() CreateMTOServiceItemParams {

	return CreateMTOServiceItemParams{}
}

// CreateMTOServiceItemParams contains all the bound params for the create m t o service item operation
// typically these are obtained from a http.Request
//
// swagger:parameters createMTOServiceItem
type CreateMTOServiceItemParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body primemessages.MTOServiceItem
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateMTOServiceItemParams() beforehand.
func (o *CreateMTOServiceItemParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		body, err := primemessages.UnmarshalMTOServiceItem(r.Body, route.Consumer)
		if err != nil {
			res = append(res, err)
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
