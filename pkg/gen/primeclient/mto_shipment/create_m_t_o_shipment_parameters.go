// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/primemessages"
)

// NewCreateMTOShipmentParams creates a new CreateMTOShipmentParams object
// with the default values initialized.
func NewCreateMTOShipmentParams() *CreateMTOShipmentParams {
	var ()
	return &CreateMTOShipmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMTOShipmentParamsWithTimeout creates a new CreateMTOShipmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMTOShipmentParamsWithTimeout(timeout time.Duration) *CreateMTOShipmentParams {
	var ()
	return &CreateMTOShipmentParams{

		timeout: timeout,
	}
}

// NewCreateMTOShipmentParamsWithContext creates a new CreateMTOShipmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMTOShipmentParamsWithContext(ctx context.Context) *CreateMTOShipmentParams {
	var ()
	return &CreateMTOShipmentParams{

		Context: ctx,
	}
}

// NewCreateMTOShipmentParamsWithHTTPClient creates a new CreateMTOShipmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMTOShipmentParamsWithHTTPClient(client *http.Client) *CreateMTOShipmentParams {
	var ()
	return &CreateMTOShipmentParams{
		HTTPClient: client,
	}
}

/*CreateMTOShipmentParams contains all the parameters to send to the API endpoint
for the create m t o shipment operation typically these are written to a http.Request
*/
type CreateMTOShipmentParams struct {

	/*Body*/
	Body *primemessages.CreateMTOShipment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create m t o shipment params
func (o *CreateMTOShipmentParams) WithTimeout(timeout time.Duration) *CreateMTOShipmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create m t o shipment params
func (o *CreateMTOShipmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create m t o shipment params
func (o *CreateMTOShipmentParams) WithContext(ctx context.Context) *CreateMTOShipmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create m t o shipment params
func (o *CreateMTOShipmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create m t o shipment params
func (o *CreateMTOShipmentParams) WithHTTPClient(client *http.Client) *CreateMTOShipmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create m t o shipment params
func (o *CreateMTOShipmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create m t o shipment params
func (o *CreateMTOShipmentParams) WithBody(body *primemessages.CreateMTOShipment) *CreateMTOShipmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create m t o shipment params
func (o *CreateMTOShipmentParams) SetBody(body *primemessages.CreateMTOShipment) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMTOShipmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
