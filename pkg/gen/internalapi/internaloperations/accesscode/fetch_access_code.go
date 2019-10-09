// Code generated by go-swagger; DO NOT EDIT.

package accesscode

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FetchAccessCodeHandlerFunc turns a function with the right signature into a fetch access code handler
type FetchAccessCodeHandlerFunc func(FetchAccessCodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FetchAccessCodeHandlerFunc) Handle(params FetchAccessCodeParams) middleware.Responder {
	return fn(params)
}

// FetchAccessCodeHandler interface for that can handle valid fetch access code params
type FetchAccessCodeHandler interface {
	Handle(FetchAccessCodeParams) middleware.Responder
}

// NewFetchAccessCode creates a new http.Handler for the fetch access code operation
func NewFetchAccessCode(ctx *middleware.Context, handler FetchAccessCodeHandler) *FetchAccessCode {
	return &FetchAccessCode{Context: ctx, Handler: handler}
}

/*FetchAccessCode swagger:route GET /access_codes accesscode fetchAccessCode

Fetches an access code

Fetches the access code for a service member.

*/
type FetchAccessCode struct {
	Context *middleware.Context
	Handler FetchAccessCodeHandler
}

func (o *FetchAccessCode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFetchAccessCodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
