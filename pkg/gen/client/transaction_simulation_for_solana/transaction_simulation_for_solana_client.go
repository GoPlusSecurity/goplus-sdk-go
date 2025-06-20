// Code generated by go-swagger; DO NOT EDIT.

package transaction_simulation_for_solana

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new transaction simulation for solana API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for transaction simulation for solana API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	PrerunTxUsingPOST(params *PrerunTxUsingPOSTParams) (*PrerunTxUsingPOSTOK, *PrerunTxUsingPOSTCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PrerunTxUsingPOST checks for potential risks in the transaction
*/
func (a *Client) PrerunTxUsingPOST(params *PrerunTxUsingPOSTParams) (*PrerunTxUsingPOSTOK, *PrerunTxUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrerunTxUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "prerunTxUsingPOST",
		Method:             "POST",
		PathPattern:        "/pis/api/v1/solana/pre_execution",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrerunTxUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PrerunTxUsingPOSTOK:
		return value, nil, nil
	case *PrerunTxUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for transaction_simulation_for_solana: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
