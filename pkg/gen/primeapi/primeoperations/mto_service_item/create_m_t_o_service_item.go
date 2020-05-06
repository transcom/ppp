// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
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

/*CreateMTOServiceItem swagger:route POST /mto-service-items mtoServiceItem createMTOServiceItem

Creates MTO service items that is added to a Move Task Order and MTO Shipment

Creates a new instance of mtoServiceItem, which come from the list of services that can be provided. Upon creation these items are associated with a Move Task Order and an MTO Shipment. Must include UUIDs for the MTO and MTO Shipment connected to this service item.


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
