// Code generated by go-swagger; DO NOT EDIT.

package service_members

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
)

// ShowServiceMemberOKCode is the HTTP code returned for type ShowServiceMemberOK
const ShowServiceMemberOKCode int = 200

/*ShowServiceMemberOK the instance of the service member

swagger:response showServiceMemberOK
*/
type ShowServiceMemberOK struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.ServiceMemberPayload `json:"body,omitempty"`
}

// NewShowServiceMemberOK creates ShowServiceMemberOK with default headers values
func NewShowServiceMemberOK() *ShowServiceMemberOK {

	return &ShowServiceMemberOK{}
}

// WithPayload adds the payload to the show service member o k response
func (o *ShowServiceMemberOK) WithPayload(payload *internalmessages.ServiceMemberPayload) *ShowServiceMemberOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show service member o k response
func (o *ShowServiceMemberOK) SetPayload(payload *internalmessages.ServiceMemberPayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowServiceMemberOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowServiceMemberBadRequestCode is the HTTP code returned for type ShowServiceMemberBadRequest
const ShowServiceMemberBadRequestCode int = 400

/*ShowServiceMemberBadRequest invalid request

swagger:response showServiceMemberBadRequest
*/
type ShowServiceMemberBadRequest struct {
}

// NewShowServiceMemberBadRequest creates ShowServiceMemberBadRequest with default headers values
func NewShowServiceMemberBadRequest() *ShowServiceMemberBadRequest {

	return &ShowServiceMemberBadRequest{}
}

// WriteResponse to the client
func (o *ShowServiceMemberBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ShowServiceMemberUnauthorizedCode is the HTTP code returned for type ShowServiceMemberUnauthorized
const ShowServiceMemberUnauthorizedCode int = 401

/*ShowServiceMemberUnauthorized request requires user authentication

swagger:response showServiceMemberUnauthorized
*/
type ShowServiceMemberUnauthorized struct {
}

// NewShowServiceMemberUnauthorized creates ShowServiceMemberUnauthorized with default headers values
func NewShowServiceMemberUnauthorized() *ShowServiceMemberUnauthorized {

	return &ShowServiceMemberUnauthorized{}
}

// WriteResponse to the client
func (o *ShowServiceMemberUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ShowServiceMemberForbiddenCode is the HTTP code returned for type ShowServiceMemberForbidden
const ShowServiceMemberForbiddenCode int = 403

/*ShowServiceMemberForbidden user is not authorized

swagger:response showServiceMemberForbidden
*/
type ShowServiceMemberForbidden struct {
}

// NewShowServiceMemberForbidden creates ShowServiceMemberForbidden with default headers values
func NewShowServiceMemberForbidden() *ShowServiceMemberForbidden {

	return &ShowServiceMemberForbidden{}
}

// WriteResponse to the client
func (o *ShowServiceMemberForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ShowServiceMemberNotFoundCode is the HTTP code returned for type ShowServiceMemberNotFound
const ShowServiceMemberNotFoundCode int = 404

/*ShowServiceMemberNotFound service member not found

swagger:response showServiceMemberNotFound
*/
type ShowServiceMemberNotFound struct {
}

// NewShowServiceMemberNotFound creates ShowServiceMemberNotFound with default headers values
func NewShowServiceMemberNotFound() *ShowServiceMemberNotFound {

	return &ShowServiceMemberNotFound{}
}

// WriteResponse to the client
func (o *ShowServiceMemberNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ShowServiceMemberInternalServerErrorCode is the HTTP code returned for type ShowServiceMemberInternalServerError
const ShowServiceMemberInternalServerErrorCode int = 500

/*ShowServiceMemberInternalServerError internal server error

swagger:response showServiceMemberInternalServerError
*/
type ShowServiceMemberInternalServerError struct {
}

// NewShowServiceMemberInternalServerError creates ShowServiceMemberInternalServerError with default headers values
func NewShowServiceMemberInternalServerError() *ShowServiceMemberInternalServerError {

	return &ShowServiceMemberInternalServerError{}
}

// WriteResponse to the client
func (o *ShowServiceMemberInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
