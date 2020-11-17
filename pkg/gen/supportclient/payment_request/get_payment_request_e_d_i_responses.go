// Code generated by go-swagger; DO NOT EDIT.

package payment_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/supportmessages"
)

// GetPaymentRequestEDIReader is a Reader for the GetPaymentRequestEDI structure.
type GetPaymentRequestEDIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentRequestEDIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPaymentRequestEDIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPaymentRequestEDIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPaymentRequestEDIUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPaymentRequestEDIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPaymentRequestEDINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetPaymentRequestEDIUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPaymentRequestEDIInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentRequestEDIOK creates a GetPaymentRequestEDIOK with default headers values
func NewGetPaymentRequestEDIOK() *GetPaymentRequestEDIOK {
	return &GetPaymentRequestEDIOK{}
}

/*GetPaymentRequestEDIOK handles this case with default header values.

Successfully retrieved payment requests associated with a given move task order
*/
type GetPaymentRequestEDIOK struct {
	Payload *supportmessages.PaymentRequestEDI
}

func (o *GetPaymentRequestEDIOK) Error() string {
	return fmt.Sprintf("[GET /payment-requests/{paymentRequestID}/edi][%d] getPaymentRequestEDIOK  %+v", 200, o.Payload)
}

func (o *GetPaymentRequestEDIOK) GetPayload() *supportmessages.PaymentRequestEDI {
	return o.Payload
}

func (o *GetPaymentRequestEDIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.PaymentRequestEDI)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentRequestEDIBadRequest creates a GetPaymentRequestEDIBadRequest with default headers values
func NewGetPaymentRequestEDIBadRequest() *GetPaymentRequestEDIBadRequest {
	return &GetPaymentRequestEDIBadRequest{}
}

/*GetPaymentRequestEDIBadRequest handles this case with default header values.

The request payload is invalid.
*/
type GetPaymentRequestEDIBadRequest struct {
	Payload *supportmessages.ClientError
}

func (o *GetPaymentRequestEDIBadRequest) Error() string {
	return fmt.Sprintf("[GET /payment-requests/{paymentRequestID}/edi][%d] getPaymentRequestEDIBadRequest  %+v", 400, o.Payload)
}

func (o *GetPaymentRequestEDIBadRequest) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *GetPaymentRequestEDIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentRequestEDIUnauthorized creates a GetPaymentRequestEDIUnauthorized with default headers values
func NewGetPaymentRequestEDIUnauthorized() *GetPaymentRequestEDIUnauthorized {
	return &GetPaymentRequestEDIUnauthorized{}
}

/*GetPaymentRequestEDIUnauthorized handles this case with default header values.

The request was denied.
*/
type GetPaymentRequestEDIUnauthorized struct {
	Payload *supportmessages.ClientError
}

func (o *GetPaymentRequestEDIUnauthorized) Error() string {
	return fmt.Sprintf("[GET /payment-requests/{paymentRequestID}/edi][%d] getPaymentRequestEDIUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPaymentRequestEDIUnauthorized) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *GetPaymentRequestEDIUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentRequestEDIForbidden creates a GetPaymentRequestEDIForbidden with default headers values
func NewGetPaymentRequestEDIForbidden() *GetPaymentRequestEDIForbidden {
	return &GetPaymentRequestEDIForbidden{}
}

/*GetPaymentRequestEDIForbidden handles this case with default header values.

The request was denied.
*/
type GetPaymentRequestEDIForbidden struct {
	Payload *supportmessages.ClientError
}

func (o *GetPaymentRequestEDIForbidden) Error() string {
	return fmt.Sprintf("[GET /payment-requests/{paymentRequestID}/edi][%d] getPaymentRequestEDIForbidden  %+v", 403, o.Payload)
}

func (o *GetPaymentRequestEDIForbidden) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *GetPaymentRequestEDIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentRequestEDINotFound creates a GetPaymentRequestEDINotFound with default headers values
func NewGetPaymentRequestEDINotFound() *GetPaymentRequestEDINotFound {
	return &GetPaymentRequestEDINotFound{}
}

/*GetPaymentRequestEDINotFound handles this case with default header values.

The requested resource wasn't found.
*/
type GetPaymentRequestEDINotFound struct {
	Payload *supportmessages.ClientError
}

func (o *GetPaymentRequestEDINotFound) Error() string {
	return fmt.Sprintf("[GET /payment-requests/{paymentRequestID}/edi][%d] getPaymentRequestEDINotFound  %+v", 404, o.Payload)
}

func (o *GetPaymentRequestEDINotFound) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *GetPaymentRequestEDINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentRequestEDIUnprocessableEntity creates a GetPaymentRequestEDIUnprocessableEntity with default headers values
func NewGetPaymentRequestEDIUnprocessableEntity() *GetPaymentRequestEDIUnprocessableEntity {
	return &GetPaymentRequestEDIUnprocessableEntity{}
}

/*GetPaymentRequestEDIUnprocessableEntity handles this case with default header values.

The payload was unprocessable.
*/
type GetPaymentRequestEDIUnprocessableEntity struct {
	Payload *supportmessages.ValidationError
}

func (o *GetPaymentRequestEDIUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /payment-requests/{paymentRequestID}/edi][%d] getPaymentRequestEDIUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetPaymentRequestEDIUnprocessableEntity) GetPayload() *supportmessages.ValidationError {
	return o.Payload
}

func (o *GetPaymentRequestEDIUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentRequestEDIInternalServerError creates a GetPaymentRequestEDIInternalServerError with default headers values
func NewGetPaymentRequestEDIInternalServerError() *GetPaymentRequestEDIInternalServerError {
	return &GetPaymentRequestEDIInternalServerError{}
}

/*GetPaymentRequestEDIInternalServerError handles this case with default header values.

A server error occurred.
*/
type GetPaymentRequestEDIInternalServerError struct {
	Payload *supportmessages.Error
}

func (o *GetPaymentRequestEDIInternalServerError) Error() string {
	return fmt.Sprintf("[GET /payment-requests/{paymentRequestID}/edi][%d] getPaymentRequestEDIInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPaymentRequestEDIInternalServerError) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *GetPaymentRequestEDIInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
