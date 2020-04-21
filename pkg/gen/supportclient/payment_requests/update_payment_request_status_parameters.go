// Code generated by go-swagger; DO NOT EDIT.

package payment_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	supportmessages "github.com/transcom/mymove/pkg/gen/supportmessages"
)

// NewUpdatePaymentRequestStatusParams creates a new UpdatePaymentRequestStatusParams object
// with the default values initialized.
func NewUpdatePaymentRequestStatusParams() *UpdatePaymentRequestStatusParams {
	var ()
	return &UpdatePaymentRequestStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePaymentRequestStatusParamsWithTimeout creates a new UpdatePaymentRequestStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePaymentRequestStatusParamsWithTimeout(timeout time.Duration) *UpdatePaymentRequestStatusParams {
	var ()
	return &UpdatePaymentRequestStatusParams{

		timeout: timeout,
	}
}

// NewUpdatePaymentRequestStatusParamsWithContext creates a new UpdatePaymentRequestStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePaymentRequestStatusParamsWithContext(ctx context.Context) *UpdatePaymentRequestStatusParams {
	var ()
	return &UpdatePaymentRequestStatusParams{

		Context: ctx,
	}
}

// NewUpdatePaymentRequestStatusParamsWithHTTPClient creates a new UpdatePaymentRequestStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePaymentRequestStatusParamsWithHTTPClient(client *http.Client) *UpdatePaymentRequestStatusParams {
	var ()
	return &UpdatePaymentRequestStatusParams{
		HTTPClient: client,
	}
}

/*UpdatePaymentRequestStatusParams contains all the parameters to send to the API endpoint
for the update payment request status operation typically these are written to a http.Request
*/
type UpdatePaymentRequestStatusParams struct {

	/*IfMatch*/
	IfMatch string
	/*Body*/
	Body *supportmessages.UpdatePaymentRequestStatus
	/*PaymentRequestID
	  UUID of payment request

	*/
	PaymentRequestID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update payment request status params
func (o *UpdatePaymentRequestStatusParams) WithTimeout(timeout time.Duration) *UpdatePaymentRequestStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update payment request status params
func (o *UpdatePaymentRequestStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update payment request status params
func (o *UpdatePaymentRequestStatusParams) WithContext(ctx context.Context) *UpdatePaymentRequestStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update payment request status params
func (o *UpdatePaymentRequestStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update payment request status params
func (o *UpdatePaymentRequestStatusParams) WithHTTPClient(client *http.Client) *UpdatePaymentRequestStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update payment request status params
func (o *UpdatePaymentRequestStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the update payment request status params
func (o *UpdatePaymentRequestStatusParams) WithIfMatch(ifMatch string) *UpdatePaymentRequestStatusParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the update payment request status params
func (o *UpdatePaymentRequestStatusParams) SetIfMatch(ifMatch string) {
	o.IfMatch = ifMatch
}

// WithBody adds the body to the update payment request status params
func (o *UpdatePaymentRequestStatusParams) WithBody(body *supportmessages.UpdatePaymentRequestStatus) *UpdatePaymentRequestStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update payment request status params
func (o *UpdatePaymentRequestStatusParams) SetBody(body *supportmessages.UpdatePaymentRequestStatus) {
	o.Body = body
}

// WithPaymentRequestID adds the paymentRequestID to the update payment request status params
func (o *UpdatePaymentRequestStatusParams) WithPaymentRequestID(paymentRequestID strfmt.UUID) *UpdatePaymentRequestStatusParams {
	o.SetPaymentRequestID(paymentRequestID)
	return o
}

// SetPaymentRequestID adds the paymentRequestId to the update payment request status params
func (o *UpdatePaymentRequestStatusParams) SetPaymentRequestID(paymentRequestID strfmt.UUID) {
	o.PaymentRequestID = paymentRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePaymentRequestStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param If-Match
	if err := r.SetHeaderParam("If-Match", o.IfMatch); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param paymentRequestID
	if err := r.SetPathParam("paymentRequestID", o.PaymentRequestID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
