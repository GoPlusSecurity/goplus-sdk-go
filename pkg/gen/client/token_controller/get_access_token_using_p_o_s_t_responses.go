// Code generated by go-swagger; DO NOT EDIT.

package token_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/models"
)

// GetAccessTokenUsingPOSTReader is a Reader for the GetAccessTokenUsingPOST structure.
type GetAccessTokenUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessTokenUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccessTokenUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewGetAccessTokenUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAccessTokenUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccessTokenUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccessTokenUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccessTokenUsingPOSTOK creates a GetAccessTokenUsingPOSTOK with default headers values
func NewGetAccessTokenUsingPOSTOK() *GetAccessTokenUsingPOSTOK {
	return &GetAccessTokenUsingPOSTOK{}
}

/*
GetAccessTokenUsingPOSTOK handles this case with default header values.

OK
*/
type GetAccessTokenUsingPOSTOK struct {
	Payload *models.ResponseWrapperGetAccessTokenResponse
}

func (o *GetAccessTokenUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/token][%d] getAccessTokenUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *GetAccessTokenUsingPOSTOK) GetPayload() *models.ResponseWrapperGetAccessTokenResponse {
	return o.Payload
}

func (o *GetAccessTokenUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperGetAccessTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessTokenUsingPOSTCreated creates a GetAccessTokenUsingPOSTCreated with default headers values
func NewGetAccessTokenUsingPOSTCreated() *GetAccessTokenUsingPOSTCreated {
	return &GetAccessTokenUsingPOSTCreated{}
}

/*
GetAccessTokenUsingPOSTCreated handles this case with default header values.

Created
*/
type GetAccessTokenUsingPOSTCreated struct {
}

func (o *GetAccessTokenUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/token][%d] getAccessTokenUsingPOSTCreated ", 201)
}

func (o *GetAccessTokenUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessTokenUsingPOSTUnauthorized creates a GetAccessTokenUsingPOSTUnauthorized with default headers values
func NewGetAccessTokenUsingPOSTUnauthorized() *GetAccessTokenUsingPOSTUnauthorized {
	return &GetAccessTokenUsingPOSTUnauthorized{}
}

/*
GetAccessTokenUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAccessTokenUsingPOSTUnauthorized struct {
}

func (o *GetAccessTokenUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/token][%d] getAccessTokenUsingPOSTUnauthorized ", 401)
}

func (o *GetAccessTokenUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessTokenUsingPOSTForbidden creates a GetAccessTokenUsingPOSTForbidden with default headers values
func NewGetAccessTokenUsingPOSTForbidden() *GetAccessTokenUsingPOSTForbidden {
	return &GetAccessTokenUsingPOSTForbidden{}
}

/*
GetAccessTokenUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type GetAccessTokenUsingPOSTForbidden struct {
}

func (o *GetAccessTokenUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/token][%d] getAccessTokenUsingPOSTForbidden ", 403)
}

func (o *GetAccessTokenUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccessTokenUsingPOSTNotFound creates a GetAccessTokenUsingPOSTNotFound with default headers values
func NewGetAccessTokenUsingPOSTNotFound() *GetAccessTokenUsingPOSTNotFound {
	return &GetAccessTokenUsingPOSTNotFound{}
}

/*
GetAccessTokenUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type GetAccessTokenUsingPOSTNotFound struct {
}

func (o *GetAccessTokenUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/token][%d] getAccessTokenUsingPOSTNotFound ", 404)
}

func (o *GetAccessTokenUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
