// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// PostWebhookNotifyHandlerFunc turns a function with the right signature into a post webhook notify handler
type PostWebhookNotifyHandlerFunc func(PostWebhookNotifyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostWebhookNotifyHandlerFunc) Handle(params PostWebhookNotifyParams) middleware.Responder {
	return fn(params)
}

// PostWebhookNotifyHandler interface for that can handle valid post webhook notify params
type PostWebhookNotifyHandler interface {
	Handle(PostWebhookNotifyParams) middleware.Responder
}

// NewPostWebhookNotify creates a new http.Handler for the post webhook notify operation
func NewPostWebhookNotify(ctx *middleware.Context, handler PostWebhookNotifyHandler) *PostWebhookNotify {
	return &PostWebhookNotify{Context: ctx, Handler: handler}
}

/*PostWebhookNotify swagger:route POST /webhook-notify webhook postWebhookNotify

Test endpoint for sending messages via webhook

This endpoint represents the receiving server, The Prime, in our webhook-client testing workflow. The `webhook-client` is responsible for retrieving messages from the webhook_notifications table and sending them to the Prime (this endpoint in our testing case) via an mTLS connection.


*/
type PostWebhookNotify struct {
	Context *middleware.Context
	Handler PostWebhookNotifyHandler
}

func (o *PostWebhookNotify) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostWebhookNotifyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostWebhookNotifyBody post webhook notify body
// swagger:model PostWebhookNotifyBody
type PostWebhookNotifyBody struct {

	// Message sent
	// Required: true
	Message string `json:"message"`
}

// Validate validates this post webhook notify body
func (o *PostWebhookNotifyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWebhookNotifyBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.RequiredString("message"+"."+"message", "body", string(o.Message)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWebhookNotifyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWebhookNotifyBody) UnmarshalBinary(b []byte) error {
	var res PostWebhookNotifyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostWebhookNotifyOKBody post webhook notify o k body
// swagger:model PostWebhookNotifyOKBody
type PostWebhookNotifyOKBody struct {

	// Message sent
	Message string `json:"message,omitempty"`
}

// Validate validates this post webhook notify o k body
func (o *PostWebhookNotifyOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWebhookNotifyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWebhookNotifyOKBody) UnmarshalBinary(b []byte) error {
	var res PostWebhookNotifyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
