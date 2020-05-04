// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	supportmessages "github.com/transcom/mymove/pkg/gen/supportmessages"
)

// UpdateMTOServiceItemStatusReader is a Reader for the UpdateMTOServiceItemStatus structure.
type UpdateMTOServiceItemStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMTOServiceItemStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMTOServiceItemStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateMTOServiceItemStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateMTOServiceItemStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMTOServiceItemStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMTOServiceItemStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateMTOServiceItemStatusConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewUpdateMTOServiceItemStatusPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateMTOServiceItemStatusUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMTOServiceItemStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateMTOServiceItemStatusOK creates a UpdateMTOServiceItemStatusOK with default headers values
func NewUpdateMTOServiceItemStatusOK() *UpdateMTOServiceItemStatusOK {
	return &UpdateMTOServiceItemStatusOK{}
}

/*UpdateMTOServiceItemStatusOK handles this case with default header values.

Successfully updated service item status for a move task order.
*/
type UpdateMTOServiceItemStatusOK struct {
	Payload *supportmessages.UpdateMTOServiceItemStatus
}

func (o *UpdateMTOServiceItemStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /service-items/{mtoServiceItemID}/status][%d] updateMTOServiceItemStatusOK  %+v", 200, o.Payload)
}

func (o *UpdateMTOServiceItemStatusOK) GetPayload() *supportmessages.UpdateMTOServiceItemStatus {
	return o.Payload
}

func (o *UpdateMTOServiceItemStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.UpdateMTOServiceItemStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOServiceItemStatusBadRequest creates a UpdateMTOServiceItemStatusBadRequest with default headers values
func NewUpdateMTOServiceItemStatusBadRequest() *UpdateMTOServiceItemStatusBadRequest {
	return &UpdateMTOServiceItemStatusBadRequest{}
}

/*UpdateMTOServiceItemStatusBadRequest handles this case with default header values.

The parameters were invalid.
*/
type UpdateMTOServiceItemStatusBadRequest struct {
	Payload *supportmessages.Error
}

func (o *UpdateMTOServiceItemStatusBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /service-items/{mtoServiceItemID}/status][%d] updateMTOServiceItemStatusBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateMTOServiceItemStatusBadRequest) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *UpdateMTOServiceItemStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOServiceItemStatusUnauthorized creates a UpdateMTOServiceItemStatusUnauthorized with default headers values
func NewUpdateMTOServiceItemStatusUnauthorized() *UpdateMTOServiceItemStatusUnauthorized {
	return &UpdateMTOServiceItemStatusUnauthorized{}
}

/*UpdateMTOServiceItemStatusUnauthorized handles this case with default header values.

The request was unauthorized.
*/
type UpdateMTOServiceItemStatusUnauthorized struct {
	Payload interface{}
}

func (o *UpdateMTOServiceItemStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /service-items/{mtoServiceItemID}/status][%d] updateMTOServiceItemStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateMTOServiceItemStatusUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOServiceItemStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOServiceItemStatusForbidden creates a UpdateMTOServiceItemStatusForbidden with default headers values
func NewUpdateMTOServiceItemStatusForbidden() *UpdateMTOServiceItemStatusForbidden {
	return &UpdateMTOServiceItemStatusForbidden{}
}

/*UpdateMTOServiceItemStatusForbidden handles this case with default header values.

The client doesn't have permissions to perform the request.
*/
type UpdateMTOServiceItemStatusForbidden struct {
	Payload interface{}
}

func (o *UpdateMTOServiceItemStatusForbidden) Error() string {
	return fmt.Sprintf("[PATCH /service-items/{mtoServiceItemID}/status][%d] updateMTOServiceItemStatusForbidden  %+v", 403, o.Payload)
}

func (o *UpdateMTOServiceItemStatusForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOServiceItemStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOServiceItemStatusNotFound creates a UpdateMTOServiceItemStatusNotFound with default headers values
func NewUpdateMTOServiceItemStatusNotFound() *UpdateMTOServiceItemStatusNotFound {
	return &UpdateMTOServiceItemStatusNotFound{}
}

/*UpdateMTOServiceItemStatusNotFound handles this case with default header values.

The requested resource wasn't found.
*/
type UpdateMTOServiceItemStatusNotFound struct {
	Payload *supportmessages.Error
}

func (o *UpdateMTOServiceItemStatusNotFound) Error() string {
	return fmt.Sprintf("[PATCH /service-items/{mtoServiceItemID}/status][%d] updateMTOServiceItemStatusNotFound  %+v", 404, o.Payload)
}

func (o *UpdateMTOServiceItemStatusNotFound) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *UpdateMTOServiceItemStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOServiceItemStatusConflict creates a UpdateMTOServiceItemStatusConflict with default headers values
func NewUpdateMTOServiceItemStatusConflict() *UpdateMTOServiceItemStatusConflict {
	return &UpdateMTOServiceItemStatusConflict{}
}

/*UpdateMTOServiceItemStatusConflict handles this case with default header values.

Conflict error due to trying to change the status of service item that is not currently "SUBMITTED".
*/
type UpdateMTOServiceItemStatusConflict struct {
	Payload interface{}
}

func (o *UpdateMTOServiceItemStatusConflict) Error() string {
	return fmt.Sprintf("[PATCH /service-items/{mtoServiceItemID}/status][%d] updateMTOServiceItemStatusConflict  %+v", 409, o.Payload)
}

func (o *UpdateMTOServiceItemStatusConflict) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMTOServiceItemStatusConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOServiceItemStatusPreconditionFailed creates a UpdateMTOServiceItemStatusPreconditionFailed with default headers values
func NewUpdateMTOServiceItemStatusPreconditionFailed() *UpdateMTOServiceItemStatusPreconditionFailed {
	return &UpdateMTOServiceItemStatusPreconditionFailed{}
}

/*UpdateMTOServiceItemStatusPreconditionFailed handles this case with default header values.

Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.
*/
type UpdateMTOServiceItemStatusPreconditionFailed struct {
	Payload *supportmessages.Error
}

func (o *UpdateMTOServiceItemStatusPreconditionFailed) Error() string {
	return fmt.Sprintf("[PATCH /service-items/{mtoServiceItemID}/status][%d] updateMTOServiceItemStatusPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateMTOServiceItemStatusPreconditionFailed) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *UpdateMTOServiceItemStatusPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOServiceItemStatusUnprocessableEntity creates a UpdateMTOServiceItemStatusUnprocessableEntity with default headers values
func NewUpdateMTOServiceItemStatusUnprocessableEntity() *UpdateMTOServiceItemStatusUnprocessableEntity {
	return &UpdateMTOServiceItemStatusUnprocessableEntity{}
}

/*UpdateMTOServiceItemStatusUnprocessableEntity handles this case with default header values.

The payload was unprocessable.
*/
type UpdateMTOServiceItemStatusUnprocessableEntity struct {
	Payload *supportmessages.ValidationError
}

func (o *UpdateMTOServiceItemStatusUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /service-items/{mtoServiceItemID}/status][%d] updateMTOServiceItemStatusUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateMTOServiceItemStatusUnprocessableEntity) GetPayload() *supportmessages.ValidationError {
	return o.Payload
}

func (o *UpdateMTOServiceItemStatusUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMTOServiceItemStatusInternalServerError creates a UpdateMTOServiceItemStatusInternalServerError with default headers values
func NewUpdateMTOServiceItemStatusInternalServerError() *UpdateMTOServiceItemStatusInternalServerError {
	return &UpdateMTOServiceItemStatusInternalServerError{}
}

/*UpdateMTOServiceItemStatusInternalServerError handles this case with default header values.

A server error occurred.
*/
type UpdateMTOServiceItemStatusInternalServerError struct {
	Payload *supportmessages.Error
}

func (o *UpdateMTOServiceItemStatusInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /service-items/{mtoServiceItemID}/status][%d] updateMTOServiceItemStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateMTOServiceItemStatusInternalServerError) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *UpdateMTOServiceItemStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}