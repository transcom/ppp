// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	adminmessages "github.com/transcom/mymove/pkg/gen/adminmessages"
)

// RevokeUserSessionOKCode is the HTTP code returned for type RevokeUserSessionOK
const RevokeUserSessionOKCode int = 200

/*RevokeUserSessionOK Successfully updated User

swagger:response revokeUserSessionOK
*/
type RevokeUserSessionOK struct {

	/*
	  In: Body
	*/
	Payload *adminmessages.User `json:"body,omitempty"`
}

// NewRevokeUserSessionOK creates RevokeUserSessionOK with default headers values
func NewRevokeUserSessionOK() *RevokeUserSessionOK {

	return &RevokeUserSessionOK{}
}

// WithPayload adds the payload to the revoke user session o k response
func (o *RevokeUserSessionOK) WithPayload(payload *adminmessages.User) *RevokeUserSessionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the revoke user session o k response
func (o *RevokeUserSessionOK) SetPayload(payload *adminmessages.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RevokeUserSessionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RevokeUserSessionBadRequestCode is the HTTP code returned for type RevokeUserSessionBadRequest
const RevokeUserSessionBadRequestCode int = 400

/*RevokeUserSessionBadRequest Invalid Request

swagger:response revokeUserSessionBadRequest
*/
type RevokeUserSessionBadRequest struct {
}

// NewRevokeUserSessionBadRequest creates RevokeUserSessionBadRequest with default headers values
func NewRevokeUserSessionBadRequest() *RevokeUserSessionBadRequest {

	return &RevokeUserSessionBadRequest{}
}

// WriteResponse to the client
func (o *RevokeUserSessionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// RevokeUserSessionUnauthorizedCode is the HTTP code returned for type RevokeUserSessionUnauthorized
const RevokeUserSessionUnauthorizedCode int = 401

/*RevokeUserSessionUnauthorized Must be authenticated to use this end point

swagger:response revokeUserSessionUnauthorized
*/
type RevokeUserSessionUnauthorized struct {
}

// NewRevokeUserSessionUnauthorized creates RevokeUserSessionUnauthorized with default headers values
func NewRevokeUserSessionUnauthorized() *RevokeUserSessionUnauthorized {

	return &RevokeUserSessionUnauthorized{}
}

// WriteResponse to the client
func (o *RevokeUserSessionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// RevokeUserSessionForbiddenCode is the HTTP code returned for type RevokeUserSessionForbidden
const RevokeUserSessionForbiddenCode int = 403

/*RevokeUserSessionForbidden Not authorized to update this user

swagger:response revokeUserSessionForbidden
*/
type RevokeUserSessionForbidden struct {
}

// NewRevokeUserSessionForbidden creates RevokeUserSessionForbidden with default headers values
func NewRevokeUserSessionForbidden() *RevokeUserSessionForbidden {

	return &RevokeUserSessionForbidden{}
}

// WriteResponse to the client
func (o *RevokeUserSessionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// RevokeUserSessionInternalServerErrorCode is the HTTP code returned for type RevokeUserSessionInternalServerError
const RevokeUserSessionInternalServerErrorCode int = 500

/*RevokeUserSessionInternalServerError Server error

swagger:response revokeUserSessionInternalServerError
*/
type RevokeUserSessionInternalServerError struct {
}

// NewRevokeUserSessionInternalServerError creates RevokeUserSessionInternalServerError with default headers values
func NewRevokeUserSessionInternalServerError() *RevokeUserSessionInternalServerError {

	return &RevokeUserSessionInternalServerError{}
}

// WriteResponse to the client
func (o *RevokeUserSessionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}