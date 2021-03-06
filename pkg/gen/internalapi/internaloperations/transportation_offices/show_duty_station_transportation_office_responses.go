// Code generated by go-swagger; DO NOT EDIT.

package transportation_offices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/internalmessages"
)

// ShowDutyStationTransportationOfficeOKCode is the HTTP code returned for type ShowDutyStationTransportationOfficeOK
const ShowDutyStationTransportationOfficeOKCode int = 200

/*ShowDutyStationTransportationOfficeOK the instance of the transportation office for a duty station

swagger:response showDutyStationTransportationOfficeOK
*/
type ShowDutyStationTransportationOfficeOK struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.TransportationOffice `json:"body,omitempty"`
}

// NewShowDutyStationTransportationOfficeOK creates ShowDutyStationTransportationOfficeOK with default headers values
func NewShowDutyStationTransportationOfficeOK() *ShowDutyStationTransportationOfficeOK {

	return &ShowDutyStationTransportationOfficeOK{}
}

// WithPayload adds the payload to the show duty station transportation office o k response
func (o *ShowDutyStationTransportationOfficeOK) WithPayload(payload *internalmessages.TransportationOffice) *ShowDutyStationTransportationOfficeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show duty station transportation office o k response
func (o *ShowDutyStationTransportationOfficeOK) SetPayload(payload *internalmessages.TransportationOffice) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowDutyStationTransportationOfficeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowDutyStationTransportationOfficeBadRequestCode is the HTTP code returned for type ShowDutyStationTransportationOfficeBadRequest
const ShowDutyStationTransportationOfficeBadRequestCode int = 400

/*ShowDutyStationTransportationOfficeBadRequest invalid request

swagger:response showDutyStationTransportationOfficeBadRequest
*/
type ShowDutyStationTransportationOfficeBadRequest struct {
}

// NewShowDutyStationTransportationOfficeBadRequest creates ShowDutyStationTransportationOfficeBadRequest with default headers values
func NewShowDutyStationTransportationOfficeBadRequest() *ShowDutyStationTransportationOfficeBadRequest {

	return &ShowDutyStationTransportationOfficeBadRequest{}
}

// WriteResponse to the client
func (o *ShowDutyStationTransportationOfficeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ShowDutyStationTransportationOfficeUnauthorizedCode is the HTTP code returned for type ShowDutyStationTransportationOfficeUnauthorized
const ShowDutyStationTransportationOfficeUnauthorizedCode int = 401

/*ShowDutyStationTransportationOfficeUnauthorized request requires user authentication

swagger:response showDutyStationTransportationOfficeUnauthorized
*/
type ShowDutyStationTransportationOfficeUnauthorized struct {
}

// NewShowDutyStationTransportationOfficeUnauthorized creates ShowDutyStationTransportationOfficeUnauthorized with default headers values
func NewShowDutyStationTransportationOfficeUnauthorized() *ShowDutyStationTransportationOfficeUnauthorized {

	return &ShowDutyStationTransportationOfficeUnauthorized{}
}

// WriteResponse to the client
func (o *ShowDutyStationTransportationOfficeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ShowDutyStationTransportationOfficeForbiddenCode is the HTTP code returned for type ShowDutyStationTransportationOfficeForbidden
const ShowDutyStationTransportationOfficeForbiddenCode int = 403

/*ShowDutyStationTransportationOfficeForbidden user is not authorized

swagger:response showDutyStationTransportationOfficeForbidden
*/
type ShowDutyStationTransportationOfficeForbidden struct {
}

// NewShowDutyStationTransportationOfficeForbidden creates ShowDutyStationTransportationOfficeForbidden with default headers values
func NewShowDutyStationTransportationOfficeForbidden() *ShowDutyStationTransportationOfficeForbidden {

	return &ShowDutyStationTransportationOfficeForbidden{}
}

// WriteResponse to the client
func (o *ShowDutyStationTransportationOfficeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ShowDutyStationTransportationOfficeNotFoundCode is the HTTP code returned for type ShowDutyStationTransportationOfficeNotFound
const ShowDutyStationTransportationOfficeNotFoundCode int = 404

/*ShowDutyStationTransportationOfficeNotFound transportation office not found

swagger:response showDutyStationTransportationOfficeNotFound
*/
type ShowDutyStationTransportationOfficeNotFound struct {
}

// NewShowDutyStationTransportationOfficeNotFound creates ShowDutyStationTransportationOfficeNotFound with default headers values
func NewShowDutyStationTransportationOfficeNotFound() *ShowDutyStationTransportationOfficeNotFound {

	return &ShowDutyStationTransportationOfficeNotFound{}
}

// WriteResponse to the client
func (o *ShowDutyStationTransportationOfficeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ShowDutyStationTransportationOfficeInternalServerErrorCode is the HTTP code returned for type ShowDutyStationTransportationOfficeInternalServerError
const ShowDutyStationTransportationOfficeInternalServerErrorCode int = 500

/*ShowDutyStationTransportationOfficeInternalServerError internal server error

swagger:response showDutyStationTransportationOfficeInternalServerError
*/
type ShowDutyStationTransportationOfficeInternalServerError struct {
}

// NewShowDutyStationTransportationOfficeInternalServerError creates ShowDutyStationTransportationOfficeInternalServerError with default headers values
func NewShowDutyStationTransportationOfficeInternalServerError() *ShowDutyStationTransportationOfficeInternalServerError {

	return &ShowDutyStationTransportationOfficeInternalServerError{}
}

// WriteResponse to the client
func (o *ShowDutyStationTransportationOfficeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
