// Code generated by go-swagger; DO NOT EDIT.

package approve_controller_v_1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/models"
)

// ApprovalContractUsingGETReader is a Reader for the ApprovalContractUsingGET structure.
type ApprovalContractUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApprovalContractUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApprovalContractUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewApprovalContractUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewApprovalContractUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewApprovalContractUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewApprovalContractUsingGETOK creates a ApprovalContractUsingGETOK with default headers values
func NewApprovalContractUsingGETOK() *ApprovalContractUsingGETOK {
	return &ApprovalContractUsingGETOK{}
}

/*
ApprovalContractUsingGETOK handles this case with default header values.

OK
*/
type ApprovalContractUsingGETOK struct {
	Payload *models.ResponseWrapperContractApproveResponse
}

func (o *ApprovalContractUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/approval_security/{chain_id}][%d] approvalContractUsingGETOK  %+v", 200, o.Payload)
}

func (o *ApprovalContractUsingGETOK) GetPayload() *models.ResponseWrapperContractApproveResponse {
	return o.Payload
}

func (o *ApprovalContractUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperContractApproveResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApprovalContractUsingGETUnauthorized creates a ApprovalContractUsingGETUnauthorized with default headers values
func NewApprovalContractUsingGETUnauthorized() *ApprovalContractUsingGETUnauthorized {
	return &ApprovalContractUsingGETUnauthorized{}
}

/*
ApprovalContractUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type ApprovalContractUsingGETUnauthorized struct {
}

func (o *ApprovalContractUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/approval_security/{chain_id}][%d] approvalContractUsingGETUnauthorized ", 401)
}

func (o *ApprovalContractUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApprovalContractUsingGETForbidden creates a ApprovalContractUsingGETForbidden with default headers values
func NewApprovalContractUsingGETForbidden() *ApprovalContractUsingGETForbidden {
	return &ApprovalContractUsingGETForbidden{}
}

/*
ApprovalContractUsingGETForbidden handles this case with default header values.

Forbidden
*/
type ApprovalContractUsingGETForbidden struct {
}

func (o *ApprovalContractUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/approval_security/{chain_id}][%d] approvalContractUsingGETForbidden ", 403)
}

func (o *ApprovalContractUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewApprovalContractUsingGETNotFound creates a ApprovalContractUsingGETNotFound with default headers values
func NewApprovalContractUsingGETNotFound() *ApprovalContractUsingGETNotFound {
	return &ApprovalContractUsingGETNotFound{}
}

/*
ApprovalContractUsingGETNotFound handles this case with default header values.

Not Found
*/
type ApprovalContractUsingGETNotFound struct {
}

func (o *ApprovalContractUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/approval_security/{chain_id}][%d] approvalContractUsingGETNotFound ", 404)
}

func (o *ApprovalContractUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
