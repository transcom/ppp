// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ShowPPMIncentiveHandlerFunc turns a function with the right signature into a show p p m incentive handler
type ShowPPMIncentiveHandlerFunc func(ShowPPMIncentiveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ShowPPMIncentiveHandlerFunc) Handle(params ShowPPMIncentiveParams) middleware.Responder {
	return fn(params)
}

// ShowPPMIncentiveHandler interface for that can handle valid show p p m incentive params
type ShowPPMIncentiveHandler interface {
	Handle(ShowPPMIncentiveParams) middleware.Responder
}

// NewShowPPMIncentive creates a new http.Handler for the show p p m incentive operation
func NewShowPPMIncentive(ctx *middleware.Context, handler ShowPPMIncentiveHandler) *ShowPPMIncentive {
	return &ShowPPMIncentive{Context: ctx, Handler: handler}
}

/*ShowPPMIncentive swagger:route GET /personally_procured_moves/incentive ppm showPPMIncentive

Return a PPM incentive value

Calculates incentive for a PPM move (excluding SIT)

*/
type ShowPPMIncentive struct {
	Context *middleware.Context
	Handler ShowPPMIncentiveHandler
}

func (o *ShowPPMIncentive) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewShowPPMIncentiveParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
