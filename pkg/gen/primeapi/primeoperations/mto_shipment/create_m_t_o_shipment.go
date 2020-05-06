// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateMTOShipmentHandlerFunc turns a function with the right signature into a create m t o shipment handler
type CreateMTOShipmentHandlerFunc func(CreateMTOShipmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateMTOShipmentHandlerFunc) Handle(params CreateMTOShipmentParams) middleware.Responder {
	return fn(params)
}

// CreateMTOShipmentHandler interface for that can handle valid create m t o shipment params
type CreateMTOShipmentHandler interface {
	Handle(CreateMTOShipmentParams) middleware.Responder
}

// NewCreateMTOShipment creates a new http.Handler for the create m t o shipment operation
func NewCreateMTOShipment(ctx *middleware.Context, handler CreateMTOShipmentHandler) *CreateMTOShipment {
	return &CreateMTOShipment{Context: ctx, Handler: handler}
}

/*CreateMTOShipment swagger:route POST /mto-shipments mtoShipment createMTOShipment

Creates MTO shipment

Creates a MTO shipment for the specified Move Task Order.
Required fields include:
* Shipment Type
* Customer requested pick-up date
* Pick-up Address
* Delivery Address
* Releasing / Receiving agents

Optional fields include:
* Customer Remarks
* Releasing / Receiving agents
* An array of optional accessorial service item codes


*/
type CreateMTOShipment struct {
	Context *middleware.Context
	Handler CreateMTOShipmentHandler
}

func (o *CreateMTOShipment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateMTOShipmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
