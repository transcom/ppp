// Code generated by go-swagger; DO NOT EDIT.

package payment_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new payment request API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payment request API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePaymentRequest(params *CreatePaymentRequestParams) (*CreatePaymentRequestCreated, error)

	CreateUpload(params *CreateUploadParams) (*CreateUploadCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreatePaymentRequest creates payment request

  Creates a new instance of a paymentRequest.
A newly created payment request is assigned the status `PENDING`.
A move task order can have multiple payment requests, and
a final payment request can be marked using boolean `isFinal`.

*/
func (a *Client) CreatePaymentRequest(params *CreatePaymentRequestParams) (*CreatePaymentRequestCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePaymentRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPaymentRequest",
		Method:             "POST",
		PathPattern:        "/payment-requests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePaymentRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePaymentRequestCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPaymentRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateUpload creates upload

  ### Functionality
This endpoint **uploads** a Proof of Service document for a PaymentRequest.

The PaymentRequest should already exist.

PaymentRequests are created with the [createPaymentRequest](#operation/createPaymentRequest) endpoint.

*/
func (a *Client) CreateUpload(params *CreateUploadParams) (*CreateUploadCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUploadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUpload",
		Method:             "POST",
		PathPattern:        "/payment-requests/{paymentRequestID}/uploads",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateUploadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateUploadCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUpload: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
