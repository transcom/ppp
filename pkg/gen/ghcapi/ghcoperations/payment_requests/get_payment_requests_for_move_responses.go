// Code generated by go-swagger; DO NOT EDIT.

package payment_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// GetPaymentRequestsForMoveOKCode is the HTTP code returned for type GetPaymentRequestsForMoveOK
const GetPaymentRequestsForMoveOKCode int = 200

/*GetPaymentRequestsForMoveOK Successfully retrieved all line items for a move task order

swagger:response getPaymentRequestsForMoveOK
*/
type GetPaymentRequestsForMoveOK struct {

	/*
	  In: Body
	*/
	Payload ghcmessages.PaymentRequests `json:"body,omitempty"`
}

// NewGetPaymentRequestsForMoveOK creates GetPaymentRequestsForMoveOK with default headers values
func NewGetPaymentRequestsForMoveOK() *GetPaymentRequestsForMoveOK {

	return &GetPaymentRequestsForMoveOK{}
}

// WithPayload adds the payload to the get payment requests for move o k response
func (o *GetPaymentRequestsForMoveOK) WithPayload(payload ghcmessages.PaymentRequests) *GetPaymentRequestsForMoveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment requests for move o k response
func (o *GetPaymentRequestsForMoveOK) SetPayload(payload ghcmessages.PaymentRequests) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentRequestsForMoveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = ghcmessages.PaymentRequests{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetPaymentRequestsForMoveForbiddenCode is the HTTP code returned for type GetPaymentRequestsForMoveForbidden
const GetPaymentRequestsForMoveForbiddenCode int = 403

/*GetPaymentRequestsForMoveForbidden The request was denied

swagger:response getPaymentRequestsForMoveForbidden
*/
type GetPaymentRequestsForMoveForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetPaymentRequestsForMoveForbidden creates GetPaymentRequestsForMoveForbidden with default headers values
func NewGetPaymentRequestsForMoveForbidden() *GetPaymentRequestsForMoveForbidden {

	return &GetPaymentRequestsForMoveForbidden{}
}

// WithPayload adds the payload to the get payment requests for move forbidden response
func (o *GetPaymentRequestsForMoveForbidden) WithPayload(payload interface{}) *GetPaymentRequestsForMoveForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment requests for move forbidden response
func (o *GetPaymentRequestsForMoveForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentRequestsForMoveForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetPaymentRequestsForMoveNotFoundCode is the HTTP code returned for type GetPaymentRequestsForMoveNotFound
const GetPaymentRequestsForMoveNotFoundCode int = 404

/*GetPaymentRequestsForMoveNotFound The requested resource wasn't found

swagger:response getPaymentRequestsForMoveNotFound
*/
type GetPaymentRequestsForMoveNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetPaymentRequestsForMoveNotFound creates GetPaymentRequestsForMoveNotFound with default headers values
func NewGetPaymentRequestsForMoveNotFound() *GetPaymentRequestsForMoveNotFound {

	return &GetPaymentRequestsForMoveNotFound{}
}

// WithPayload adds the payload to the get payment requests for move not found response
func (o *GetPaymentRequestsForMoveNotFound) WithPayload(payload interface{}) *GetPaymentRequestsForMoveNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment requests for move not found response
func (o *GetPaymentRequestsForMoveNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentRequestsForMoveNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetPaymentRequestsForMoveUnprocessableEntityCode is the HTTP code returned for type GetPaymentRequestsForMoveUnprocessableEntity
const GetPaymentRequestsForMoveUnprocessableEntityCode int = 422

/*GetPaymentRequestsForMoveUnprocessableEntity Validation error

swagger:response getPaymentRequestsForMoveUnprocessableEntity
*/
type GetPaymentRequestsForMoveUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.ValidationError `json:"body,omitempty"`
}

// NewGetPaymentRequestsForMoveUnprocessableEntity creates GetPaymentRequestsForMoveUnprocessableEntity with default headers values
func NewGetPaymentRequestsForMoveUnprocessableEntity() *GetPaymentRequestsForMoveUnprocessableEntity {

	return &GetPaymentRequestsForMoveUnprocessableEntity{}
}

// WithPayload adds the payload to the get payment requests for move unprocessable entity response
func (o *GetPaymentRequestsForMoveUnprocessableEntity) WithPayload(payload *ghcmessages.ValidationError) *GetPaymentRequestsForMoveUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment requests for move unprocessable entity response
func (o *GetPaymentRequestsForMoveUnprocessableEntity) SetPayload(payload *ghcmessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentRequestsForMoveUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPaymentRequestsForMoveInternalServerErrorCode is the HTTP code returned for type GetPaymentRequestsForMoveInternalServerError
const GetPaymentRequestsForMoveInternalServerErrorCode int = 500

/*GetPaymentRequestsForMoveInternalServerError A server error occurred

swagger:response getPaymentRequestsForMoveInternalServerError
*/
type GetPaymentRequestsForMoveInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetPaymentRequestsForMoveInternalServerError creates GetPaymentRequestsForMoveInternalServerError with default headers values
func NewGetPaymentRequestsForMoveInternalServerError() *GetPaymentRequestsForMoveInternalServerError {

	return &GetPaymentRequestsForMoveInternalServerError{}
}

// WithPayload adds the payload to the get payment requests for move internal server error response
func (o *GetPaymentRequestsForMoveInternalServerError) WithPayload(payload interface{}) *GetPaymentRequestsForMoveInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment requests for move internal server error response
func (o *GetPaymentRequestsForMoveInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentRequestsForMoveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
