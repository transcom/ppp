// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
)

// UpdatePersonallyProcuredMoveOKCode is the HTTP code returned for type UpdatePersonallyProcuredMoveOK
const UpdatePersonallyProcuredMoveOKCode int = 200

/*UpdatePersonallyProcuredMoveOK updated instance of personally_procured_move

swagger:response updatePersonallyProcuredMoveOK
*/
type UpdatePersonallyProcuredMoveOK struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.PersonallyProcuredMovePayload `json:"body,omitempty"`
}

// NewUpdatePersonallyProcuredMoveOK creates UpdatePersonallyProcuredMoveOK with default headers values
func NewUpdatePersonallyProcuredMoveOK() *UpdatePersonallyProcuredMoveOK {

	return &UpdatePersonallyProcuredMoveOK{}
}

// WithPayload adds the payload to the update personally procured move o k response
func (o *UpdatePersonallyProcuredMoveOK) WithPayload(payload *internalmessages.PersonallyProcuredMovePayload) *UpdatePersonallyProcuredMoveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update personally procured move o k response
func (o *UpdatePersonallyProcuredMoveOK) SetPayload(payload *internalmessages.PersonallyProcuredMovePayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePersonallyProcuredMoveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePersonallyProcuredMoveBadRequestCode is the HTTP code returned for type UpdatePersonallyProcuredMoveBadRequest
const UpdatePersonallyProcuredMoveBadRequestCode int = 400

/*UpdatePersonallyProcuredMoveBadRequest invalid request

swagger:response updatePersonallyProcuredMoveBadRequest
*/
type UpdatePersonallyProcuredMoveBadRequest struct {
}

// NewUpdatePersonallyProcuredMoveBadRequest creates UpdatePersonallyProcuredMoveBadRequest with default headers values
func NewUpdatePersonallyProcuredMoveBadRequest() *UpdatePersonallyProcuredMoveBadRequest {

	return &UpdatePersonallyProcuredMoveBadRequest{}
}

// WriteResponse to the client
func (o *UpdatePersonallyProcuredMoveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdatePersonallyProcuredMoveUnauthorizedCode is the HTTP code returned for type UpdatePersonallyProcuredMoveUnauthorized
const UpdatePersonallyProcuredMoveUnauthorizedCode int = 401

/*UpdatePersonallyProcuredMoveUnauthorized request requires user authentication

swagger:response updatePersonallyProcuredMoveUnauthorized
*/
type UpdatePersonallyProcuredMoveUnauthorized struct {
}

// NewUpdatePersonallyProcuredMoveUnauthorized creates UpdatePersonallyProcuredMoveUnauthorized with default headers values
func NewUpdatePersonallyProcuredMoveUnauthorized() *UpdatePersonallyProcuredMoveUnauthorized {

	return &UpdatePersonallyProcuredMoveUnauthorized{}
}

// WriteResponse to the client
func (o *UpdatePersonallyProcuredMoveUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// UpdatePersonallyProcuredMoveForbiddenCode is the HTTP code returned for type UpdatePersonallyProcuredMoveForbidden
const UpdatePersonallyProcuredMoveForbiddenCode int = 403

/*UpdatePersonallyProcuredMoveForbidden user is not authorized

swagger:response updatePersonallyProcuredMoveForbidden
*/
type UpdatePersonallyProcuredMoveForbidden struct {
}

// NewUpdatePersonallyProcuredMoveForbidden creates UpdatePersonallyProcuredMoveForbidden with default headers values
func NewUpdatePersonallyProcuredMoveForbidden() *UpdatePersonallyProcuredMoveForbidden {

	return &UpdatePersonallyProcuredMoveForbidden{}
}

// WriteResponse to the client
func (o *UpdatePersonallyProcuredMoveForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// UpdatePersonallyProcuredMoveInternalServerErrorCode is the HTTP code returned for type UpdatePersonallyProcuredMoveInternalServerError
const UpdatePersonallyProcuredMoveInternalServerErrorCode int = 500

/*UpdatePersonallyProcuredMoveInternalServerError internal server error

swagger:response updatePersonallyProcuredMoveInternalServerError
*/
type UpdatePersonallyProcuredMoveInternalServerError struct {
}

// NewUpdatePersonallyProcuredMoveInternalServerError creates UpdatePersonallyProcuredMoveInternalServerError with default headers values
func NewUpdatePersonallyProcuredMoveInternalServerError() *UpdatePersonallyProcuredMoveInternalServerError {

	return &UpdatePersonallyProcuredMoveInternalServerError{}
}

// WriteResponse to the client
func (o *UpdatePersonallyProcuredMoveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
