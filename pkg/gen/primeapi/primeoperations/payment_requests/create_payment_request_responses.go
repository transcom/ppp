// Code generated by go-swagger; DO NOT EDIT.

package payment_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// CreatePaymentRequestCreatedCode is the HTTP code returned for type CreatePaymentRequestCreated
const CreatePaymentRequestCreatedCode int = 201

/*CreatePaymentRequestCreated Successfully created a paymentRequest object.

swagger:response createPaymentRequestCreated
*/
type CreatePaymentRequestCreated struct {

	/*
	  In: Body
	*/
	Payload *primemessages.PaymentRequest `json:"body,omitempty"`
}

// NewCreatePaymentRequestCreated creates CreatePaymentRequestCreated with default headers values
func NewCreatePaymentRequestCreated() *CreatePaymentRequestCreated {

	return &CreatePaymentRequestCreated{}
}

// WithPayload adds the payload to the create payment request created response
func (o *CreatePaymentRequestCreated) WithPayload(payload *primemessages.PaymentRequest) *CreatePaymentRequestCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create payment request created response
func (o *CreatePaymentRequestCreated) SetPayload(payload *primemessages.PaymentRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePaymentRequestCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePaymentRequestBadRequestCode is the HTTP code returned for type CreatePaymentRequestBadRequest
const CreatePaymentRequestBadRequestCode int = 400

/*CreatePaymentRequestBadRequest Request payload is invalid.

swagger:response createPaymentRequestBadRequest
*/
type CreatePaymentRequestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewCreatePaymentRequestBadRequest creates CreatePaymentRequestBadRequest with default headers values
func NewCreatePaymentRequestBadRequest() *CreatePaymentRequestBadRequest {

	return &CreatePaymentRequestBadRequest{}
}

// WithPayload adds the payload to the create payment request bad request response
func (o *CreatePaymentRequestBadRequest) WithPayload(payload *primemessages.ClientError) *CreatePaymentRequestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create payment request bad request response
func (o *CreatePaymentRequestBadRequest) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePaymentRequestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePaymentRequestUnauthorizedCode is the HTTP code returned for type CreatePaymentRequestUnauthorized
const CreatePaymentRequestUnauthorizedCode int = 401

/*CreatePaymentRequestUnauthorized The request was denied.

swagger:response createPaymentRequestUnauthorized
*/
type CreatePaymentRequestUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewCreatePaymentRequestUnauthorized creates CreatePaymentRequestUnauthorized with default headers values
func NewCreatePaymentRequestUnauthorized() *CreatePaymentRequestUnauthorized {

	return &CreatePaymentRequestUnauthorized{}
}

// WithPayload adds the payload to the create payment request unauthorized response
func (o *CreatePaymentRequestUnauthorized) WithPayload(payload *primemessages.ClientError) *CreatePaymentRequestUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create payment request unauthorized response
func (o *CreatePaymentRequestUnauthorized) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePaymentRequestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePaymentRequestForbiddenCode is the HTTP code returned for type CreatePaymentRequestForbidden
const CreatePaymentRequestForbiddenCode int = 403

/*CreatePaymentRequestForbidden The request was denied.

swagger:response createPaymentRequestForbidden
*/
type CreatePaymentRequestForbidden struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewCreatePaymentRequestForbidden creates CreatePaymentRequestForbidden with default headers values
func NewCreatePaymentRequestForbidden() *CreatePaymentRequestForbidden {

	return &CreatePaymentRequestForbidden{}
}

// WithPayload adds the payload to the create payment request forbidden response
func (o *CreatePaymentRequestForbidden) WithPayload(payload *primemessages.ClientError) *CreatePaymentRequestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create payment request forbidden response
func (o *CreatePaymentRequestForbidden) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePaymentRequestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePaymentRequestNotFoundCode is the HTTP code returned for type CreatePaymentRequestNotFound
const CreatePaymentRequestNotFoundCode int = 404

/*CreatePaymentRequestNotFound The requested resource wasn't found.

swagger:response createPaymentRequestNotFound
*/
type CreatePaymentRequestNotFound struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewCreatePaymentRequestNotFound creates CreatePaymentRequestNotFound with default headers values
func NewCreatePaymentRequestNotFound() *CreatePaymentRequestNotFound {

	return &CreatePaymentRequestNotFound{}
}

// WithPayload adds the payload to the create payment request not found response
func (o *CreatePaymentRequestNotFound) WithPayload(payload *primemessages.ClientError) *CreatePaymentRequestNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create payment request not found response
func (o *CreatePaymentRequestNotFound) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePaymentRequestNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePaymentRequestUnprocessableEntityCode is the HTTP code returned for type CreatePaymentRequestUnprocessableEntity
const CreatePaymentRequestUnprocessableEntityCode int = 422

/*CreatePaymentRequestUnprocessableEntity The payload was unprocessable.

swagger:response createPaymentRequestUnprocessableEntity
*/
type CreatePaymentRequestUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ValidationError `json:"body,omitempty"`
}

// NewCreatePaymentRequestUnprocessableEntity creates CreatePaymentRequestUnprocessableEntity with default headers values
func NewCreatePaymentRequestUnprocessableEntity() *CreatePaymentRequestUnprocessableEntity {

	return &CreatePaymentRequestUnprocessableEntity{}
}

// WithPayload adds the payload to the create payment request unprocessable entity response
func (o *CreatePaymentRequestUnprocessableEntity) WithPayload(payload *primemessages.ValidationError) *CreatePaymentRequestUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create payment request unprocessable entity response
func (o *CreatePaymentRequestUnprocessableEntity) SetPayload(payload *primemessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePaymentRequestUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreatePaymentRequestInternalServerErrorCode is the HTTP code returned for type CreatePaymentRequestInternalServerError
const CreatePaymentRequestInternalServerErrorCode int = 500

/*CreatePaymentRequestInternalServerError A server error occurred.

swagger:response createPaymentRequestInternalServerError
*/
type CreatePaymentRequestInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *primemessages.Error `json:"body,omitempty"`
}

// NewCreatePaymentRequestInternalServerError creates CreatePaymentRequestInternalServerError with default headers values
func NewCreatePaymentRequestInternalServerError() *CreatePaymentRequestInternalServerError {

	return &CreatePaymentRequestInternalServerError{}
}

// WithPayload adds the payload to the create payment request internal server error response
func (o *CreatePaymentRequestInternalServerError) WithPayload(payload *primemessages.Error) *CreatePaymentRequestInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create payment request internal server error response
func (o *CreatePaymentRequestInternalServerError) SetPayload(payload *primemessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePaymentRequestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
