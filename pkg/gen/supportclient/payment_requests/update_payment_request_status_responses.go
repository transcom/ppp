// Code generated by go-swagger; DO NOT EDIT.

package payment_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	supportmessages "github.com/transcom/mymove/pkg/gen/supportmessages"
)

// UpdatePaymentRequestStatusReader is a Reader for the UpdatePaymentRequestStatus structure.
type UpdatePaymentRequestStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePaymentRequestStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePaymentRequestStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePaymentRequestStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdatePaymentRequestStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePaymentRequestStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePaymentRequestStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewUpdatePaymentRequestStatusPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdatePaymentRequestStatusUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdatePaymentRequestStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePaymentRequestStatusOK creates a UpdatePaymentRequestStatusOK with default headers values
func NewUpdatePaymentRequestStatusOK() *UpdatePaymentRequestStatusOK {
	return &UpdatePaymentRequestStatusOK{}
}

/*UpdatePaymentRequestStatusOK handles this case with default header values.

Successfully updated payment request status.
*/
type UpdatePaymentRequestStatusOK struct {
	Payload *supportmessages.PaymentRequest
}

func (o *UpdatePaymentRequestStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /payment-requests/{paymentRequestID}/status][%d] updatePaymentRequestStatusOK  %+v", 200, o.Payload)
}

func (o *UpdatePaymentRequestStatusOK) GetPayload() *supportmessages.PaymentRequest {
	return o.Payload
}

func (o *UpdatePaymentRequestStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.PaymentRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePaymentRequestStatusBadRequest creates a UpdatePaymentRequestStatusBadRequest with default headers values
func NewUpdatePaymentRequestStatusBadRequest() *UpdatePaymentRequestStatusBadRequest {
	return &UpdatePaymentRequestStatusBadRequest{}
}

/*UpdatePaymentRequestStatusBadRequest handles this case with default header values.

The parameters were invalid.
*/
type UpdatePaymentRequestStatusBadRequest struct {
	Payload *supportmessages.Error
}

func (o *UpdatePaymentRequestStatusBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /payment-requests/{paymentRequestID}/status][%d] updatePaymentRequestStatusBadRequest  %+v", 400, o.Payload)
}

func (o *UpdatePaymentRequestStatusBadRequest) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *UpdatePaymentRequestStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePaymentRequestStatusUnauthorized creates a UpdatePaymentRequestStatusUnauthorized with default headers values
func NewUpdatePaymentRequestStatusUnauthorized() *UpdatePaymentRequestStatusUnauthorized {
	return &UpdatePaymentRequestStatusUnauthorized{}
}

/*UpdatePaymentRequestStatusUnauthorized handles this case with default header values.

The request was unauthorized.
*/
type UpdatePaymentRequestStatusUnauthorized struct {
	Payload interface{}
}

func (o *UpdatePaymentRequestStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /payment-requests/{paymentRequestID}/status][%d] updatePaymentRequestStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdatePaymentRequestStatusUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdatePaymentRequestStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePaymentRequestStatusForbidden creates a UpdatePaymentRequestStatusForbidden with default headers values
func NewUpdatePaymentRequestStatusForbidden() *UpdatePaymentRequestStatusForbidden {
	return &UpdatePaymentRequestStatusForbidden{}
}

/*UpdatePaymentRequestStatusForbidden handles this case with default header values.

The client doesn't have permissions to perform the request.
*/
type UpdatePaymentRequestStatusForbidden struct {
	Payload interface{}
}

func (o *UpdatePaymentRequestStatusForbidden) Error() string {
	return fmt.Sprintf("[PATCH /payment-requests/{paymentRequestID}/status][%d] updatePaymentRequestStatusForbidden  %+v", 403, o.Payload)
}

func (o *UpdatePaymentRequestStatusForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdatePaymentRequestStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePaymentRequestStatusNotFound creates a UpdatePaymentRequestStatusNotFound with default headers values
func NewUpdatePaymentRequestStatusNotFound() *UpdatePaymentRequestStatusNotFound {
	return &UpdatePaymentRequestStatusNotFound{}
}

/*UpdatePaymentRequestStatusNotFound handles this case with default header values.

The requested resource wasn't found.
*/
type UpdatePaymentRequestStatusNotFound struct {
	Payload *supportmessages.Error
}

func (o *UpdatePaymentRequestStatusNotFound) Error() string {
	return fmt.Sprintf("[PATCH /payment-requests/{paymentRequestID}/status][%d] updatePaymentRequestStatusNotFound  %+v", 404, o.Payload)
}

func (o *UpdatePaymentRequestStatusNotFound) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *UpdatePaymentRequestStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePaymentRequestStatusPreconditionFailed creates a UpdatePaymentRequestStatusPreconditionFailed with default headers values
func NewUpdatePaymentRequestStatusPreconditionFailed() *UpdatePaymentRequestStatusPreconditionFailed {
	return &UpdatePaymentRequestStatusPreconditionFailed{}
}

/*UpdatePaymentRequestStatusPreconditionFailed handles this case with default header values.

Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.
*/
type UpdatePaymentRequestStatusPreconditionFailed struct {
	Payload *supportmessages.Error
}

func (o *UpdatePaymentRequestStatusPreconditionFailed) Error() string {
	return fmt.Sprintf("[PATCH /payment-requests/{paymentRequestID}/status][%d] updatePaymentRequestStatusPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdatePaymentRequestStatusPreconditionFailed) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *UpdatePaymentRequestStatusPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePaymentRequestStatusUnprocessableEntity creates a UpdatePaymentRequestStatusUnprocessableEntity with default headers values
func NewUpdatePaymentRequestStatusUnprocessableEntity() *UpdatePaymentRequestStatusUnprocessableEntity {
	return &UpdatePaymentRequestStatusUnprocessableEntity{}
}

/*UpdatePaymentRequestStatusUnprocessableEntity handles this case with default header values.

The payload was unprocessable.
*/
type UpdatePaymentRequestStatusUnprocessableEntity struct {
	Payload *supportmessages.ValidationError
}

func (o *UpdatePaymentRequestStatusUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /payment-requests/{paymentRequestID}/status][%d] updatePaymentRequestStatusUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdatePaymentRequestStatusUnprocessableEntity) GetPayload() *supportmessages.ValidationError {
	return o.Payload
}

func (o *UpdatePaymentRequestStatusUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePaymentRequestStatusInternalServerError creates a UpdatePaymentRequestStatusInternalServerError with default headers values
func NewUpdatePaymentRequestStatusInternalServerError() *UpdatePaymentRequestStatusInternalServerError {
	return &UpdatePaymentRequestStatusInternalServerError{}
}

/*UpdatePaymentRequestStatusInternalServerError handles this case with default header values.

A server error occurred.
*/
type UpdatePaymentRequestStatusInternalServerError struct {
	Payload *supportmessages.Error
}

func (o *UpdatePaymentRequestStatusInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /payment-requests/{paymentRequestID}/status][%d] updatePaymentRequestStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdatePaymentRequestStatusInternalServerError) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *UpdatePaymentRequestStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}