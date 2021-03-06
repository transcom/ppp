// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// UpdateOrderOKCode is the HTTP code returned for type UpdateOrderOK
const UpdateOrderOKCode int = 200

/*UpdateOrderOK updated instance of orders

swagger:response updateOrderOK
*/
type UpdateOrderOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Order `json:"body,omitempty"`
}

// NewUpdateOrderOK creates UpdateOrderOK with default headers values
func NewUpdateOrderOK() *UpdateOrderOK {

	return &UpdateOrderOK{}
}

// WithPayload adds the payload to the update order o k response
func (o *UpdateOrderOK) WithPayload(payload *ghcmessages.Order) *UpdateOrderOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update order o k response
func (o *UpdateOrderOK) SetPayload(payload *ghcmessages.Order) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOrderBadRequestCode is the HTTP code returned for type UpdateOrderBadRequest
const UpdateOrderBadRequestCode int = 400

/*UpdateOrderBadRequest The request payload is invalid

swagger:response updateOrderBadRequest
*/
type UpdateOrderBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateOrderBadRequest creates UpdateOrderBadRequest with default headers values
func NewUpdateOrderBadRequest() *UpdateOrderBadRequest {

	return &UpdateOrderBadRequest{}
}

// WithPayload adds the payload to the update order bad request response
func (o *UpdateOrderBadRequest) WithPayload(payload interface{}) *UpdateOrderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update order bad request response
func (o *UpdateOrderBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateOrderUnauthorizedCode is the HTTP code returned for type UpdateOrderUnauthorized
const UpdateOrderUnauthorizedCode int = 401

/*UpdateOrderUnauthorized The request was unauthenticated

swagger:response updateOrderUnauthorized
*/
type UpdateOrderUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateOrderUnauthorized creates UpdateOrderUnauthorized with default headers values
func NewUpdateOrderUnauthorized() *UpdateOrderUnauthorized {

	return &UpdateOrderUnauthorized{}
}

// WithPayload adds the payload to the update order unauthorized response
func (o *UpdateOrderUnauthorized) WithPayload(payload interface{}) *UpdateOrderUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update order unauthorized response
func (o *UpdateOrderUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrderUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateOrderForbiddenCode is the HTTP code returned for type UpdateOrderForbidden
const UpdateOrderForbiddenCode int = 403

/*UpdateOrderForbidden The request was unauthorized

swagger:response updateOrderForbidden
*/
type UpdateOrderForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateOrderForbidden creates UpdateOrderForbidden with default headers values
func NewUpdateOrderForbidden() *UpdateOrderForbidden {

	return &UpdateOrderForbidden{}
}

// WithPayload adds the payload to the update order forbidden response
func (o *UpdateOrderForbidden) WithPayload(payload interface{}) *UpdateOrderForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update order forbidden response
func (o *UpdateOrderForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrderForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateOrderNotFoundCode is the HTTP code returned for type UpdateOrderNotFound
const UpdateOrderNotFoundCode int = 404

/*UpdateOrderNotFound The requested resource wasn't found

swagger:response updateOrderNotFound
*/
type UpdateOrderNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateOrderNotFound creates UpdateOrderNotFound with default headers values
func NewUpdateOrderNotFound() *UpdateOrderNotFound {

	return &UpdateOrderNotFound{}
}

// WithPayload adds the payload to the update order not found response
func (o *UpdateOrderNotFound) WithPayload(payload interface{}) *UpdateOrderNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update order not found response
func (o *UpdateOrderNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateOrderPreconditionFailedCode is the HTTP code returned for type UpdateOrderPreconditionFailed
const UpdateOrderPreconditionFailedCode int = 412

/*UpdateOrderPreconditionFailed Precondition failed

swagger:response updateOrderPreconditionFailed
*/
type UpdateOrderPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateOrderPreconditionFailed creates UpdateOrderPreconditionFailed with default headers values
func NewUpdateOrderPreconditionFailed() *UpdateOrderPreconditionFailed {

	return &UpdateOrderPreconditionFailed{}
}

// WithPayload adds the payload to the update order precondition failed response
func (o *UpdateOrderPreconditionFailed) WithPayload(payload interface{}) *UpdateOrderPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update order precondition failed response
func (o *UpdateOrderPreconditionFailed) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrderPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateOrderInternalServerErrorCode is the HTTP code returned for type UpdateOrderInternalServerError
const UpdateOrderInternalServerErrorCode int = 500

/*UpdateOrderInternalServerError internal server error

swagger:response updateOrderInternalServerError
*/
type UpdateOrderInternalServerError struct {
}

// NewUpdateOrderInternalServerError creates UpdateOrderInternalServerError with default headers values
func NewUpdateOrderInternalServerError() *UpdateOrderInternalServerError {

	return &UpdateOrderInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateOrderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
