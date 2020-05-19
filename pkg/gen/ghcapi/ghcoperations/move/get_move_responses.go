// Code generated by go-swagger; DO NOT EDIT.

package move

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// GetMoveOKCode is the HTTP code returned for type GetMoveOK
const GetMoveOKCode int = 200

/*GetMoveOK Successfully retrieved the individual move

swagger:response getMoveOK
*/
type GetMoveOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Move `json:"body,omitempty"`
}

// NewGetMoveOK creates GetMoveOK with default headers values
func NewGetMoveOK() *GetMoveOK {

	return &GetMoveOK{}
}

// WithPayload adds the payload to the get move o k response
func (o *GetMoveOK) WithPayload(payload *ghcmessages.Move) *GetMoveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move o k response
func (o *GetMoveOK) SetPayload(payload *ghcmessages.Move) {
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

/*GetMoveBadRequest The locator provided is invalid

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

/*GetMoveUnauthorized The request was denied

swagger:response getMoveUnauthorized
*/
type GetMoveUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetMoveUnauthorized creates GetMoveUnauthorized with default headers values
func NewGetMoveUnauthorized() *GetMoveUnauthorized {

	return &GetMoveUnauthorized{}
}

// WithPayload adds the payload to the get move unauthorized response
func (o *GetMoveUnauthorized) WithPayload(payload interface{}) *GetMoveUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move unauthorized response
func (o *GetMoveUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMoveForbiddenCode is the HTTP code returned for type GetMoveForbidden
const GetMoveForbiddenCode int = 403

/*GetMoveForbidden The request was denied

swagger:response getMoveForbidden
*/
type GetMoveForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetMoveForbidden creates GetMoveForbidden with default headers values
func NewGetMoveForbidden() *GetMoveForbidden {

	return &GetMoveForbidden{}
}

// WithPayload adds the payload to the get move forbidden response
func (o *GetMoveForbidden) WithPayload(payload interface{}) *GetMoveForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move forbidden response
func (o *GetMoveForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMoveNotFoundCode is the HTTP code returned for type GetMoveNotFound
const GetMoveNotFoundCode int = 404

/*GetMoveNotFound The requested resource wasn't found

swagger:response getMoveNotFound
*/
type GetMoveNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetMoveNotFound creates GetMoveNotFound with default headers values
func NewGetMoveNotFound() *GetMoveNotFound {

	return &GetMoveNotFound{}
}

// WithPayload adds the payload to the get move not found response
func (o *GetMoveNotFound) WithPayload(payload interface{}) *GetMoveNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move not found response
func (o *GetMoveNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMoveInternalServerErrorCode is the HTTP code returned for type GetMoveInternalServerError
const GetMoveInternalServerErrorCode int = 500

/*GetMoveInternalServerError A server error occurred

swagger:response getMoveInternalServerError
*/
type GetMoveInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetMoveInternalServerError creates GetMoveInternalServerError with default headers values
func NewGetMoveInternalServerError() *GetMoveInternalServerError {

	return &GetMoveInternalServerError{}
}

// WithPayload adds the payload to the get move internal server error response
func (o *GetMoveInternalServerError) WithPayload(payload interface{}) *GetMoveInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get move internal server error response
func (o *GetMoveInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMoveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}