// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// UpdateMoveTaskOrderOKCode is the HTTP code returned for type UpdateMoveTaskOrderOK
const UpdateMoveTaskOrderOKCode int = 200

/*UpdateMoveTaskOrderOK Successfully retrieved move task order

swagger:response updateMoveTaskOrderOK
*/
type UpdateMoveTaskOrderOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.MoveTaskOrder `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderOK creates UpdateMoveTaskOrderOK with default headers values
func NewUpdateMoveTaskOrderOK() *UpdateMoveTaskOrderOK {

	return &UpdateMoveTaskOrderOK{}
}

// WithPayload adds the payload to the update move task order o k response
func (o *UpdateMoveTaskOrderOK) WithPayload(payload *ghcmessages.MoveTaskOrder) *UpdateMoveTaskOrderOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order o k response
func (o *UpdateMoveTaskOrderOK) SetPayload(payload *ghcmessages.MoveTaskOrder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMoveTaskOrderBadRequestCode is the HTTP code returned for type UpdateMoveTaskOrderBadRequest
const UpdateMoveTaskOrderBadRequestCode int = 400

/*UpdateMoveTaskOrderBadRequest The request payload is invalid

swagger:response updateMoveTaskOrderBadRequest
*/
type UpdateMoveTaskOrderBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderBadRequest creates UpdateMoveTaskOrderBadRequest with default headers values
func NewUpdateMoveTaskOrderBadRequest() *UpdateMoveTaskOrderBadRequest {

	return &UpdateMoveTaskOrderBadRequest{}
}

// WithPayload adds the payload to the update move task order bad request response
func (o *UpdateMoveTaskOrderBadRequest) WithPayload(payload interface{}) *UpdateMoveTaskOrderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order bad request response
func (o *UpdateMoveTaskOrderBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderUnauthorizedCode is the HTTP code returned for type UpdateMoveTaskOrderUnauthorized
const UpdateMoveTaskOrderUnauthorizedCode int = 401

/*UpdateMoveTaskOrderUnauthorized The request was denied

swagger:response updateMoveTaskOrderUnauthorized
*/
type UpdateMoveTaskOrderUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderUnauthorized creates UpdateMoveTaskOrderUnauthorized with default headers values
func NewUpdateMoveTaskOrderUnauthorized() *UpdateMoveTaskOrderUnauthorized {

	return &UpdateMoveTaskOrderUnauthorized{}
}

// WithPayload adds the payload to the update move task order unauthorized response
func (o *UpdateMoveTaskOrderUnauthorized) WithPayload(payload interface{}) *UpdateMoveTaskOrderUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order unauthorized response
func (o *UpdateMoveTaskOrderUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderForbiddenCode is the HTTP code returned for type UpdateMoveTaskOrderForbidden
const UpdateMoveTaskOrderForbiddenCode int = 403

/*UpdateMoveTaskOrderForbidden The request was denied

swagger:response updateMoveTaskOrderForbidden
*/
type UpdateMoveTaskOrderForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderForbidden creates UpdateMoveTaskOrderForbidden with default headers values
func NewUpdateMoveTaskOrderForbidden() *UpdateMoveTaskOrderForbidden {

	return &UpdateMoveTaskOrderForbidden{}
}

// WithPayload adds the payload to the update move task order forbidden response
func (o *UpdateMoveTaskOrderForbidden) WithPayload(payload interface{}) *UpdateMoveTaskOrderForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order forbidden response
func (o *UpdateMoveTaskOrderForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderNotFoundCode is the HTTP code returned for type UpdateMoveTaskOrderNotFound
const UpdateMoveTaskOrderNotFoundCode int = 404

/*UpdateMoveTaskOrderNotFound The requested resource wasn't found

swagger:response updateMoveTaskOrderNotFound
*/
type UpdateMoveTaskOrderNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderNotFound creates UpdateMoveTaskOrderNotFound with default headers values
func NewUpdateMoveTaskOrderNotFound() *UpdateMoveTaskOrderNotFound {

	return &UpdateMoveTaskOrderNotFound{}
}

// WithPayload adds the payload to the update move task order not found response
func (o *UpdateMoveTaskOrderNotFound) WithPayload(payload interface{}) *UpdateMoveTaskOrderNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order not found response
func (o *UpdateMoveTaskOrderNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderPreconditionFailedCode is the HTTP code returned for type UpdateMoveTaskOrderPreconditionFailed
const UpdateMoveTaskOrderPreconditionFailedCode int = 412

/*UpdateMoveTaskOrderPreconditionFailed Precondition failed

swagger:response updateMoveTaskOrderPreconditionFailed
*/
type UpdateMoveTaskOrderPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderPreconditionFailed creates UpdateMoveTaskOrderPreconditionFailed with default headers values
func NewUpdateMoveTaskOrderPreconditionFailed() *UpdateMoveTaskOrderPreconditionFailed {

	return &UpdateMoveTaskOrderPreconditionFailed{}
}

// WithPayload adds the payload to the update move task order precondition failed response
func (o *UpdateMoveTaskOrderPreconditionFailed) WithPayload(payload interface{}) *UpdateMoveTaskOrderPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order precondition failed response
func (o *UpdateMoveTaskOrderPreconditionFailed) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderInternalServerErrorCode is the HTTP code returned for type UpdateMoveTaskOrderInternalServerError
const UpdateMoveTaskOrderInternalServerErrorCode int = 500

/*UpdateMoveTaskOrderInternalServerError A server error occurred

swagger:response updateMoveTaskOrderInternalServerError
*/
type UpdateMoveTaskOrderInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderInternalServerError creates UpdateMoveTaskOrderInternalServerError with default headers values
func NewUpdateMoveTaskOrderInternalServerError() *UpdateMoveTaskOrderInternalServerError {

	return &UpdateMoveTaskOrderInternalServerError{}
}

// WithPayload adds the payload to the update move task order internal server error response
func (o *UpdateMoveTaskOrderInternalServerError) WithPayload(payload interface{}) *UpdateMoveTaskOrderInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order internal server error response
func (o *UpdateMoveTaskOrderInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
