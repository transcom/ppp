// Code generated by go-swagger; DO NOT EDIT.

package customer_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/adminmessages"
)

// IndexCustomerUsersOKCode is the HTTP code returned for type IndexCustomerUsersOK
const IndexCustomerUsersOKCode int = 200

/*IndexCustomerUsersOK success

swagger:response indexCustomerUsersOK
*/
type IndexCustomerUsersOK struct {
	/*Used for pagination

	 */
	ContentRange string `json:"Content-Range"`

	/*
	  In: Body
	*/
	Payload adminmessages.CustomerUsers `json:"body,omitempty"`
}

// NewIndexCustomerUsersOK creates IndexCustomerUsersOK with default headers values
func NewIndexCustomerUsersOK() *IndexCustomerUsersOK {

	return &IndexCustomerUsersOK{}
}

// WithContentRange adds the contentRange to the index customer users o k response
func (o *IndexCustomerUsersOK) WithContentRange(contentRange string) *IndexCustomerUsersOK {
	o.ContentRange = contentRange
	return o
}

// SetContentRange sets the contentRange to the index customer users o k response
func (o *IndexCustomerUsersOK) SetContentRange(contentRange string) {
	o.ContentRange = contentRange
}

// WithPayload adds the payload to the index customer users o k response
func (o *IndexCustomerUsersOK) WithPayload(payload adminmessages.CustomerUsers) *IndexCustomerUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the index customer users o k response
func (o *IndexCustomerUsersOK) SetPayload(payload adminmessages.CustomerUsers) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IndexCustomerUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Range

	contentRange := o.ContentRange
	if contentRange != "" {
		rw.Header().Set("Content-Range", contentRange)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = adminmessages.CustomerUsers{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IndexCustomerUsersBadRequestCode is the HTTP code returned for type IndexCustomerUsersBadRequest
const IndexCustomerUsersBadRequestCode int = 400

/*IndexCustomerUsersBadRequest invalid request

swagger:response indexCustomerUsersBadRequest
*/
type IndexCustomerUsersBadRequest struct {
}

// NewIndexCustomerUsersBadRequest creates IndexCustomerUsersBadRequest with default headers values
func NewIndexCustomerUsersBadRequest() *IndexCustomerUsersBadRequest {

	return &IndexCustomerUsersBadRequest{}
}

// WriteResponse to the client
func (o *IndexCustomerUsersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// IndexCustomerUsersUnauthorizedCode is the HTTP code returned for type IndexCustomerUsersUnauthorized
const IndexCustomerUsersUnauthorizedCode int = 401

/*IndexCustomerUsersUnauthorized request requires user authentication

swagger:response indexCustomerUsersUnauthorized
*/
type IndexCustomerUsersUnauthorized struct {
}

// NewIndexCustomerUsersUnauthorized creates IndexCustomerUsersUnauthorized with default headers values
func NewIndexCustomerUsersUnauthorized() *IndexCustomerUsersUnauthorized {

	return &IndexCustomerUsersUnauthorized{}
}

// WriteResponse to the client
func (o *IndexCustomerUsersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// IndexCustomerUsersNotFoundCode is the HTTP code returned for type IndexCustomerUsersNotFound
const IndexCustomerUsersNotFoundCode int = 404

/*IndexCustomerUsersNotFound customer users not found

swagger:response indexCustomerUsersNotFound
*/
type IndexCustomerUsersNotFound struct {
}

// NewIndexCustomerUsersNotFound creates IndexCustomerUsersNotFound with default headers values
func NewIndexCustomerUsersNotFound() *IndexCustomerUsersNotFound {

	return &IndexCustomerUsersNotFound{}
}

// WriteResponse to the client
func (o *IndexCustomerUsersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// IndexCustomerUsersInternalServerErrorCode is the HTTP code returned for type IndexCustomerUsersInternalServerError
const IndexCustomerUsersInternalServerErrorCode int = 500

/*IndexCustomerUsersInternalServerError server error

swagger:response indexCustomerUsersInternalServerError
*/
type IndexCustomerUsersInternalServerError struct {
}

// NewIndexCustomerUsersInternalServerError creates IndexCustomerUsersInternalServerError with default headers values
func NewIndexCustomerUsersInternalServerError() *IndexCustomerUsersInternalServerError {

	return &IndexCustomerUsersInternalServerError{}
}

// WriteResponse to the client
func (o *IndexCustomerUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
