// Code generated by go-swagger; DO NOT EDIT.

package dapp_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new dapp controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dapp controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetDappInfoUsingGET(params *GetDappInfoUsingGETParams) (*GetDappInfoUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetDappInfoUsingGET checks risk of dapp through URL
*/
func (a *Client) GetDappInfoUsingGET(params *GetDappInfoUsingGETParams) (*GetDappInfoUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDappInfoUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDappInfoUsingGET",
		Method:             "GET",
		PathPattern:        "/api/v1/dapp_security",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDappInfoUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDappInfoUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDappInfoUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
