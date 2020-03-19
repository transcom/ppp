// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

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

// CreateMTOServiceItemHandlerFunc turns a function with the right signature into a create m t o service item handler
type CreateMTOServiceItemHandlerFunc func(CreateMTOServiceItemParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateMTOServiceItemHandlerFunc) Handle(params CreateMTOServiceItemParams) middleware.Responder {
	return fn(params)
}

// CreateMTOServiceItemHandler interface for that can handle valid create m t o service item params
type CreateMTOServiceItemHandler interface {
	Handle(CreateMTOServiceItemParams) middleware.Responder
}

// NewCreateMTOServiceItem creates a new http.Handler for the create m t o service item operation
func NewCreateMTOServiceItem(ctx *middleware.Context, handler CreateMTOServiceItemHandler) *CreateMTOServiceItem {
	return &CreateMTOServiceItem{Context: ctx, Handler: handler}
}

/*CreateMTOServiceItem swagger:route POST /move_task_orders/{moveTaskOrderID}/mto_service_items mtoServiceItem createMTOServiceItem

Creates a service item for a move order by id

Creates a service item for a move order by id

*/
type CreateMTOServiceItem struct {
	Context *middleware.Context
	Handler CreateMTOServiceItemHandler
}

func (o *CreateMTOServiceItem) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateMTOServiceItemParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateMTOServiceItemBody create m t o service item body
// swagger:model CreateMTOServiceItemBody
type CreateMTOServiceItemBody struct {

	// mto shipment ID
	// Required: true
	// Format: uuid
	MtoShipmentID *strfmt.UUID `json:"mtoShipmentID"`

	// re service ID
	// Required: true
	// Format: uuid
	ReServiceID *strfmt.UUID `json:"reServiceID"`
}

// Validate validates this create m t o service item body
func (o *CreateMTOServiceItemBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMtoShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReServiceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateMTOServiceItemBody) validateMtoShipmentID(formats strfmt.Registry) error {

	if err := validate.Required("createMTOServiceItemBody"+"."+"mtoShipmentID", "body", o.MtoShipmentID); err != nil {
		return err
	}

	if err := validate.FormatOf("createMTOServiceItemBody"+"."+"mtoShipmentID", "body", "uuid", o.MtoShipmentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *CreateMTOServiceItemBody) validateReServiceID(formats strfmt.Registry) error {

	if err := validate.Required("createMTOServiceItemBody"+"."+"reServiceID", "body", o.ReServiceID); err != nil {
		return err
	}

	if err := validate.FormatOf("createMTOServiceItemBody"+"."+"reServiceID", "body", "uuid", o.ReServiceID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateMTOServiceItemBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateMTOServiceItemBody) UnmarshalBinary(b []byte) error {
	var res CreateMTOServiceItemBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
