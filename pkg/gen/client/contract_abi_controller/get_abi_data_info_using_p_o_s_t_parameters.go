// Code generated by go-swagger; DO NOT EDIT.

package contract_abi_controller

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

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/models"
)

// NewGetAbiDataInfoUsingPOSTParams creates a new GetAbiDataInfoUsingPOSTParams object
// with the default values initialized.
func NewGetAbiDataInfoUsingPOSTParams() *GetAbiDataInfoUsingPOSTParams {
	var ()
	return &GetAbiDataInfoUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAbiDataInfoUsingPOSTParamsWithTimeout creates a new GetAbiDataInfoUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAbiDataInfoUsingPOSTParamsWithTimeout(timeout time.Duration) *GetAbiDataInfoUsingPOSTParams {
	var ()
	return &GetAbiDataInfoUsingPOSTParams{

		timeout: timeout,
	}
}

// NewGetAbiDataInfoUsingPOSTParamsWithContext creates a new GetAbiDataInfoUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAbiDataInfoUsingPOSTParamsWithContext(ctx context.Context) *GetAbiDataInfoUsingPOSTParams {
	var ()
	return &GetAbiDataInfoUsingPOSTParams{

		Context: ctx,
	}
}

// NewGetAbiDataInfoUsingPOSTParamsWithHTTPClient creates a new GetAbiDataInfoUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAbiDataInfoUsingPOSTParamsWithHTTPClient(client *http.Client) *GetAbiDataInfoUsingPOSTParams {
	var ()
	return &GetAbiDataInfoUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
GetAbiDataInfoUsingPOSTParams contains all the parameters to send to the API endpoint
for the get abi data info using p o s t operation typically these are written to a http.Request
*/
type GetAbiDataInfoUsingPOSTParams struct {

	/*Authorization
	  Authorization token in the format: Bearer <token> (e.g., Bearer eyJsZXZlbCI6NSwiYXBwTmFtZSI6ImF2cyIsImFwcEtleSI6IjFaW...)

	*/
	Authorization *string
	/*AbiDataRequest
	  abiDataRequest

	*/
	AbiDataRequest *models.ParseAbiDataRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get abi data info using p o s t params
func (o *GetAbiDataInfoUsingPOSTParams) WithTimeout(timeout time.Duration) *GetAbiDataInfoUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get abi data info using p o s t params
func (o *GetAbiDataInfoUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get abi data info using p o s t params
func (o *GetAbiDataInfoUsingPOSTParams) WithContext(ctx context.Context) *GetAbiDataInfoUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get abi data info using p o s t params
func (o *GetAbiDataInfoUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get abi data info using p o s t params
func (o *GetAbiDataInfoUsingPOSTParams) WithHTTPClient(client *http.Client) *GetAbiDataInfoUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get abi data info using p o s t params
func (o *GetAbiDataInfoUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get abi data info using p o s t params
func (o *GetAbiDataInfoUsingPOSTParams) WithAuthorization(authorization *string) *GetAbiDataInfoUsingPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get abi data info using p o s t params
func (o *GetAbiDataInfoUsingPOSTParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithAbiDataRequest adds the abiDataRequest to the get abi data info using p o s t params
func (o *GetAbiDataInfoUsingPOSTParams) WithAbiDataRequest(abiDataRequest *models.ParseAbiDataRequest) *GetAbiDataInfoUsingPOSTParams {
	o.SetAbiDataRequest(abiDataRequest)
	return o
}

// SetAbiDataRequest adds the abiDataRequest to the get abi data info using p o s t params
func (o *GetAbiDataInfoUsingPOSTParams) SetAbiDataRequest(abiDataRequest *models.ParseAbiDataRequest) {
	o.AbiDataRequest = abiDataRequest
}

// WriteToRequest writes these params to a swagger request
func (o *GetAbiDataInfoUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AbiDataRequest != nil {
		if err := r.SetBodyParam(o.AbiDataRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
