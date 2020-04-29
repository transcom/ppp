// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new move task order API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for move task order API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetMoveTaskOrder gets a move task order by ID

Gets an individual move task order by ID.
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
		ConsumesMediaTypes: []string{""},
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
UpdateMoveTaskOrderStatus updates the status of a move task order to make it available to prime

Updates move task order `isAvailableToPrime` to TRUE to make it available to prime.
*/
func (a *Client) UpdateMoveTaskOrderStatus(params *UpdateMoveTaskOrderStatusParams) (*UpdateMoveTaskOrderStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMoveTaskOrderStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateMoveTaskOrderStatus",
		Method:             "PATCH",
		PathPattern:        "/move-task-orders/{moveTaskOrderID}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateMoveTaskOrderStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateMoveTaskOrderStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateMoveTaskOrderStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
