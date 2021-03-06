// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostWebhookNotifyReader is a Reader for the PostWebhookNotify structure.
type PostWebhookNotifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWebhookNotifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWebhookNotifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWebhookNotifyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWebhookNotifyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWebhookNotifyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWebhookNotifyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWebhookNotifyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostWebhookNotifyOK creates a PostWebhookNotifyOK with default headers values
func NewPostWebhookNotifyOK() *PostWebhookNotifyOK {
	return &PostWebhookNotifyOK{}
}

/*PostWebhookNotifyOK handles this case with default header values.

Sent
*/
type PostWebhookNotifyOK struct {
	Payload *PostWebhookNotifyOKBody
}

func (o *PostWebhookNotifyOK) Error() string {
	return fmt.Sprintf("[POST /webhook-notify][%d] postWebhookNotifyOK  %+v", 200, o.Payload)
}

func (o *PostWebhookNotifyOK) GetPayload() *PostWebhookNotifyOKBody {
	return o.Payload
}

func (o *PostWebhookNotifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostWebhookNotifyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebhookNotifyBadRequest creates a PostWebhookNotifyBadRequest with default headers values
func NewPostWebhookNotifyBadRequest() *PostWebhookNotifyBadRequest {
	return &PostWebhookNotifyBadRequest{}
}

/*PostWebhookNotifyBadRequest handles this case with default header values.

Bad request
*/
type PostWebhookNotifyBadRequest struct {
}

func (o *PostWebhookNotifyBadRequest) Error() string {
	return fmt.Sprintf("[POST /webhook-notify][%d] postWebhookNotifyBadRequest ", 400)
}

func (o *PostWebhookNotifyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWebhookNotifyUnauthorized creates a PostWebhookNotifyUnauthorized with default headers values
func NewPostWebhookNotifyUnauthorized() *PostWebhookNotifyUnauthorized {
	return &PostWebhookNotifyUnauthorized{}
}

/*PostWebhookNotifyUnauthorized handles this case with default header values.

must be authenticated to use this endpoint
*/
type PostWebhookNotifyUnauthorized struct {
}

func (o *PostWebhookNotifyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /webhook-notify][%d] postWebhookNotifyUnauthorized ", 401)
}

func (o *PostWebhookNotifyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWebhookNotifyForbidden creates a PostWebhookNotifyForbidden with default headers values
func NewPostWebhookNotifyForbidden() *PostWebhookNotifyForbidden {
	return &PostWebhookNotifyForbidden{}
}

/*PostWebhookNotifyForbidden handles this case with default header values.

Forbidden
*/
type PostWebhookNotifyForbidden struct {
}

func (o *PostWebhookNotifyForbidden) Error() string {
	return fmt.Sprintf("[POST /webhook-notify][%d] postWebhookNotifyForbidden ", 403)
}

func (o *PostWebhookNotifyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWebhookNotifyNotFound creates a PostWebhookNotifyNotFound with default headers values
func NewPostWebhookNotifyNotFound() *PostWebhookNotifyNotFound {
	return &PostWebhookNotifyNotFound{}
}

/*PostWebhookNotifyNotFound handles this case with default header values.

No orders found
*/
type PostWebhookNotifyNotFound struct {
}

func (o *PostWebhookNotifyNotFound) Error() string {
	return fmt.Sprintf("[POST /webhook-notify][%d] postWebhookNotifyNotFound ", 404)
}

func (o *PostWebhookNotifyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWebhookNotifyInternalServerError creates a PostWebhookNotifyInternalServerError with default headers values
func NewPostWebhookNotifyInternalServerError() *PostWebhookNotifyInternalServerError {
	return &PostWebhookNotifyInternalServerError{}
}

/*PostWebhookNotifyInternalServerError handles this case with default header values.

Server error
*/
type PostWebhookNotifyInternalServerError struct {
}

func (o *PostWebhookNotifyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /webhook-notify][%d] postWebhookNotifyInternalServerError ", 500)
}

func (o *PostWebhookNotifyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostWebhookNotifyBody post webhook notify body
swagger:model PostWebhookNotifyBody
*/
type PostWebhookNotifyBody struct {

	// Name of event triggered
	EventName string `json:"eventName,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// object
	Object string `json:"object,omitempty"`

	// The type of object that's being updated
	ObjectType string `json:"objectType,omitempty"`

	// Time representing when the event was triggered
	// Format: date-time
	TriggeredAt strfmt.DateTime `json:"triggeredAt,omitempty"`
}

// Validate validates this post webhook notify body
func (o *PostWebhookNotifyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTriggeredAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWebhookNotifyBody) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"id", "body", "uuid", o.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostWebhookNotifyBody) validateTriggeredAt(formats strfmt.Registry) error {

	if swag.IsZero(o.TriggeredAt) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"triggeredAt", "body", "date-time", o.TriggeredAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWebhookNotifyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWebhookNotifyBody) UnmarshalBinary(b []byte) error {
	var res PostWebhookNotifyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWebhookNotifyOKBody post webhook notify o k body
swagger:model PostWebhookNotifyOKBody
*/
type PostWebhookNotifyOKBody struct {

	// Name of event triggered
	EventName string `json:"eventName,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// object
	Object string `json:"object,omitempty"`

	// The type of object that's being updated
	ObjectType string `json:"objectType,omitempty"`

	// Time representing when the event was triggered
	// Format: date-time
	TriggeredAt strfmt.DateTime `json:"triggeredAt,omitempty"`
}

// Validate validates this post webhook notify o k body
func (o *PostWebhookNotifyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTriggeredAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWebhookNotifyOKBody) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("postWebhookNotifyOK"+"."+"id", "body", "uuid", o.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PostWebhookNotifyOKBody) validateTriggeredAt(formats strfmt.Registry) error {

	if swag.IsZero(o.TriggeredAt) { // not required
		return nil
	}

	if err := validate.FormatOf("postWebhookNotifyOK"+"."+"triggeredAt", "body", "date-time", o.TriggeredAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWebhookNotifyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWebhookNotifyOKBody) UnmarshalBinary(b []byte) error {
	var res PostWebhookNotifyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
