// Code generated by go-swagger; DO NOT EDIT.

package adminmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateWebhookSubscription create webhook subscription
//
// swagger:model CreateWebhookSubscription
type CreateWebhookSubscription struct {

	// The URL to which the notifications for this subscription will be pushed to.
	// Required: true
	CallbackURL *string `json:"callbackUrl"`

	// A string used to represent which events this subscriber expects to be notified about. Corresponds to the possible event_key values in webhook_notifications.
	// Required: true
	EventKey *string `json:"eventKey"`

	// status
	// Required: true
	Status WebhookSubscriptionStatus `json:"status"`

	// subscriber Id
	// Required: true
	// Format: uuid
	SubscriberID *strfmt.UUID `json:"subscriberId"`
}

// Validate validates this create webhook subscription
func (m *CreateWebhookSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallbackURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateWebhookSubscription) validateCallbackURL(formats strfmt.Registry) error {

	if err := validate.Required("callbackUrl", "body", m.CallbackURL); err != nil {
		return err
	}

	return nil
}

func (m *CreateWebhookSubscription) validateEventKey(formats strfmt.Registry) error {

	if err := validate.Required("eventKey", "body", m.EventKey); err != nil {
		return err
	}

	return nil
}

func (m *CreateWebhookSubscription) validateStatus(formats strfmt.Registry) error {

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *CreateWebhookSubscription) validateSubscriberID(formats strfmt.Registry) error {

	if err := validate.Required("subscriberId", "body", m.SubscriberID); err != nil {
		return err
	}

	if err := validate.FormatOf("subscriberId", "body", "uuid", m.SubscriberID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateWebhookSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateWebhookSubscription) UnmarshalBinary(b []byte) error {
	var res CreateWebhookSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
