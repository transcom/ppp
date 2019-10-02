// Code generated by go-swagger; DO NOT EDIT.

package backup_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"
)

// CreateServiceMemberBackupContactCreatedCode is the HTTP code returned for type CreateServiceMemberBackupContactCreated
const CreateServiceMemberBackupContactCreatedCode int = 201

/*CreateServiceMemberBackupContactCreated created instance of service member backup contact

swagger:response createServiceMemberBackupContactCreated
*/
type CreateServiceMemberBackupContactCreated struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.ServiceMemberBackupContactPayload `json:"body,omitempty"`
}

// NewCreateServiceMemberBackupContactCreated creates CreateServiceMemberBackupContactCreated with default headers values
func NewCreateServiceMemberBackupContactCreated() *CreateServiceMemberBackupContactCreated {

	return &CreateServiceMemberBackupContactCreated{}
}

// WithPayload adds the payload to the create service member backup contact created response
func (o *CreateServiceMemberBackupContactCreated) WithPayload(payload *internalmessages.ServiceMemberBackupContactPayload) *CreateServiceMemberBackupContactCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create service member backup contact created response
func (o *CreateServiceMemberBackupContactCreated) SetPayload(payload *internalmessages.ServiceMemberBackupContactPayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServiceMemberBackupContactCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServiceMemberBackupContactBadRequestCode is the HTTP code returned for type CreateServiceMemberBackupContactBadRequest
const CreateServiceMemberBackupContactBadRequestCode int = 400

/*CreateServiceMemberBackupContactBadRequest invalid request

swagger:response createServiceMemberBackupContactBadRequest
*/
type CreateServiceMemberBackupContactBadRequest struct {
}

// NewCreateServiceMemberBackupContactBadRequest creates CreateServiceMemberBackupContactBadRequest with default headers values
func NewCreateServiceMemberBackupContactBadRequest() *CreateServiceMemberBackupContactBadRequest {

	return &CreateServiceMemberBackupContactBadRequest{}
}

// WriteResponse to the client
func (o *CreateServiceMemberBackupContactBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateServiceMemberBackupContactUnauthorizedCode is the HTTP code returned for type CreateServiceMemberBackupContactUnauthorized
const CreateServiceMemberBackupContactUnauthorizedCode int = 401

/*CreateServiceMemberBackupContactUnauthorized request requires user authentication

swagger:response createServiceMemberBackupContactUnauthorized
*/
type CreateServiceMemberBackupContactUnauthorized struct {
}

// NewCreateServiceMemberBackupContactUnauthorized creates CreateServiceMemberBackupContactUnauthorized with default headers values
func NewCreateServiceMemberBackupContactUnauthorized() *CreateServiceMemberBackupContactUnauthorized {

	return &CreateServiceMemberBackupContactUnauthorized{}
}

// WriteResponse to the client
func (o *CreateServiceMemberBackupContactUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// CreateServiceMemberBackupContactForbiddenCode is the HTTP code returned for type CreateServiceMemberBackupContactForbidden
const CreateServiceMemberBackupContactForbiddenCode int = 403

/*CreateServiceMemberBackupContactForbidden user is not authorized to create this backup contact

swagger:response createServiceMemberBackupContactForbidden
*/
type CreateServiceMemberBackupContactForbidden struct {
}

// NewCreateServiceMemberBackupContactForbidden creates CreateServiceMemberBackupContactForbidden with default headers values
func NewCreateServiceMemberBackupContactForbidden() *CreateServiceMemberBackupContactForbidden {

	return &CreateServiceMemberBackupContactForbidden{}
}

// WriteResponse to the client
func (o *CreateServiceMemberBackupContactForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// CreateServiceMemberBackupContactNotFoundCode is the HTTP code returned for type CreateServiceMemberBackupContactNotFound
const CreateServiceMemberBackupContactNotFoundCode int = 404

/*CreateServiceMemberBackupContactNotFound contact not found

swagger:response createServiceMemberBackupContactNotFound
*/
type CreateServiceMemberBackupContactNotFound struct {
}

// NewCreateServiceMemberBackupContactNotFound creates CreateServiceMemberBackupContactNotFound with default headers values
func NewCreateServiceMemberBackupContactNotFound() *CreateServiceMemberBackupContactNotFound {

	return &CreateServiceMemberBackupContactNotFound{}
}

// WriteResponse to the client
func (o *CreateServiceMemberBackupContactNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// CreateServiceMemberBackupContactInternalServerErrorCode is the HTTP code returned for type CreateServiceMemberBackupContactInternalServerError
const CreateServiceMemberBackupContactInternalServerErrorCode int = 500

/*CreateServiceMemberBackupContactInternalServerError internal server error

swagger:response createServiceMemberBackupContactInternalServerError
*/
type CreateServiceMemberBackupContactInternalServerError struct {
}

// NewCreateServiceMemberBackupContactInternalServerError creates CreateServiceMemberBackupContactInternalServerError with default headers values
func NewCreateServiceMemberBackupContactInternalServerError() *CreateServiceMemberBackupContactInternalServerError {

	return &CreateServiceMemberBackupContactInternalServerError{}
}

// WriteResponse to the client
func (o *CreateServiceMemberBackupContactInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
