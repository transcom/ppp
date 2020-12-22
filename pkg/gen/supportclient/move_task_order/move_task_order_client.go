// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new move task order API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for move task order API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateMoveTaskOrder(params *CreateMoveTaskOrderParams) (*CreateMoveTaskOrderCreated, error)

	GetMoveTaskOrder(params *GetMoveTaskOrderParams) (*GetMoveTaskOrderOK, error)

	HideNonFakeMoveTaskOrders(params *HideNonFakeMoveTaskOrdersParams) (*HideNonFakeMoveTaskOrdersOK, error)

	ListMTOs(params *ListMTOsParams) (*ListMTOsOK, error)

	MakeMoveTaskOrderAvailable(params *MakeMoveTaskOrderAvailableParams) (*MakeMoveTaskOrderAvailableOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateMoveTaskOrder creates move task order

  Creates an instance of moveTaskOrder.
Current this will also create a number of nested objects but not all.
It will currently create
* MoveTaskOrder
* MoveOrder
* Customer
* User
* Entitlement

It will not create addresses, duty stations, shipments, payment requests or service items. It requires an existing contractor ID, destination duty station ID,
origin duty station ID, and an uploaded orders ID to be passed into the request.

This is a support endpoint and will not be available in production.

*/
func (a *Client) CreateMoveTaskOrder(params *CreateMoveTaskOrderParams) (*CreateMoveTaskOrderCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMoveTaskOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createMoveTaskOrder",
		Method:             "POST",
		PathPattern:        "/move-task-orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateMoveTaskOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateMoveTaskOrderCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMoveTaskOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMoveTaskOrder gets move task order

  ### Functionality
This endpoint gets an individual MoveTaskOrder by ID.

It will provide nested information about the Customer and any associated MTOShipments, MTOServiceItems and PaymentRequests.

This is a support endpoint and is not available in production.

*/
func (a *Client) GetMoveTaskOrder(params *GetMoveTaskOrderParams) (*GetMoveTaskOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMoveTaskOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMoveTaskOrder",
		Method:             "GET",
		PathPattern:        "/move-task-orders/{moveTaskOrderID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMoveTaskOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMoveTaskOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMoveTaskOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  HideNonFakeMoveTaskOrders hides non fake move task orders

  Updates move task order without fake user data `show` to false. No request body required. <br />
<br />
This is a support endpoint and will not be available in production.

*/
func (a *Client) HideNonFakeMoveTaskOrders(params *HideNonFakeMoveTaskOrdersParams) (*HideNonFakeMoveTaskOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHideNonFakeMoveTaskOrdersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "hideNonFakeMoveTaskOrders",
		Method:             "PATCH",
		PathPattern:        "/move-task-orders/hide",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HideNonFakeMoveTaskOrdersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HideNonFakeMoveTaskOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for hideNonFakeMoveTaskOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListMTOs lists m t os

  ### Functionality
This endpoint lists all MoveTaskOrders regardless of whether or not they have been made available to Prime.

It will provide nested information about the Customer and any associated MTOShipments, MTOServiceItems and PaymentRequests.

*/
func (a *Client) ListMTOs(params *ListMTOsParams) (*ListMTOsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMTOsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listMTOs",
		Method:             "GET",
		PathPattern:        "/move-task-orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListMTOsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListMTOsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listMTOs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MakeMoveTaskOrderAvailable makes move task order available

  Updates move task order `availableToPrimeAt` to make it available to prime. No request body required. <br />
<br />
This is a support endpoint and will not be available in production.

*/
func (a *Client) MakeMoveTaskOrderAvailable(params *MakeMoveTaskOrderAvailableParams) (*MakeMoveTaskOrderAvailableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMakeMoveTaskOrderAvailableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "makeMoveTaskOrderAvailable",
		Method:             "PATCH",
		PathPattern:        "/move-task-orders/{moveTaskOrderID}/available-to-prime",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MakeMoveTaskOrderAvailableReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MakeMoveTaskOrderAvailableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for makeMoveTaskOrderAvailable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
