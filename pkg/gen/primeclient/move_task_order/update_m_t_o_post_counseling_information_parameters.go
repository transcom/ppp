// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

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
)

// NewUpdateMTOPostCounselingInformationParams creates a new UpdateMTOPostCounselingInformationParams object
// with the default values initialized.
func NewUpdateMTOPostCounselingInformationParams() *UpdateMTOPostCounselingInformationParams {
	var ()
	return &UpdateMTOPostCounselingInformationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMTOPostCounselingInformationParamsWithTimeout creates a new UpdateMTOPostCounselingInformationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMTOPostCounselingInformationParamsWithTimeout(timeout time.Duration) *UpdateMTOPostCounselingInformationParams {
	var ()
	return &UpdateMTOPostCounselingInformationParams{

		timeout: timeout,
	}
}

// NewUpdateMTOPostCounselingInformationParamsWithContext creates a new UpdateMTOPostCounselingInformationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMTOPostCounselingInformationParamsWithContext(ctx context.Context) *UpdateMTOPostCounselingInformationParams {
	var ()
	return &UpdateMTOPostCounselingInformationParams{

		Context: ctx,
	}
}

// NewUpdateMTOPostCounselingInformationParamsWithHTTPClient creates a new UpdateMTOPostCounselingInformationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMTOPostCounselingInformationParamsWithHTTPClient(client *http.Client) *UpdateMTOPostCounselingInformationParams {
	var ()
	return &UpdateMTOPostCounselingInformationParams{
		HTTPClient: client,
	}
}

/*UpdateMTOPostCounselingInformationParams contains all the parameters to send to the API endpoint
for the update m t o post counseling information operation typically these are written to a http.Request
*/
type UpdateMTOPostCounselingInformationParams struct {

	/*IfMatch*/
	IfMatch string
	/*Body*/
	Body UpdateMTOPostCounselingInformationBody
	/*MoveTaskOrderID
	  ID of move task order to use

	*/
	MoveTaskOrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update m t o post counseling information params
func (o *UpdateMTOPostCounselingInformationParams) WithTimeout(timeout time.Duration) *UpdateMTOPostCounselingInformationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update m t o post counseling information params
func (o *UpdateMTOPostCounselingInformationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update m t o post counseling information params
func (o *UpdateMTOPostCounselingInformationParams) WithContext(ctx context.Context) *UpdateMTOPostCounselingInformationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update m t o post counseling information params
func (o *UpdateMTOPostCounselingInformationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update m t o post counseling information params
func (o *UpdateMTOPostCounselingInformationParams) WithHTTPClient(client *http.Client) *UpdateMTOPostCounselingInformationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update m t o post counseling information params
func (o *UpdateMTOPostCounselingInformationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the update m t o post counseling information params
func (o *UpdateMTOPostCounselingInformationParams) WithIfMatch(ifMatch string) *UpdateMTOPostCounselingInformationParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the update m t o post counseling information params
func (o *UpdateMTOPostCounselingInformationParams) SetIfMatch(ifMatch string) {
	o.IfMatch = ifMatch
}

// WithBody adds the body to the update m t o post counseling information params
func (o *UpdateMTOPostCounselingInformationParams) WithBody(body UpdateMTOPostCounselingInformationBody) *UpdateMTOPostCounselingInformationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update m t o post counseling information params
func (o *UpdateMTOPostCounselingInformationParams) SetBody(body UpdateMTOPostCounselingInformationBody) {
	o.Body = body
}

// WithMoveTaskOrderID adds the moveTaskOrderID to the update m t o post counseling information params
func (o *UpdateMTOPostCounselingInformationParams) WithMoveTaskOrderID(moveTaskOrderID string) *UpdateMTOPostCounselingInformationParams {
	o.SetMoveTaskOrderID(moveTaskOrderID)
	return o
}

// SetMoveTaskOrderID adds the moveTaskOrderId to the update m t o post counseling information params
func (o *UpdateMTOPostCounselingInformationParams) SetMoveTaskOrderID(moveTaskOrderID string) {
	o.MoveTaskOrderID = moveTaskOrderID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMTOPostCounselingInformationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param If-Match
	if err := r.SetHeaderParam("If-Match", o.IfMatch); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param moveTaskOrderID
	if err := r.SetPathParam("moveTaskOrderID", o.MoveTaskOrderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
