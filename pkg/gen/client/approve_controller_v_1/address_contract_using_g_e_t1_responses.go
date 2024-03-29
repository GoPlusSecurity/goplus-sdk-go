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

// AddressContractUsingGET1Reader is a Reader for the AddressContractUsingGET1 structure.
type AddressContractUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressContractUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddressContractUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddressContractUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddressContractUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddressContractUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddressContractUsingGET1OK creates a AddressContractUsingGET1OK with default headers values
func NewAddressContractUsingGET1OK() *AddressContractUsingGET1OK {
	return &AddressContractUsingGET1OK{}
}

/*
AddressContractUsingGET1OK handles this case with default header values.

OK
*/
type AddressContractUsingGET1OK struct {
	Payload *models.ResponseWrapperAddressContract
}

func (o *AddressContractUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /api/v1/address_security/{address}][%d] addressContractUsingGET1OK  %+v", 200, o.Payload)
}

func (o *AddressContractUsingGET1OK) GetPayload() *models.ResponseWrapperAddressContract {
	return o.Payload
}

func (o *AddressContractUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapperAddressContract)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressContractUsingGET1Unauthorized creates a AddressContractUsingGET1Unauthorized with default headers values
func NewAddressContractUsingGET1Unauthorized() *AddressContractUsingGET1Unauthorized {
	return &AddressContractUsingGET1Unauthorized{}
}

/*
AddressContractUsingGET1Unauthorized handles this case with default header values.

Unauthorized
*/
type AddressContractUsingGET1Unauthorized struct {
}

func (o *AddressContractUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/address_security/{address}][%d] addressContractUsingGET1Unauthorized ", 401)
}

func (o *AddressContractUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddressContractUsingGET1Forbidden creates a AddressContractUsingGET1Forbidden with default headers values
func NewAddressContractUsingGET1Forbidden() *AddressContractUsingGET1Forbidden {
	return &AddressContractUsingGET1Forbidden{}
}

/*
AddressContractUsingGET1Forbidden handles this case with default header values.

Forbidden
*/
type AddressContractUsingGET1Forbidden struct {
}

func (o *AddressContractUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/address_security/{address}][%d] addressContractUsingGET1Forbidden ", 403)
}

func (o *AddressContractUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddressContractUsingGET1NotFound creates a AddressContractUsingGET1NotFound with default headers values
func NewAddressContractUsingGET1NotFound() *AddressContractUsingGET1NotFound {
	return &AddressContractUsingGET1NotFound{}
}

/*
AddressContractUsingGET1NotFound handles this case with default header values.

Not Found
*/
type AddressContractUsingGET1NotFound struct {
}

func (o *AddressContractUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/address_security/{address}][%d] addressContractUsingGET1NotFound ", 404)
}

func (o *AddressContractUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
