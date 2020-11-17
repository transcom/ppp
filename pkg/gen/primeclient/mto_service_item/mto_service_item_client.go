// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new mto service item API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mto service item API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateMTOServiceItem(params *CreateMTOServiceItemParams) (*CreateMTOServiceItemOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateMTOServiceItem creates m t o service item

  Creates a new instance, or multiple instances, of mtoServiceItem, which come from the list of services that can be provided. Upon creation these items are associated with a Move Task Order and an MTO Shipment. Must include UUIDs for the MTO and MTO Shipment connected to this service item. Some service item types require additional service items to be autogenerated when added - all created service items, autogenerated included, will be returned in the response.

*/
func (a *Client) CreateMTOServiceItem(params *CreateMTOServiceItemParams) (*CreateMTOServiceItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMTOServiceItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createMTOServiceItem",
		Method:             "POST",
		PathPattern:        "/mto-service-items",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateMTOServiceItemReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMTOServiceItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMTOServiceItem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
