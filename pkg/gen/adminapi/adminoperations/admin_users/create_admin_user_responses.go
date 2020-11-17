// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/adminmessages"
)

// CreateAdminUserCreatedCode is the HTTP code returned for type CreateAdminUserCreated
const CreateAdminUserCreatedCode int = 201

/*CreateAdminUserCreated Successfully created Admin User

swagger:response createAdminUserCreated
*/
type CreateAdminUserCreated struct {

	/*
	  In: Body
	*/
	Payload *adminmessages.AdminUser `json:"body,omitempty"`
}

// NewCreateAdminUserCreated creates CreateAdminUserCreated with default headers values
func NewCreateAdminUserCreated() *CreateAdminUserCreated {

	return &CreateAdminUserCreated{}
}

// WithPayload adds the payload to the create admin user created response
func (o *CreateAdminUserCreated) WithPayload(payload *adminmessages.AdminUser) *CreateAdminUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create admin user created response
func (o *CreateAdminUserCreated) SetPayload(payload *adminmessages.AdminUser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAdminUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAdminUserBadRequestCode is the HTTP code returned for type CreateAdminUserBadRequest
const CreateAdminUserBadRequestCode int = 400

/*CreateAdminUserBadRequest Invalid Request

swagger:response createAdminUserBadRequest
*/
type CreateAdminUserBadRequest struct {
}

// NewCreateAdminUserBadRequest creates CreateAdminUserBadRequest with default headers values
func NewCreateAdminUserBadRequest() *CreateAdminUserBadRequest {

	return &CreateAdminUserBadRequest{}
}

// WriteResponse to the client
func (o *CreateAdminUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateAdminUserUnauthorizedCode is the HTTP code returned for type CreateAdminUserUnauthorized
const CreateAdminUserUnauthorizedCode int = 401

/*CreateAdminUserUnauthorized Must be authenticated to use this end point

swagger:response createAdminUserUnauthorized
*/
type CreateAdminUserUnauthorized struct {
}

// NewCreateAdminUserUnauthorized creates CreateAdminUserUnauthorized with default headers values
func NewCreateAdminUserUnauthorized() *CreateAdminUserUnauthorized {

	return &CreateAdminUserUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAdminUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// CreateAdminUserForbiddenCode is the HTTP code returned for type CreateAdminUserForbidden
const CreateAdminUserForbiddenCode int = 403

/*CreateAdminUserForbidden Not authorized to create an admin user

swagger:response createAdminUserForbidden
*/
type CreateAdminUserForbidden struct {
}

// NewCreateAdminUserForbidden creates CreateAdminUserForbidden with default headers values
func NewCreateAdminUserForbidden() *CreateAdminUserForbidden {

	return &CreateAdminUserForbidden{}
}

// WriteResponse to the client
func (o *CreateAdminUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// CreateAdminUserInternalServerErrorCode is the HTTP code returned for type CreateAdminUserInternalServerError
const CreateAdminUserInternalServerErrorCode int = 500

/*CreateAdminUserInternalServerError Server error

swagger:response createAdminUserInternalServerError
*/
type CreateAdminUserInternalServerError struct {
}

// NewCreateAdminUserInternalServerError creates CreateAdminUserInternalServerError with default headers values
func NewCreateAdminUserInternalServerError() *CreateAdminUserInternalServerError {

	return &CreateAdminUserInternalServerError{}
}

// WriteResponse to the client
func (o *CreateAdminUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
