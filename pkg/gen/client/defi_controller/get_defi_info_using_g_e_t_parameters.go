// Code generated by go-swagger; DO NOT EDIT.

package defi_controller

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
)

// NewGetDefiInfoUsingGETParams creates a new GetDefiInfoUsingGETParams object
// with the default values initialized.
func NewGetDefiInfoUsingGETParams() *GetDefiInfoUsingGETParams {
	var ()
	return &GetDefiInfoUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDefiInfoUsingGETParamsWithTimeout creates a new GetDefiInfoUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDefiInfoUsingGETParamsWithTimeout(timeout time.Duration) *GetDefiInfoUsingGETParams {
	var ()
	return &GetDefiInfoUsingGETParams{

		timeout: timeout,
	}
}

// NewGetDefiInfoUsingGETParamsWithContext creates a new GetDefiInfoUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDefiInfoUsingGETParamsWithContext(ctx context.Context) *GetDefiInfoUsingGETParams {
	var ()
	return &GetDefiInfoUsingGETParams{

		Context: ctx,
	}
}

// NewGetDefiInfoUsingGETParamsWithHTTPClient creates a new GetDefiInfoUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDefiInfoUsingGETParamsWithHTTPClient(client *http.Client) *GetDefiInfoUsingGETParams {
	var ()
	return &GetDefiInfoUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetDefiInfoUsingGETParams contains all the parameters to send to the API endpoint
for the get defi info using g e t operation typically these are written to a http.Request
*/
type GetDefiInfoUsingGETParams struct {

	/*Authorization
	  Authorization (test：Bearer 81|9ihH8JzEuFu4MQ9DjWmH5WrNCPW...)

	*/
	Authorization *string
	/*ChainID
	  Chain id, (eth: 1, bsc: 56)

	*/
	ChainID string
	/*ContractAddresses
	  Defi protocol address

	*/
	ContractAddresses string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get defi info using g e t params
func (o *GetDefiInfoUsingGETParams) WithTimeout(timeout time.Duration) *GetDefiInfoUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get defi info using g e t params
func (o *GetDefiInfoUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get defi info using g e t params
func (o *GetDefiInfoUsingGETParams) WithContext(ctx context.Context) *GetDefiInfoUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get defi info using g e t params
func (o *GetDefiInfoUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get defi info using g e t params
func (o *GetDefiInfoUsingGETParams) WithHTTPClient(client *http.Client) *GetDefiInfoUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get defi info using g e t params
func (o *GetDefiInfoUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get defi info using g e t params
func (o *GetDefiInfoUsingGETParams) WithAuthorization(authorization *string) *GetDefiInfoUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get defi info using g e t params
func (o *GetDefiInfoUsingGETParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithChainID adds the chainID to the get defi info using g e t params
func (o *GetDefiInfoUsingGETParams) WithChainID(chainID string) *GetDefiInfoUsingGETParams {
	o.SetChainID(chainID)
	return o
}

// SetChainID adds the chainId to the get defi info using g e t params
func (o *GetDefiInfoUsingGETParams) SetChainID(chainID string) {
	o.ChainID = chainID
}

// WithContractAddresses adds the contractAddresses to the get defi info using g e t params
func (o *GetDefiInfoUsingGETParams) WithContractAddresses(contractAddresses string) *GetDefiInfoUsingGETParams {
	o.SetContractAddresses(contractAddresses)
	return o
}

// SetContractAddresses adds the contractAddresses to the get defi info using g e t params
func (o *GetDefiInfoUsingGETParams) SetContractAddresses(contractAddresses string) {
	o.ContractAddresses = contractAddresses
}

// WriteToRequest writes these params to a swagger request
func (o *GetDefiInfoUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}

	}

	// path param chain_id
	if err := r.SetPathParam("chain_id", o.ChainID); err != nil {
		return err
	}

	// query param contract_addresses
	qrContractAddresses := o.ContractAddresses
	qContractAddresses := qrContractAddresses
	if qContractAddresses != "" {
		if err := r.SetQueryParam("contract_addresses", qContractAddresses); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
