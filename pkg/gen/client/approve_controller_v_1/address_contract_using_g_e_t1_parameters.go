// Code generated by go-swagger; DO NOT EDIT.

package approve_controller_v_1

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

// NewAddressContractUsingGET1Params creates a new AddressContractUsingGET1Params object
// with the default values initialized.
func NewAddressContractUsingGET1Params() *AddressContractUsingGET1Params {
	var ()
	return &AddressContractUsingGET1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddressContractUsingGET1ParamsWithTimeout creates a new AddressContractUsingGET1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddressContractUsingGET1ParamsWithTimeout(timeout time.Duration) *AddressContractUsingGET1Params {
	var ()
	return &AddressContractUsingGET1Params{

		timeout: timeout,
	}
}

// NewAddressContractUsingGET1ParamsWithContext creates a new AddressContractUsingGET1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAddressContractUsingGET1ParamsWithContext(ctx context.Context) *AddressContractUsingGET1Params {
	var ()
	return &AddressContractUsingGET1Params{

		Context: ctx,
	}
}

// NewAddressContractUsingGET1ParamsWithHTTPClient creates a new AddressContractUsingGET1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddressContractUsingGET1ParamsWithHTTPClient(client *http.Client) *AddressContractUsingGET1Params {
	var ()
	return &AddressContractUsingGET1Params{
		HTTPClient: client,
	}
}

/*
AddressContractUsingGET1Params contains all the parameters to send to the API endpoint
for the address contract using g e t 1 operation typically these are written to a http.Request
*/
type AddressContractUsingGET1Params struct {

	/*Authorization
	  Authorization (test：Bearer 81|9ihH8JzEuFu4MQ9DjWmH5WrNCPW1zQ9cCv8WrbB1)

	*/
	Authorization *string
	/*Address
	  address

	*/
	Address string
	/*ChainID
	  The chain_id of the blockchain.
	"1" means Ethereum;
	"10" means Optimism;
	“25” means Cronos;
	"56" means BSC;
	“66” means OKC;
	"100" means Gnosis;
	"128" means HECO;
	"137" means Polygon;
	"250" means Fantom;
	"321" means KCC;
	"324" means zkSync Era;
	"10001" means ETHW;
	"201022" means FON;
	"42161" means Arbitrum;
	"43114" means Avalanche;
	"59140" means Linea;
	"1666600000" means Harmony;
	"tron" means Tron.

	*/
	ChainID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the address contract using g e t 1 params
func (o *AddressContractUsingGET1Params) WithTimeout(timeout time.Duration) *AddressContractUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the address contract using g e t 1 params
func (o *AddressContractUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the address contract using g e t 1 params
func (o *AddressContractUsingGET1Params) WithContext(ctx context.Context) *AddressContractUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the address contract using g e t 1 params
func (o *AddressContractUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the address contract using g e t 1 params
func (o *AddressContractUsingGET1Params) WithHTTPClient(client *http.Client) *AddressContractUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the address contract using g e t 1 params
func (o *AddressContractUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the address contract using g e t 1 params
func (o *AddressContractUsingGET1Params) WithAuthorization(authorization *string) *AddressContractUsingGET1Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the address contract using g e t 1 params
func (o *AddressContractUsingGET1Params) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithAddress adds the address to the address contract using g e t 1 params
func (o *AddressContractUsingGET1Params) WithAddress(address string) *AddressContractUsingGET1Params {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the address contract using g e t 1 params
func (o *AddressContractUsingGET1Params) SetAddress(address string) {
	o.Address = address
}

// WithChainID adds the chainID to the address contract using g e t 1 params
func (o *AddressContractUsingGET1Params) WithChainID(chainID *string) *AddressContractUsingGET1Params {
	o.SetChainID(chainID)
	return o
}

// SetChainID adds the chainId to the address contract using g e t 1 params
func (o *AddressContractUsingGET1Params) SetChainID(chainID *string) {
	o.ChainID = chainID
}

// WriteToRequest writes these params to a swagger request
func (o *AddressContractUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if o.ChainID != nil {

		// query param chain_id
		var qrChainID string
		if o.ChainID != nil {
			qrChainID = *o.ChainID
		}
		qChainID := qrChainID
		if qChainID != "" {
			if err := r.SetQueryParam("chain_id", qChainID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
