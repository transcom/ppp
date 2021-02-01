// Code generated by go-swagger; DO NOT EDIT.

package move

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/adminmessages"
)

// GetMoveOKCode is the HTTP code returned for type GetMoveOK
const GetMoveOKCode int = 200

/*GetMoveOK Success

swagger:response getMoveOK
*/
type GetMoveOK struct {

	/*
	  In: Body
	*/
	Payload *adminmessages.Move `json:"body,omitempty"`
}

// NewGetMoveOK creates GetMoveOK with default headers values
func NewGetMoveOK() *GetMoveOK {

	return &GetMoveOK{}
}

// WithPayload adds the payload to the get move o k response
func (o *GetMoveOK) WithPayload(payload *adminmessages.Move) *GetMoveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move o k response
func (o *GetMoveOK) SetPayload(payload *adminmessages.Move) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMoveBadRequestCode is the HTTP code returned for type GetMoveBadRequest
const GetMoveBadRequestCode int = 400

/*GetMoveBadRequest Invalid request

swagger:response getMoveBadRequest
*/
type GetMoveBadRequest struct {
}

// NewGetMoveBadRequest creates GetMoveBadRequest with default headers values
func NewGetMoveBadRequest() *GetMoveBadRequest {

	return &GetMoveBadRequest{}
}

// WriteResponse to the client
func (o *GetMoveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetMoveUnauthorizedCode is the HTTP code returned for type GetMoveUnauthorized
const GetMoveUnauthorizedCode int = 401

/*GetMoveUnauthorized Must be authenticated to use this endpoint

swagger:response getMoveUnauthorized
*/
type GetMoveUnauthorized struct {
}

// NewGetMoveUnauthorized creates GetMoveUnauthorized with default headers values
func NewGetMoveUnauthorized() *GetMoveUnauthorized {

	return &GetMoveUnauthorized{}
}

// WriteResponse to the client
func (o *GetMoveUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetMoveNotFoundCode is the HTTP code returned for type GetMoveNotFound
const GetMoveNotFoundCode int = 404

/*GetMoveNotFound Move not found

swagger:response getMoveNotFound
*/
type GetMoveNotFound struct {
}

// NewGetMoveNotFound creates GetMoveNotFound with default headers values
func NewGetMoveNotFound() *GetMoveNotFound {

	return &GetMoveNotFound{}
}

// WriteResponse to the client
func (o *GetMoveNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetMoveInternalServerErrorCode is the HTTP code returned for type GetMoveInternalServerError
const GetMoveInternalServerErrorCode int = 500

/*GetMoveInternalServerError Server error

swagger:response getMoveInternalServerError
*/
type GetMoveInternalServerError struct {
}

// NewGetMoveInternalServerError creates GetMoveInternalServerError with default headers values
func NewGetMoveInternalServerError() *GetMoveInternalServerError {

	return &GetMoveInternalServerError{}
}

// WriteResponse to the client
func (o *GetMoveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}