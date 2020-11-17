// Code generated by go-swagger; DO NOT EDIT.

package documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateDocumentHandlerFunc turns a function with the right signature into a create document handler
type CreateDocumentHandlerFunc func(CreateDocumentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateDocumentHandlerFunc) Handle(params CreateDocumentParams) middleware.Responder {
	return fn(params)
}

// CreateDocumentHandler interface for that can handle valid create document params
type CreateDocumentHandler interface {
	Handle(CreateDocumentParams) middleware.Responder
}

// NewCreateDocument creates a new http.Handler for the create document operation
func NewCreateDocument(ctx *middleware.Context, handler CreateDocumentHandler) *CreateDocument {
	return &CreateDocument{Context: ctx, Handler: handler}
}

/*CreateDocument swagger:route POST /documents documents createDocument

Create a new document

Documents represent a physical artifact such as a scanned document or a PDF file

*/
type CreateDocument struct {
	Context *middleware.Context
	Handler CreateDocumentHandler
}

func (o *CreateDocument) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateDocumentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
