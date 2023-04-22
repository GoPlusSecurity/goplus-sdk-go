// Code generated by go-swagger; DO NOT EDIT.

package nft_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new nft controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for nft controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetNftInfoUsingGET1(params *GetNftInfoUsingGET1Params) (*GetNftInfoUsingGET1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetNftInfoUsingGET1 gets n f t s security and risk data
*/
func (a *Client) GetNftInfoUsingGET1(params *GetNftInfoUsingGET1Params) (*GetNftInfoUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNftInfoUsingGET1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNftInfoUsingGET_1",
		Method:             "GET",
		PathPattern:        "/api/v1/nft_security/{chain_id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNftInfoUsingGET1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNftInfoUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNftInfoUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
