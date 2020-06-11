// Code generated by go-swagger; DO NOT EDIT.

package gex

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	adminmessages "github.com/transcom/mymove/pkg/gen/adminmessages"
)

// NewSendGexRequestParams creates a new SendGexRequestParams object
// no default values defined in spec.
func NewSendGexRequestParams() SendGexRequestParams {

	return SendGexRequestParams{}
}

// SendGexRequestParams contains all the bound params for the send gex request operation
// typically these are obtained from a http.Request
//
// swagger:parameters sendGexRequest
type SendGexRequestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	SendGexRequest *adminmessages.SendGexRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSendGexRequestParams() beforehand.
func (o *SendGexRequestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body adminmessages.SendGexRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("sendGexRequest", "body"))
			} else {
				res = append(res, errors.NewParseError("sendGexRequest", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.SendGexRequest = &body
			}
		}
	} else {
		res = append(res, errors.Required("sendGexRequest", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
