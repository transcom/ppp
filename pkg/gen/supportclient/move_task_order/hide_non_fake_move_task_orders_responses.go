// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/supportmessages"
)

// HideNonFakeMoveTaskOrdersReader is a Reader for the HideNonFakeMoveTaskOrders structure.
type HideNonFakeMoveTaskOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HideNonFakeMoveTaskOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHideNonFakeMoveTaskOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHideNonFakeMoveTaskOrdersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHideNonFakeMoveTaskOrdersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHideNonFakeMoveTaskOrdersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHideNonFakeMoveTaskOrdersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewHideNonFakeMoveTaskOrdersPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewHideNonFakeMoveTaskOrdersUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHideNonFakeMoveTaskOrdersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHideNonFakeMoveTaskOrdersOK creates a HideNonFakeMoveTaskOrdersOK with default headers values
func NewHideNonFakeMoveTaskOrdersOK() *HideNonFakeMoveTaskOrdersOK {
	return &HideNonFakeMoveTaskOrdersOK{}
}

/*HideNonFakeMoveTaskOrdersOK handles this case with default header values.

Successfully hid MTOs.
*/
type HideNonFakeMoveTaskOrdersOK struct {
	Payload supportmessages.MoveTaskOrders
}

func (o *HideNonFakeMoveTaskOrdersOK) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/hide][%d] hideNonFakeMoveTaskOrdersOK  %+v", 200, o.Payload)
}

func (o *HideNonFakeMoveTaskOrdersOK) GetPayload() supportmessages.MoveTaskOrders {
	return o.Payload
}

func (o *HideNonFakeMoveTaskOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHideNonFakeMoveTaskOrdersBadRequest creates a HideNonFakeMoveTaskOrdersBadRequest with default headers values
func NewHideNonFakeMoveTaskOrdersBadRequest() *HideNonFakeMoveTaskOrdersBadRequest {
	return &HideNonFakeMoveTaskOrdersBadRequest{}
}

/*HideNonFakeMoveTaskOrdersBadRequest handles this case with default header values.

The request payload is invalid.
*/
type HideNonFakeMoveTaskOrdersBadRequest struct {
	Payload *supportmessages.ClientError
}

func (o *HideNonFakeMoveTaskOrdersBadRequest) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/hide][%d] hideNonFakeMoveTaskOrdersBadRequest  %+v", 400, o.Payload)
}

func (o *HideNonFakeMoveTaskOrdersBadRequest) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *HideNonFakeMoveTaskOrdersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHideNonFakeMoveTaskOrdersUnauthorized creates a HideNonFakeMoveTaskOrdersUnauthorized with default headers values
func NewHideNonFakeMoveTaskOrdersUnauthorized() *HideNonFakeMoveTaskOrdersUnauthorized {
	return &HideNonFakeMoveTaskOrdersUnauthorized{}
}

/*HideNonFakeMoveTaskOrdersUnauthorized handles this case with default header values.

The request was denied.
*/
type HideNonFakeMoveTaskOrdersUnauthorized struct {
	Payload *supportmessages.ClientError
}

func (o *HideNonFakeMoveTaskOrdersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/hide][%d] hideNonFakeMoveTaskOrdersUnauthorized  %+v", 401, o.Payload)
}

func (o *HideNonFakeMoveTaskOrdersUnauthorized) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *HideNonFakeMoveTaskOrdersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHideNonFakeMoveTaskOrdersForbidden creates a HideNonFakeMoveTaskOrdersForbidden with default headers values
func NewHideNonFakeMoveTaskOrdersForbidden() *HideNonFakeMoveTaskOrdersForbidden {
	return &HideNonFakeMoveTaskOrdersForbidden{}
}

/*HideNonFakeMoveTaskOrdersForbidden handles this case with default header values.

The request was denied.
*/
type HideNonFakeMoveTaskOrdersForbidden struct {
	Payload *supportmessages.ClientError
}

