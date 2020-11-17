// Code generated by go-swagger; DO NOT EDIT.

package move_docs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteMoveDocumentHandlerFunc turns a function with the right signature into a delete move document handler
type DeleteMoveDocumentHandlerFunc func(DeleteMoveDocumentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMoveDocumentHandlerFunc) Handle(params DeleteMoveDocumentParams) middleware.Responder {
	return fn(params)
}

// DeleteMoveDocumentHandler interface for that can handle valid delete move document params
type DeleteMoveDocumentHandler interface {
	Handle(DeleteMoveDocumentParams) middleware.Responder
}

// NewDeleteMoveDocument creates a new http.Handler for the delete move document operation
func NewDeleteMoveDocument(ctx *middleware.Context, handler DeleteMoveDocumentHandler) *DeleteMoveDocument {
	return &DeleteMoveDocument{Context: ctx, Handler: handler}
}

/*DeleteMoveDocument swagger:route DELETE /move_documents/{moveDocumentId} move_docs deleteMoveDocument

Deletes a move document

Deletes a move document with the given information

*/
type DeleteMoveDocument struct {
	Context *middleware.Context
	Handler DeleteMoveDocumentHandler
}

func (o *DeleteMoveDocument) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteMoveDocumentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
