// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

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

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// NewUpdateMoveTaskOrderStatusParams creates a new UpdateMoveTaskOrderStatusParams object
// no default values defined in spec.
func NewUpdateMoveTaskOrderStatusParams() UpdateMoveTaskOrderStatusParams {

	return UpdateMoveTaskOrderStatusParams{}
}

// UpdateMoveTaskOrderStatusParams contains all the bound params for the update move task order status operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateMoveTaskOrderStatus
type UpdateMoveTaskOrderStatusParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	IfMatch string
	/*ID of move order to use
	  Required: true
	  In: path
	*/
	MoveTaskOrderID string
	/*
	  Required: true
	  In: body
	*/
	ServiceItemCodes ghcmessages.MTOApprovalServiceItemCodes
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateMoveTaskOrderStatusParams() beforehand.
func (o *UpdateMoveTaskOrderStatusParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindIfMatch(r.Header[http.CanonicalHeaderKey("If-Match")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rMoveTaskOrderID, rhkMoveTaskOrderID, _ := route.Params.GetOK("moveTaskOrderID")
	if err := o.bindMoveTaskOrderID(rMoveTaskOrderID, rhkMoveTaskOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body ghcmessages.MTOApprovalServiceItemCodes
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("serviceItemCodes", "body"))
			} else {
				res = append(res, errors.NewParseError("serviceItemCodes", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ServiceItemCodes = body
			}
		}
	} else {
		res = append(res, errors.Required("serviceItemCodes", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIfMatch binds and validates parameter IfMatch from header.
func (o *UpdateMoveTaskOrderStatusParams) bindIfMatch(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("If-Match", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("If-Match", "header", raw); err != nil {
		return err
	}

	o.IfMatch = raw

	return nil
}

// bindMoveTaskOrderID binds and validates parameter MoveTaskOrderID from path.
func (o *UpdateMoveTaskOrderStatusParams) bindMoveTaskOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.MoveTaskOrderID = raw

	return nil
}
