// Code generated by go-swagger; DO NOT EDIT.

package office_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// IndexOfficeUsersHandlerFunc turns a function with the right signature into a index office users handler
type IndexOfficeUsersHandlerFunc func(IndexOfficeUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn IndexOfficeUsersHandlerFunc) Handle(params IndexOfficeUsersParams) middleware.Responder {
	return fn(params)
}

// IndexOfficeUsersHandler interface for that can handle valid index office users params
type IndexOfficeUsersHandler interface {
	Handle(IndexOfficeUsersParams) middleware.Responder
}

// NewIndexOfficeUsers creates a new http.Handler for the index office users operation
func NewIndexOfficeUsers(ctx *middleware.Context, handler IndexOfficeUsersHandler) *IndexOfficeUsers {
	return &IndexOfficeUsers{Context: ctx, Handler: handler}
}

/*IndexOfficeUsers swagger:route GET /office_users office_users indexOfficeUsers

List office users

Returns a list of office users

*/
type IndexOfficeUsers struct {
	Context *middleware.Context
	Handler IndexOfficeUsersHandler
}

func (o *IndexOfficeUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewIndexOfficeUsersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
