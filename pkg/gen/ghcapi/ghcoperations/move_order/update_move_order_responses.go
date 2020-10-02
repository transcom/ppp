// Code generated by go-swagger; DO NOT EDIT.

package move_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// UpdateMoveOrderOKCode is the HTTP code returned for type UpdateMoveOrderOK
const UpdateMoveOrderOKCode int = 200

/*UpdateMoveOrderOK updated instance of orders

swagger:response updateMoveOrderOK
*/
type UpdateMoveOrderOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.MoveOrder `json:"body,omitempty"`
}

// NewUpdateMoveOrderOK creates UpdateMoveOrderOK with default headers values
func NewUpdateMoveOrderOK() *UpdateMoveOrderOK {

	return &UpdateMoveOrderOK{}
}

// WithPayload adds the payload to the update move order o k response
func (o *UpdateMoveOrderOK) WithPayload(payload *ghcmessages.MoveOrder) *UpdateMoveOrderOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move order o k response
func (o *UpdateMoveOrderOK) SetPayload(payload *ghcmessages.MoveOrder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveOrderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMoveOrderBadRequestCode is the HTTP code returned for type UpdateMoveOrderBadRequest
const UpdateMoveOrderBadRequestCode int = 400

/*UpdateMoveOrderBadRequest The request payload is invalid

swagger:response updateMoveOrderBadRequest
*/
type UpdateMoveOrderBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveOrderBadRequest creates UpdateMoveOrderBadRequest with default headers values
func NewUpdateMoveOrderBadRequest() *UpdateMoveOrderBadRequest {

	return &UpdateMoveOrderBadRequest{}
}

// WithPayload adds the payload to the update move order bad request response
func (o *UpdateMoveOrderBadRequest) WithPayload(payload interface{}) *UpdateMoveOrderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move order bad request response
func (o *UpdateMoveOrderBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveOrderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveOrderUnauthorizedCode is the HTTP code returned for type UpdateMoveOrderUnauthorized
const UpdateMoveOrderUnauthorizedCode int = 401

/*UpdateMoveOrderUnauthorized The request was unauthenticated

swagger:response updateMoveOrderUnauthorized
*/
type UpdateMoveOrderUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveOrderUnauthorized creates UpdateMoveOrderUnauthorized with default headers values
func NewUpdateMoveOrderUnauthorized() *UpdateMoveOrderUnauthorized {

	return &UpdateMoveOrderUnauthorized{}
}

// WithPayload adds the payload to the update move order unauthorized response
func (o *UpdateMoveOrderUnauthorized) WithPayload(payload interface{}) *UpdateMoveOrderUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move order unauthorized response
func (o *UpdateMoveOrderUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveOrderUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveOrderForbiddenCode is the HTTP code returned for type UpdateMoveOrderForbidden
const UpdateMoveOrderForbiddenCode int = 403

/*UpdateMoveOrderForbidden The request was unauthorized

swagger:response updateMoveOrderForbidden
*/
type UpdateMoveOrderForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveOrderForbidden creates UpdateMoveOrderForbidden with default headers values
func NewUpdateMoveOrderForbidden() *UpdateMoveOrderForbidden {

	return &UpdateMoveOrderForbidden{}
}

// WithPayload adds the payload to the update move order forbidden response
func (o *UpdateMoveOrderForbidden) WithPayload(payload interface{}) *UpdateMoveOrderForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move order forbidden response
func (o *UpdateMoveOrderForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveOrderForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveOrderNotFoundCode is the HTTP code returned for type UpdateMoveOrderNotFound
const UpdateMoveOrderNotFoundCode int = 404

/*UpdateMoveOrderNotFound The requested resource wasn't found

swagger:response updateMoveOrderNotFound
*/
type UpdateMoveOrderNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveOrderNotFound creates UpdateMoveOrderNotFound with default headers values
func NewUpdateMoveOrderNotFound() *UpdateMoveOrderNotFound {

	return &UpdateMoveOrderNotFound{}
}

// WithPayload adds the payload to the update move order not found response
func (o *UpdateMoveOrderNotFound) WithPayload(payload interface{}) *UpdateMoveOrderNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move order not found response
func (o *UpdateMoveOrderNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveOrderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveOrderPreconditionFailedCode is the HTTP code returned for type UpdateMoveOrderPreconditionFailed
const UpdateMoveOrderPreconditionFailedCode int = 412

/*UpdateMoveOrderPreconditionFailed Precondition failed

swagger:response updateMoveOrderPreconditionFailed
*/
type UpdateMoveOrderPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveOrderPreconditionFailed creates UpdateMoveOrderPreconditionFailed with default headers values
func NewUpdateMoveOrderPreconditionFailed() *UpdateMoveOrderPreconditionFailed {

	return &UpdateMoveOrderPreconditionFailed{}
}

// WithPayload adds the payload to the update move order precondition failed response
func (o *UpdateMoveOrderPreconditionFailed) WithPayload(payload interface{}) *UpdateMoveOrderPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move order precondition failed response
func (o *UpdateMoveOrderPreconditionFailed) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveOrderPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveOrderInternalServerErrorCode is the HTTP code returned for type UpdateMoveOrderInternalServerError
const UpdateMoveOrderInternalServerErrorCode int = 500

/*UpdateMoveOrderInternalServerError internal server error

swagger:response updateMoveOrderInternalServerError
*/
type UpdateMoveOrderInternalServerError struct {
}

// NewUpdateMoveOrderInternalServerError creates UpdateMoveOrderInternalServerError with default headers values
func NewUpdateMoveOrderInternalServerError() *UpdateMoveOrderInternalServerError {

	return &UpdateMoveOrderInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateMoveOrderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}