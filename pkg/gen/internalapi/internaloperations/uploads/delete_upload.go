// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteUploadHandlerFunc turns a function with the right signature into a delete upload handler
type DeleteUploadHandlerFunc func(DeleteUploadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUploadHandlerFunc) Handle(params DeleteUploadParams) middleware.Responder {
	return fn(params)
}

// DeleteUploadHandler interface for that can handle valid delete upload params
type DeleteUploadHandler interface {
	Handle(DeleteUploadParams) middleware.Responder
}

// NewDeleteUpload creates a new http.Handler for the delete upload operation
func NewDeleteUpload(ctx *middleware.Context, handler DeleteUploadHandler) *DeleteUpload {
	return &DeleteUpload{Context: ctx, Handler: handler}
}

/*DeleteUpload swagger:route DELETE /uploads/{uploadId} uploads deleteUpload

Deletes an upload

Uploads represent a single digital file, such as a JPEG or PDF.

*/
type DeleteUpload struct {
	Context *middleware.Context
	Handler DeleteUploadHandler
}

func (o *DeleteUpload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteUploadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
