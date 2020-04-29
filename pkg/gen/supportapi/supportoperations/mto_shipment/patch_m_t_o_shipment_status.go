// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchMTOShipmentStatusHandlerFunc turns a function with the right signature into a patch m t o shipment status handler
type PatchMTOShipmentStatusHandlerFunc func(PatchMTOShipmentStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchMTOShipmentStatusHandlerFunc) Handle(params PatchMTOShipmentStatusParams) middleware.Responder {
	return fn(params)
}

// PatchMTOShipmentStatusHandler interface for that can handle valid patch m t o shipment status params
type PatchMTOShipmentStatusHandler interface {
	Handle(PatchMTOShipmentStatusParams) middleware.Responder
}

// NewPatchMTOShipmentStatus creates a new http.Handler for the patch m t o shipment status operation
func NewPatchMTOShipmentStatus(ctx *middleware.Context, handler PatchMTOShipmentStatusHandler) *PatchMTOShipmentStatus {
	return &PatchMTOShipmentStatus{Context: ctx, Handler: handler}
}

/*PatchMTOShipmentStatus swagger:route PATCH /mto-shipments/{mtoShipmentID}/status mtoShipment patchMTOShipmentStatus

Updates a shipment's status.

Updates a shipment's status to APPROVED or REJECTED for the purpose of testing the Prime API. If APPROVED, `rejectionReason` should be blank and any value passed through the body will be ignored. If REJECTED, a value in `rejectionReason` is required.


*/
type PatchMTOShipmentStatus struct {
	Context *middleware.Context
	Handler PatchMTOShipmentStatusHandler
}

func (o *PatchMTOShipmentStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchMTOShipmentStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
