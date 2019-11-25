// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// UpdateMoveTaskOrderDestinationAddressOKCode is the HTTP code returned for type UpdateMoveTaskOrderDestinationAddressOK
const UpdateMoveTaskOrderDestinationAddressOKCode int = 200

/*UpdateMoveTaskOrderDestinationAddressOK Successfully updated move task order destination address

swagger:response updateMoveTaskOrderDestinationAddressOK
*/
type UpdateMoveTaskOrderDestinationAddressOK struct {

	/*
	  In: Body
	*/
	Payload *primemessages.MoveTaskOrder `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderDestinationAddressOK creates UpdateMoveTaskOrderDestinationAddressOK with default headers values
func NewUpdateMoveTaskOrderDestinationAddressOK() *UpdateMoveTaskOrderDestinationAddressOK {

	return &UpdateMoveTaskOrderDestinationAddressOK{}
}

// WithPayload adds the payload to the update move task order destination address o k response
func (o *UpdateMoveTaskOrderDestinationAddressOK) WithPayload(payload *primemessages.MoveTaskOrder) *UpdateMoveTaskOrderDestinationAddressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order destination address o k response
func (o *UpdateMoveTaskOrderDestinationAddressOK) SetPayload(payload *primemessages.MoveTaskOrder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderDestinationAddressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMoveTaskOrderDestinationAddressUnauthorizedCode is the HTTP code returned for type UpdateMoveTaskOrderDestinationAddressUnauthorized
const UpdateMoveTaskOrderDestinationAddressUnauthorizedCode int = 401

/*UpdateMoveTaskOrderDestinationAddressUnauthorized The request was denied

swagger:response updateMoveTaskOrderDestinationAddressUnauthorized
*/
type UpdateMoveTaskOrderDestinationAddressUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderDestinationAddressUnauthorized creates UpdateMoveTaskOrderDestinationAddressUnauthorized with default headers values
func NewUpdateMoveTaskOrderDestinationAddressUnauthorized() *UpdateMoveTaskOrderDestinationAddressUnauthorized {

	return &UpdateMoveTaskOrderDestinationAddressUnauthorized{}
}

// WithPayload adds the payload to the update move task order destination address unauthorized response
func (o *UpdateMoveTaskOrderDestinationAddressUnauthorized) WithPayload(payload interface{}) *UpdateMoveTaskOrderDestinationAddressUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order destination address unauthorized response
func (o *UpdateMoveTaskOrderDestinationAddressUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderDestinationAddressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderDestinationAddressForbiddenCode is the HTTP code returned for type UpdateMoveTaskOrderDestinationAddressForbidden
const UpdateMoveTaskOrderDestinationAddressForbiddenCode int = 403

/*UpdateMoveTaskOrderDestinationAddressForbidden The request was denied

swagger:response updateMoveTaskOrderDestinationAddressForbidden
*/
type UpdateMoveTaskOrderDestinationAddressForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderDestinationAddressForbidden creates UpdateMoveTaskOrderDestinationAddressForbidden with default headers values
func NewUpdateMoveTaskOrderDestinationAddressForbidden() *UpdateMoveTaskOrderDestinationAddressForbidden {

	return &UpdateMoveTaskOrderDestinationAddressForbidden{}
}

// WithPayload adds the payload to the update move task order destination address forbidden response
func (o *UpdateMoveTaskOrderDestinationAddressForbidden) WithPayload(payload interface{}) *UpdateMoveTaskOrderDestinationAddressForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order destination address forbidden response
func (o *UpdateMoveTaskOrderDestinationAddressForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderDestinationAddressForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderDestinationAddressNotFoundCode is the HTTP code returned for type UpdateMoveTaskOrderDestinationAddressNotFound
const UpdateMoveTaskOrderDestinationAddressNotFoundCode int = 404

/*UpdateMoveTaskOrderDestinationAddressNotFound The requested resource wasn't found

swagger:response updateMoveTaskOrderDestinationAddressNotFound
*/
type UpdateMoveTaskOrderDestinationAddressNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderDestinationAddressNotFound creates UpdateMoveTaskOrderDestinationAddressNotFound with default headers values
func NewUpdateMoveTaskOrderDestinationAddressNotFound() *UpdateMoveTaskOrderDestinationAddressNotFound {

	return &UpdateMoveTaskOrderDestinationAddressNotFound{}
}

// WithPayload adds the payload to the update move task order destination address not found response
func (o *UpdateMoveTaskOrderDestinationAddressNotFound) WithPayload(payload interface{}) *UpdateMoveTaskOrderDestinationAddressNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order destination address not found response
func (o *UpdateMoveTaskOrderDestinationAddressNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderDestinationAddressNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateMoveTaskOrderDestinationAddressUnprocessableEntityCode is the HTTP code returned for type UpdateMoveTaskOrderDestinationAddressUnprocessableEntity
const UpdateMoveTaskOrderDestinationAddressUnprocessableEntityCode int = 422

/*UpdateMoveTaskOrderDestinationAddressUnprocessableEntity The request payload is invalid

swagger:response updateMoveTaskOrderDestinationAddressUnprocessableEntity
*/
type UpdateMoveTaskOrderDestinationAddressUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ValidationError `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderDestinationAddressUnprocessableEntity creates UpdateMoveTaskOrderDestinationAddressUnprocessableEntity with default headers values
func NewUpdateMoveTaskOrderDestinationAddressUnprocessableEntity() *UpdateMoveTaskOrderDestinationAddressUnprocessableEntity {

	return &UpdateMoveTaskOrderDestinationAddressUnprocessableEntity{}
}

// WithPayload adds the payload to the update move task order destination address unprocessable entity response
func (o *UpdateMoveTaskOrderDestinationAddressUnprocessableEntity) WithPayload(payload *primemessages.ValidationError) *UpdateMoveTaskOrderDestinationAddressUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order destination address unprocessable entity response
func (o *UpdateMoveTaskOrderDestinationAddressUnprocessableEntity) SetPayload(payload *primemessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderDestinationAddressUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateMoveTaskOrderDestinationAddressInternalServerErrorCode is the HTTP code returned for type UpdateMoveTaskOrderDestinationAddressInternalServerError
const UpdateMoveTaskOrderDestinationAddressInternalServerErrorCode int = 500

/*UpdateMoveTaskOrderDestinationAddressInternalServerError A server error occurred

swagger:response updateMoveTaskOrderDestinationAddressInternalServerError
*/
type UpdateMoveTaskOrderDestinationAddressInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewUpdateMoveTaskOrderDestinationAddressInternalServerError creates UpdateMoveTaskOrderDestinationAddressInternalServerError with default headers values
func NewUpdateMoveTaskOrderDestinationAddressInternalServerError() *UpdateMoveTaskOrderDestinationAddressInternalServerError {

	return &UpdateMoveTaskOrderDestinationAddressInternalServerError{}
}

// WithPayload adds the payload to the update move task order destination address internal server error response
func (o *UpdateMoveTaskOrderDestinationAddressInternalServerError) WithPayload(payload interface{}) *UpdateMoveTaskOrderDestinationAddressInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update move task order destination address internal server error response
func (o *UpdateMoveTaskOrderDestinationAddressInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateMoveTaskOrderDestinationAddressInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
