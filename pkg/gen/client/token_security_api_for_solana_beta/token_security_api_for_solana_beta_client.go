// Code generated by go-swagger; DO NOT EDIT.

package token_security_api_for_solana_beta

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new token security api for solana beta API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for token security api for solana beta API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	SolanaTokenSecurityUsingGET(params *SolanaTokenSecurityUsingGETParams) (*SolanaTokenSecurityUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
SolanaTokenSecurityUsingGET gets token s security and risk data
*/
func (a *Client) SolanaTokenSecurityUsingGET(params *SolanaTokenSecurityUsingGETParams) (*SolanaTokenSecurityUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSolanaTokenSecurityUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "solanaTokenSecurityUsingGET",
		Method:             "GET",
		PathPattern:        "/api/v1/solana/token_security",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SolanaTokenSecurityUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SolanaTokenSecurityUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for solanaTokenSecurityUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