func (o *HideNonFakeMoveTaskOrdersForbidden) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/hide][%d] hideNonFakeMoveTaskOrdersForbidden  %+v", 403, o.Payload)
}

func (o *HideNonFakeMoveTaskOrdersForbidden) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *HideNonFakeMoveTaskOrdersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHideNonFakeMoveTaskOrdersNotFound creates a HideNonFakeMoveTaskOrdersNotFound with default headers values
func NewHideNonFakeMoveTaskOrdersNotFound() *HideNonFakeMoveTaskOrdersNotFound {
	return &HideNonFakeMoveTaskOrdersNotFound{}
}

/*HideNonFakeMoveTaskOrdersNotFound handles this case with default header values.

The requested resource wasn't found.
*/
type HideNonFakeMoveTaskOrdersNotFound struct {
	Payload *supportmessages.ClientError
}

func (o *HideNonFakeMoveTaskOrdersNotFound) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/hide][%d] hideNonFakeMoveTaskOrdersNotFound  %+v", 404, o.Payload)
}

func (o *HideNonFakeMoveTaskOrdersNotFound) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *HideNonFakeMoveTaskOrdersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHideNonFakeMoveTaskOrdersPreconditionFailed creates a HideNonFakeMoveTaskOrdersPreconditionFailed with default headers values
func NewHideNonFakeMoveTaskOrdersPreconditionFailed() *HideNonFakeMoveTaskOrdersPreconditionFailed {
	return &HideNonFakeMoveTaskOrdersPreconditionFailed{}
}

/*HideNonFakeMoveTaskOrdersPreconditionFailed handles this case with default header values.

Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.
*/
type HideNonFakeMoveTaskOrdersPreconditionFailed struct {
	Payload *supportmessages.ClientError
}

func (o *HideNonFakeMoveTaskOrdersPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/hide][%d] hideNonFakeMoveTaskOrdersPreconditionFailed  %+v", 412, o.Payload)
}

func (o *HideNonFakeMoveTaskOrdersPreconditionFailed) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *HideNonFakeMoveTaskOrdersPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHideNonFakeMoveTaskOrdersUnprocessableEntity creates a HideNonFakeMoveTaskOrdersUnprocessableEntity with default headers values
func NewHideNonFakeMoveTaskOrdersUnprocessableEntity() *HideNonFakeMoveTaskOrdersUnprocessableEntity {
	return &HideNonFakeMoveTaskOrdersUnprocessableEntity{}
}

/*HideNonFakeMoveTaskOrdersUnprocessableEntity handles this case with default header values.

The payload was unprocessable.
*/
type HideNonFakeMoveTaskOrdersUnprocessableEntity struct {
	Payload *supportmessages.ValidationError
}

func (o *HideNonFakeMoveTaskOrdersUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/hide][%d] hideNonFakeMoveTaskOrdersUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *HideNonFakeMoveTaskOrdersUnprocessableEntity) GetPayload() *supportmessages.ValidationError {
	return o.Payload
}

func (o *HideNonFakeMoveTaskOrdersUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHideNonFakeMoveTaskOrdersInternalServerError creates a HideNonFakeMoveTaskOrdersInternalServerError with default headers values
func NewHideNonFakeMoveTaskOrdersInternalServerError() *HideNonFakeMoveTaskOrdersInternalServerError {
	return &HideNonFakeMoveTaskOrdersInternalServerError{}
}

/*HideNonFakeMoveTaskOrdersInternalServerError handles this case with default header values.

A server error occurred.
*/
type HideNonFakeMoveTaskOrdersInternalServerError struct {
	Payload *supportmessages.Error
}

func (o *HideNonFakeMoveTaskOrdersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /move-task-orders/hide][%d] hideNonFakeMoveTaskOrdersInternalServerError  %+v", 500, o.Payload)
}

func (o *HideNonFakeMoveTaskOrdersInternalServerError) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *HideNonFakeMoveTaskOrdersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
