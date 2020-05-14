// Code generated by go-swagger; DO NOT EDIT.

package payment_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreatePaymentRequestHandlerFunc turns a function with the right signature into a create payment request handler
type CreatePaymentRequestHandlerFunc func(CreatePaymentRequestParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreatePaymentRequestHandlerFunc) Handle(params CreatePaymentRequestParams) middleware.Responder {
	return fn(params)
}

// CreatePaymentRequestHandler interface for that can handle valid create payment request params
type CreatePaymentRequestHandler interface {
	Handle(CreatePaymentRequestParams) middleware.Responder
}

// NewCreatePaymentRequest creates a new http.Handler for the create payment request operation
func NewCreatePaymentRequest(ctx *middleware.Context, handler CreatePaymentRequestHandler) *CreatePaymentRequest {
	return &CreatePaymentRequest{Context: ctx, Handler: handler}
}

/*CreatePaymentRequest swagger:route POST /payment-requests paymentRequests createPaymentRequest

createPaymentRequest

Creates a new instance of a paymentRequest.
A newly created payment request is assigned the status `PENDING`.
A move task order can have multiple payment requests, and
a final payment request can be marked using boolean `isFinal`.


*/
type CreatePaymentRequest struct {
	Context *middleware.Context
	Handler CreatePaymentRequestHandler
}

func (o *CreatePaymentRequest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreatePaymentRequestParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
