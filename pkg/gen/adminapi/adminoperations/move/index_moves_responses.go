// Code generated by go-swagger; DO NOT EDIT.

package move

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/adminmessages"
)

// IndexMovesOKCode is the HTTP code returned for type IndexMovesOK
const IndexMovesOKCode int = 200

/*IndexMovesOK success

swagger:response indexMovesOK
*/
type IndexMovesOK struct {
	/*Used for pagination

	 */
	ContentRange string `json:"Content-Range"`

	/*
	  In: Body
	*/
	Payload adminmessages.Moves `json:"body,omitempty"`
}

// NewIndexMovesOK creates IndexMovesOK with default headers values
func NewIndexMovesOK() *IndexMovesOK {

	return &IndexMovesOK{}
}

// WithContentRange adds the contentRange to the index moves o k response
func (o *IndexMovesOK) WithContentRange(contentRange string) *IndexMovesOK {
	o.ContentRange = contentRange
	return o
}

// SetContentRange sets the contentRange to the index moves o k response
func (o *IndexMovesOK) SetContentRange(contentRange string) {
	o.ContentRange = contentRange
}

// WithPayload adds the payload to the index moves o k response
func (o *IndexMovesOK) WithPayload(payload adminmessages.Moves) *IndexMovesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the index moves o k response
func (o *IndexMovesOK) SetPayload(payload adminmessages.Moves) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IndexMovesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Range

	contentRange := o.ContentRange
	if contentRange != "" {
		rw.Header().Set("Content-Range", contentRange)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = adminmessages.Moves{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IndexMovesBadRequestCode is the HTTP code returned for type IndexMovesBadRequest
const IndexMovesBadRequestCode int = 400

/*IndexMovesBadRequest invalid request

swagger:response indexMovesBadRequest
*/
type IndexMovesBadRequest struct {
}

// NewIndexMovesBadRequest creates IndexMovesBadRequest with default headers values
func NewIndexMovesBadRequest() *IndexMovesBadRequest {

	return &IndexMovesBadRequest{}
}

// WriteResponse to the client
func (o *IndexMovesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// IndexMovesUnauthorizedCode is the HTTP code returned for type IndexMovesUnauthorized
const IndexMovesUnauthorizedCode int = 401

/*IndexMovesUnauthorized request requires user authentication

swagger:response indexMovesUnauthorized
*/
type IndexMovesUnauthorized struct {
}

// NewIndexMovesUnauthorized creates IndexMovesUnauthorized with default headers values
func NewIndexMovesUnauthorized() *IndexMovesUnauthorized {

	return &IndexMovesUnauthorized{}
}

// WriteResponse to the client
func (o *IndexMovesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// IndexMovesNotFoundCode is the HTTP code returned for type IndexMovesNotFound
const IndexMovesNotFoundCode int = 404

/*IndexMovesNotFound not found

swagger:response indexMovesNotFound
*/
type IndexMovesNotFound struct {
}

// NewIndexMovesNotFound creates IndexMovesNotFound with default headers values
func NewIndexMovesNotFound() *IndexMovesNotFound {

	return &IndexMovesNotFound{}
}

// WriteResponse to the client
func (o *IndexMovesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// IndexMovesInternalServerErrorCode is the HTTP code returned for type IndexMovesInternalServerError
const IndexMovesInternalServerErrorCode int = 500

/*IndexMovesInternalServerError server error

swagger:response indexMovesInternalServerError
*/
type IndexMovesInternalServerError struct {
}

// NewIndexMovesInternalServerError creates IndexMovesInternalServerError with default headers values
func NewIndexMovesInternalServerError() *IndexMovesInternalServerError {

	return &IndexMovesInternalServerError{}
}

// WriteResponse to the client
func (o *IndexMovesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
