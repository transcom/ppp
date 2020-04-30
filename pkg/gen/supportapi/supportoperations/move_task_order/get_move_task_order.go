// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetMoveTaskOrderHandlerFunc turns a function with the right signature into a get move task order handler
type GetMoveTaskOrderHandlerFunc func(GetMoveTaskOrderParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMoveTaskOrderHandlerFunc) Handle(params GetMoveTaskOrderParams) middleware.Responder {
	return fn(params)
}

// GetMoveTaskOrderHandler interface for that can handle valid get move task order params
type GetMoveTaskOrderHandler interface {
	Handle(GetMoveTaskOrderParams) middleware.Responder
}

// NewGetMoveTaskOrder creates a new http.Handler for the get move task order operation
func NewGetMoveTaskOrder(ctx *middleware.Context, handler GetMoveTaskOrderHandler) *GetMoveTaskOrder {
	return &GetMoveTaskOrder{Context: ctx, Handler: handler}
}

/*GetMoveTaskOrder swagger:route GET /move-task-orders/{moveTaskOrderID} moveTaskOrder getMoveTaskOrder

Gets a move task order by ID

Gets an individual move task order by ID.

*/
type GetMoveTaskOrder struct {
	Context *middleware.Context
	Handler GetMoveTaskOrderHandler
}

func (o *GetMoveTaskOrder) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMoveTaskOrderParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
