// Code generated by go-swagger; DO NOT EDIT.

package queues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetPaymentRequestsQueueHandlerFunc turns a function with the right signature into a get payment requests queue handler
type GetPaymentRequestsQueueHandlerFunc func(GetPaymentRequestsQueueParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPaymentRequestsQueueHandlerFunc) Handle(params GetPaymentRequestsQueueParams) middleware.Responder {
	return fn(params)
}

// GetPaymentRequestsQueueHandler interface for that can handle valid get payment requests queue params
type GetPaymentRequestsQueueHandler interface {
	Handle(GetPaymentRequestsQueueParams) middleware.Responder
}

// NewGetPaymentRequestsQueue creates a new http.Handler for the get payment requests queue operation
func NewGetPaymentRequestsQueue(ctx *middleware.Context, handler GetPaymentRequestsQueueHandler) *GetPaymentRequestsQueue {
	return &GetPaymentRequestsQueue{Context: ctx, Handler: handler}
}

/*GetPaymentRequestsQueue swagger:route GET /queues/payment-requests queues getPaymentRequestsQueue

Gets queued list of all payment requests by GBLOC origin

An office TIO user will be assigned a transportation office that will determine which payment requests are displayed in their queue based on the origin duty station.


*/
type GetPaymentRequestsQueue struct {
	Context *middleware.Context
	Handler GetPaymentRequestsQueueHandler
}

func (o *GetPaymentRequestsQueue) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPaymentRequestsQueueParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
