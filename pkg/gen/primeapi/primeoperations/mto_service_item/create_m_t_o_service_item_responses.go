// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// CreateMTOServiceItemOKCode is the HTTP code returned for type CreateMTOServiceItemOK
const CreateMTOServiceItemOKCode int = 200

/*CreateMTOServiceItemOK created instance of a mto service item

swagger:response createMTOServiceItemOK
*/
type CreateMTOServiceItemOK struct {

	/*
	  In: Body
	*/
	Payload primemessages.MTOServiceItem `json:"body,omitempty"`
}

// NewCreateMTOServiceItemOK creates CreateMTOServiceItemOK with default headers values
func NewCreateMTOServiceItemOK() *CreateMTOServiceItemOK {

	return &CreateMTOServiceItemOK{}
}

// WithPayload adds the payload to the create m t o service item o k response
func (o *CreateMTOServiceItemOK) WithPayload(payload primemessages.MTOServiceItem) *CreateMTOServiceItemOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o service item o k response
func (o *CreateMTOServiceItemOK) SetPayload(payload primemessages.MTOServiceItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOServiceItemOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMTOServiceItemBadRequestCode is the HTTP code returned for type CreateMTOServiceItemBadRequest
const CreateMTOServiceItemBadRequestCode int = 400

/*CreateMTOServiceItemBadRequest invalid request

swagger:response createMTOServiceItemBadRequest
*/
type CreateMTOServiceItemBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewCreateMTOServiceItemBadRequest creates CreateMTOServiceItemBadRequest with default headers values
func NewCreateMTOServiceItemBadRequest() *CreateMTOServiceItemBadRequest {

	return &CreateMTOServiceItemBadRequest{}
}

// WithPayload adds the payload to the create m t o service item bad request response
func (o *CreateMTOServiceItemBadRequest) WithPayload(payload interface{}) *CreateMTOServiceItemBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o service item bad request response
func (o *CreateMTOServiceItemBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOServiceItemBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CreateMTOServiceItemNotFoundCode is the HTTP code returned for type CreateMTOServiceItemNotFound
const CreateMTOServiceItemNotFoundCode int = 404

/*CreateMTOServiceItemNotFound The requested resource wasn't found

swagger:response createMTOServiceItemNotFound
*/
type CreateMTOServiceItemNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewCreateMTOServiceItemNotFound creates CreateMTOServiceItemNotFound with default headers values
func NewCreateMTOServiceItemNotFound() *CreateMTOServiceItemNotFound {

	return &CreateMTOServiceItemNotFound{}
}

// WithPayload adds the payload to the create m t o service item not found response
func (o *CreateMTOServiceItemNotFound) WithPayload(payload interface{}) *CreateMTOServiceItemNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o service item not found response
func (o *CreateMTOServiceItemNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOServiceItemNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CreateMTOServiceItemInternalServerErrorCode is the HTTP code returned for type CreateMTOServiceItemInternalServerError
const CreateMTOServiceItemInternalServerErrorCode int = 500

/*CreateMTOServiceItemInternalServerError internal server error

swagger:response createMTOServiceItemInternalServerError
*/
type CreateMTOServiceItemInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewCreateMTOServiceItemInternalServerError creates CreateMTOServiceItemInternalServerError with default headers values
func NewCreateMTOServiceItemInternalServerError() *CreateMTOServiceItemInternalServerError {

	return &CreateMTOServiceItemInternalServerError{}
}

// WithPayload adds the payload to the create m t o service item internal server error response
func (o *CreateMTOServiceItemInternalServerError) WithPayload(payload interface{}) *CreateMTOServiceItemInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o service item internal server error response
func (o *CreateMTOServiceItemInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOServiceItemInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
