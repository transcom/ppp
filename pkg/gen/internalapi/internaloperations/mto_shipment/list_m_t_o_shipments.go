// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListMTOShipmentsHandlerFunc turns a function with the right signature into a list m t o shipments handler
type ListMTOShipmentsHandlerFunc func(ListMTOShipmentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListMTOShipmentsHandlerFunc) Handle(params ListMTOShipmentsParams) middleware.Responder {
	return fn(params)
}

// ListMTOShipmentsHandler interface for that can handle valid list m t o shipments params
type ListMTOShipmentsHandler interface {
	Handle(ListMTOShipmentsParams) middleware.Responder
}

// NewListMTOShipments creates a new http.Handler for the list m t o shipments operation
func NewListMTOShipments(ctx *middleware.Context, handler ListMTOShipmentsHandler) *ListMTOShipments {
	return &ListMTOShipments{Context: ctx, Handler: handler}
}

/*ListMTOShipments swagger:route GET /moves/{moveTaskOrderID}/mto_shipments mtoShipment listMTOShipments

Gets all shipments for a move task order

Gets all MTO shipments for the specified Move Task Order.


*/
type ListMTOShipments struct {
	Context *middleware.Context
	Handler ListMTOShipmentsHandler
}

func (o *ListMTOShipments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListMTOShipmentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
