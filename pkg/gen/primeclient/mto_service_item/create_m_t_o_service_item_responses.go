// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/primemessages"
)

// CreateMTOServiceItemReader is a Reader for the CreateMTOServiceItem structure.
type CreateMTOServiceItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMTOServiceItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMTOServiceItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateMTOServiceItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateMTOServiceItemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateMTOServiceItemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateMTOServiceItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateMTOServiceItemConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateMTOServiceItemUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateMTOServiceItemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateMTOServiceItemOK creates a CreateMTOServiceItemOK with default headers values
func NewCreateMTOServiceItemOK() *CreateMTOServiceItemOK {
	return &CreateMTOServiceItemOK{}
}

/*CreateMTOServiceItemOK handles this case with default header values.

Successfully created an MTO service item.
*/
type CreateMTOServiceItemOK struct {
	Payload []primemessages.MTOServiceItem
}

func (o *CreateMTOServiceItemOK) Error() string {
	return fmt.Sprintf("[POST /mto-service-items][%d] createMTOServiceItemOK  %+v", 200, o.Payload)
}

func (o *CreateMTOServiceItemOK) GetPayload() []primemessages.MTOServiceItem {
	return o.Payload
}

func (o *CreateMTOServiceItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := primemessages.UnmarshalMTOServiceItemSlice(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCreateMTOServiceItemBadRequest creates a CreateMTOServiceItemBadRequest with default headers values
func NewCreateMTOServiceItemBadRequest() *CreateMTOServiceItemBadRequest {
	return &CreateMTOServiceItemBadRequest{}
}

/*CreateMTOServiceItemBadRequest handles this case with default header values.

The request payload is invalid.
*/
type CreateMTOServiceItemBadRequest struct {
	Payload *primemessages.ClientError
}

func (o *CreateMTOServiceItemBadRequest) Error() string {
	return fmt.Sprintf("[POST /mto-service-items][%d] createMTOServiceItemBadRequest  %+v", 400, o.Payload)
}

func (o *CreateMTOServiceItemBadRequest) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateMTOServiceItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMTOServiceItemUnauthorized creates a CreateMTOServiceItemUnauthorized with default headers values
func NewCreateMTOServiceItemUnauthorized() *CreateMTOServiceItemUnauthorized {
	return &CreateMTOServiceItemUnauthorized{}
}

/*CreateMTOServiceItemUnauthorized handles this case with default header values.

The request was denied.
*/
type CreateMTOServiceItemUnauthorized struct {
	Payload *primemessages.ClientError
}

func (o *CreateMTOServiceItemUnauthorized) Error() string {
	return fmt.Sprintf("[POST /mto-service-items][%d] createMTOServiceItemUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateMTOServiceItemUnauthorized) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateMTOServiceItemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMTOServiceItemForbidden creates a CreateMTOServiceItemForbidden with default headers values
func NewCreateMTOServiceItemForbidden() *CreateMTOServiceItemForbidden {
	return &CreateMTOServiceItemForbidden{}
}

/*CreateMTOServiceItemForbidden handles this case with default header values.

The request was denied.
*/
type CreateMTOServiceItemForbidden struct {
	Payload *primemessages.ClientError
}

func (o *CreateMTOServiceItemForbidden) Error() string {
	return fmt.Sprintf("[POST /mto-service-items][%d] createMTOServiceItemForbidden  %+v", 403, o.Payload)
}

func (o *CreateMTOServiceItemForbidden) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateMTOServiceItemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMTOServiceItemNotFound creates a CreateMTOServiceItemNotFound with default headers values
func NewCreateMTOServiceItemNotFound() *CreateMTOServiceItemNotFound {
	return &CreateMTOServiceItemNotFound{}
}

/*CreateMTOServiceItemNotFound handles this case with default header values.

The requested resource wasn't found.
*/
type CreateMTOServiceItemNotFound struct {
	Payload *primemessages.ClientError
}

func (o *CreateMTOServiceItemNotFound) Error() string {
	return fmt.Sprintf("[POST /mto-service-items][%d] createMTOServiceItemNotFound  %+v", 404, o.Payload)
}

func (o *CreateMTOServiceItemNotFound) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateMTOServiceItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMTOServiceItemConflict creates a CreateMTOServiceItemConflict with default headers values
func NewCreateMTOServiceItemConflict() *CreateMTOServiceItemConflict {
	return &CreateMTOServiceItemConflict{}
}

/*CreateMTOServiceItemConflict handles this case with default header values.

The request could not be processed because of conflict in the current state of the resource.
*/
type CreateMTOServiceItemConflict struct {
	Payload *primemessages.ClientError
}

func (o *CreateMTOServiceItemConflict) Error() string {
	return fmt.Sprintf("[POST /mto-service-items][%d] createMTOServiceItemConflict  %+v", 409, o.Payload)
}

func (o *CreateMTOServiceItemConflict) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *CreateMTOServiceItemConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMTOServiceItemUnprocessableEntity creates a CreateMTOServiceItemUnprocessableEntity with default headers values
func NewCreateMTOServiceItemUnprocessableEntity() *CreateMTOServiceItemUnprocessableEntity {
	return &CreateMTOServiceItemUnprocessableEntity{}
}

/*CreateMTOServiceItemUnprocessableEntity handles this case with default header values.

The payload was unprocessable.
*/
type CreateMTOServiceItemUnprocessableEntity struct {
	Payload *primemessages.ValidationError
}

func (o *CreateMTOServiceItemUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /mto-service-items][%d] createMTOServiceItemUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateMTOServiceItemUnprocessableEntity) GetPayload() *primemessages.ValidationError {
	return o.Payload
}

func (o *CreateMTOServiceItemUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMTOServiceItemInternalServerError creates a CreateMTOServiceItemInternalServerError with default headers values
func NewCreateMTOServiceItemInternalServerError() *CreateMTOServiceItemInternalServerError {
	return &CreateMTOServiceItemInternalServerError{}
}

/*CreateMTOServiceItemInternalServerError handles this case with default header values.

A server error occurred.
*/
type CreateMTOServiceItemInternalServerError struct {
	Payload *primemessages.Error
}

func (o *CreateMTOServiceItemInternalServerError) Error() string {
	return fmt.Sprintf("[POST /mto-service-items][%d] createMTOServiceItemInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateMTOServiceItemInternalServerError) GetPayload() *primemessages.Error {
	return o.Payload
}

func (o *CreateMTOServiceItemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
