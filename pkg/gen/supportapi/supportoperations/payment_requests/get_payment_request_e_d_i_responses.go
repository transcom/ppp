// Code generated by go-swagger; DO NOT EDIT.

package payment_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	supportmessages "github.com/transcom/mymove/pkg/gen/supportmessages"
)

// GetPaymentRequestEDIOKCode is the HTTP code returned for type GetPaymentRequestEDIOK
const GetPaymentRequestEDIOKCode int = 200

/*GetPaymentRequestEDIOK Successfully retrieved payment requests associated with a given move task order

swagger:response getPaymentRequestEDIOK
*/
type GetPaymentRequestEDIOK struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.PaymentRequestEDI `json:"body,omitempty"`
}

// NewGetPaymentRequestEDIOK creates GetPaymentRequestEDIOK with default headers values
func NewGetPaymentRequestEDIOK() *GetPaymentRequestEDIOK {

	return &GetPaymentRequestEDIOK{}
}

// WithPayload adds the payload to the get payment request e d i o k response
func (o *GetPaymentRequestEDIOK) WithPayload(payload *supportmessages.PaymentRequestEDI) *GetPaymentRequestEDIOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment request e d i o k response
func (o *GetPaymentRequestEDIOK) SetPayload(payload *supportmessages.PaymentRequestEDI) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentRequestEDIOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPaymentRequestEDIBadRequestCode is the HTTP code returned for type GetPaymentRequestEDIBadRequest
const GetPaymentRequestEDIBadRequestCode int = 400

/*GetPaymentRequestEDIBadRequest The request payload is invalid.

swagger:response getPaymentRequestEDIBadRequest
*/
type GetPaymentRequestEDIBadRequest struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewGetPaymentRequestEDIBadRequest creates GetPaymentRequestEDIBadRequest with default headers values
func NewGetPaymentRequestEDIBadRequest() *GetPaymentRequestEDIBadRequest {

	return &GetPaymentRequestEDIBadRequest{}
}

// WithPayload adds the payload to the get payment request e d i bad request response
func (o *GetPaymentRequestEDIBadRequest) WithPayload(payload *supportmessages.ClientError) *GetPaymentRequestEDIBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment request e d i bad request response
func (o *GetPaymentRequestEDIBadRequest) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentRequestEDIBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPaymentRequestEDIUnauthorizedCode is the HTTP code returned for type GetPaymentRequestEDIUnauthorized
const GetPaymentRequestEDIUnauthorizedCode int = 401

/*GetPaymentRequestEDIUnauthorized The request was denied.

swagger:response getPaymentRequestEDIUnauthorized
*/
type GetPaymentRequestEDIUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewGetPaymentRequestEDIUnauthorized creates GetPaymentRequestEDIUnauthorized with default headers values
func NewGetPaymentRequestEDIUnauthorized() *GetPaymentRequestEDIUnauthorized {

	return &GetPaymentRequestEDIUnauthorized{}
}

// WithPayload adds the payload to the get payment request e d i unauthorized response
func (o *GetPaymentRequestEDIUnauthorized) WithPayload(payload *supportmessages.ClientError) *GetPaymentRequestEDIUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment request e d i unauthorized response
func (o *GetPaymentRequestEDIUnauthorized) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentRequestEDIUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPaymentRequestEDIForbiddenCode is the HTTP code returned for type GetPaymentRequestEDIForbidden
const GetPaymentRequestEDIForbiddenCode int = 403

/*GetPaymentRequestEDIForbidden The request was denied.

swagger:response getPaymentRequestEDIForbidden
*/
type GetPaymentRequestEDIForbidden struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewGetPaymentRequestEDIForbidden creates GetPaymentRequestEDIForbidden with default headers values
func NewGetPaymentRequestEDIForbidden() *GetPaymentRequestEDIForbidden {

	return &GetPaymentRequestEDIForbidden{}
}

// WithPayload adds the payload to the get payment request e d i forbidden response
func (o *GetPaymentRequestEDIForbidden) WithPayload(payload *supportmessages.ClientError) *GetPaymentRequestEDIForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment request e d i forbidden response
func (o *GetPaymentRequestEDIForbidden) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentRequestEDIForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPaymentRequestEDINotFoundCode is the HTTP code returned for type GetPaymentRequestEDINotFound
const GetPaymentRequestEDINotFoundCode int = 404

/*GetPaymentRequestEDINotFound The requested resource wasn't found.

swagger:response getPaymentRequestEDINotFound
*/
type GetPaymentRequestEDINotFound struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.ClientError `json:"body,omitempty"`
}

// NewGetPaymentRequestEDINotFound creates GetPaymentRequestEDINotFound with default headers values
func NewGetPaymentRequestEDINotFound() *GetPaymentRequestEDINotFound {

	return &GetPaymentRequestEDINotFound{}
}

// WithPayload adds the payload to the get payment request e d i not found response
func (o *GetPaymentRequestEDINotFound) WithPayload(payload *supportmessages.ClientError) *GetPaymentRequestEDINotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment request e d i not found response
func (o *GetPaymentRequestEDINotFound) SetPayload(payload *supportmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentRequestEDINotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPaymentRequestEDIInternalServerErrorCode is the HTTP code returned for type GetPaymentRequestEDIInternalServerError
const GetPaymentRequestEDIInternalServerErrorCode int = 500

/*GetPaymentRequestEDIInternalServerError A server error occurred.

swagger:response getPaymentRequestEDIInternalServerError
*/
type GetPaymentRequestEDIInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *supportmessages.Error `json:"body,omitempty"`
}

// NewGetPaymentRequestEDIInternalServerError creates GetPaymentRequestEDIInternalServerError with default headers values
func NewGetPaymentRequestEDIInternalServerError() *GetPaymentRequestEDIInternalServerError {

	return &GetPaymentRequestEDIInternalServerError{}
}

// WithPayload adds the payload to the get payment request e d i internal server error response
func (o *GetPaymentRequestEDIInternalServerError) WithPayload(payload *supportmessages.Error) *GetPaymentRequestEDIInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payment request e d i internal server error response
func (o *GetPaymentRequestEDIInternalServerError) SetPayload(payload *supportmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentRequestEDIInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
