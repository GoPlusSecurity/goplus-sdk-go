// Code generated by go-swagger; DO NOT EDIT.

package token_controller_v_1

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

// NewGetChainsListUsingGETParams creates a new GetChainsListUsingGETParams object
// with the default values initialized.
func NewGetChainsListUsingGETParams() *GetChainsListUsingGETParams {
	var ()
	return &GetChainsListUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChainsListUsingGETParamsWithTimeout creates a new GetChainsListUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChainsListUsingGETParamsWithTimeout(timeout time.Duration) *GetChainsListUsingGETParams {
	var ()
	return &GetChainsListUsingGETParams{

		timeout: timeout,
	}
}

// NewGetChainsListUsingGETParamsWithContext creates a new GetChainsListUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChainsListUsingGETParamsWithContext(ctx context.Context) *GetChainsListUsingGETParams {
	var ()
	return &GetChainsListUsingGETParams{

		Context: ctx,
	}
}

// NewGetChainsListUsingGETParamsWithHTTPClient creates a new GetChainsListUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChainsListUsingGETParamsWithHTTPClient(client *http.Client) *GetChainsListUsingGETParams {
	var ()
	return &GetChainsListUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetChainsListUsingGETParams contains all the parameters to send to the API endpoint
for the get chains list using g e t operation typically these are written to a http.Request
*/
type GetChainsListUsingGETParams struct {

	/*Authorization
	  Authorization token in the format: Bearer <token> (e.g., Bearer eyJsZXZlbCI6NSwiYXBwTmFtZSI6ImF2cyIsImFwcEtleSI6IjFaW...)

	*/
	Authorization *string
	/*Name
	  API name.

	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get chains list using g e t params
func (o *GetChainsListUsingGETParams) WithTimeout(timeout time.Duration) *GetChainsListUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get chains list using g e t params
func (o *GetChainsListUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get chains list using g e t params
func (o *GetChainsListUsingGETParams) WithContext(ctx context.Context) *GetChainsListUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get chains list using g e t params
func (o *GetChainsListUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get chains list using g e t params
func (o *GetChainsListUsingGETParams) WithHTTPClient(client *http.Client) *GetChainsListUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get chains list using g e t params
func (o *GetChainsListUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get chains list using g e t params
func (o *GetChainsListUsingGETParams) WithAuthorization(authorization *string) *GetChainsListUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get chains list using g e t params
func (o *GetChainsListUsingGETParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithName adds the name to the get chains list using g e t params
func (o *GetChainsListUsingGETParams) WithName(name *string) *GetChainsListUsingGETParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get chains list using g e t params
func (o *GetChainsListUsingGETParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetChainsListUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
