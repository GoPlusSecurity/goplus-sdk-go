// Code generated by go-swagger; DO NOT EDIT.

package token_security_api_for_sui

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

// NewSuiTokenSecurityUsingGETParams creates a new SuiTokenSecurityUsingGETParams object
// with the default values initialized.
func NewSuiTokenSecurityUsingGETParams() *SuiTokenSecurityUsingGETParams {
	var ()
	return &SuiTokenSecurityUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSuiTokenSecurityUsingGETParamsWithTimeout creates a new SuiTokenSecurityUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSuiTokenSecurityUsingGETParamsWithTimeout(timeout time.Duration) *SuiTokenSecurityUsingGETParams {
	var ()
	return &SuiTokenSecurityUsingGETParams{

		timeout: timeout,
	}
}

// NewSuiTokenSecurityUsingGETParamsWithContext creates a new SuiTokenSecurityUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewSuiTokenSecurityUsingGETParamsWithContext(ctx context.Context) *SuiTokenSecurityUsingGETParams {
	var ()
	return &SuiTokenSecurityUsingGETParams{

		Context: ctx,
	}
}

// NewSuiTokenSecurityUsingGETParamsWithHTTPClient creates a new SuiTokenSecurityUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSuiTokenSecurityUsingGETParamsWithHTTPClient(client *http.Client) *SuiTokenSecurityUsingGETParams {
	var ()
	return &SuiTokenSecurityUsingGETParams{
		HTTPClient: client,
	}
}

/*
SuiTokenSecurityUsingGETParams contains all the parameters to send to the API endpoint
for the sui token security using g e t operation typically these are written to a http.Request
*/
type SuiTokenSecurityUsingGETParams struct {

	/*Authorization
	  Authorization token in the format: Bearer <token> (e.g., Bearer eyJsZXZlbCI6NSwiYXBwTmFtZSI6ImF2cyIsImFwcEtleSI6IjFaW...)

	*/
	Authorization *string
	/*ContractAddresses
	  The contract address of sui tokens.

	*/
	ContractAddresses string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sui token security using g e t params
func (o *SuiTokenSecurityUsingGETParams) WithTimeout(timeout time.Duration) *SuiTokenSecurityUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sui token security using g e t params
func (o *SuiTokenSecurityUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sui token security using g e t params
func (o *SuiTokenSecurityUsingGETParams) WithContext(ctx context.Context) *SuiTokenSecurityUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sui token security using g e t params
func (o *SuiTokenSecurityUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sui token security using g e t params
func (o *SuiTokenSecurityUsingGETParams) WithHTTPClient(client *http.Client) *SuiTokenSecurityUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sui token security using g e t params
func (o *SuiTokenSecurityUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the sui token security using g e t params
func (o *SuiTokenSecurityUsingGETParams) WithAuthorization(authorization *string) *SuiTokenSecurityUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the sui token security using g e t params
func (o *SuiTokenSecurityUsingGETParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithContractAddresses adds the contractAddresses to the sui token security using g e t params
func (o *SuiTokenSecurityUsingGETParams) WithContractAddresses(contractAddresses string) *SuiTokenSecurityUsingGETParams {
	o.SetContractAddresses(contractAddresses)
	return o
}

// SetContractAddresses adds the contractAddresses to the sui token security using g e t params
func (o *SuiTokenSecurityUsingGETParams) SetContractAddresses(contractAddresses string) {
	o.ContractAddresses = contractAddresses
}

// WriteToRequest writes these params to a swagger request
func (o *SuiTokenSecurityUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
