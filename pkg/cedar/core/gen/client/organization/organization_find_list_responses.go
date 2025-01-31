// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// OrganizationFindListReader is a Reader for the OrganizationFindList structure.
type OrganizationFindListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationFindListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationFindListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationFindListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationFindListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationFindListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationFindListOK creates a OrganizationFindListOK with default headers values
func NewOrganizationFindListOK() *OrganizationFindListOK {
	return &OrganizationFindListOK{}
}

/* OrganizationFindListOK describes a response with status code 200, with default header values.

OK
*/
type OrganizationFindListOK struct {
	Payload *models.OrganizationFindResponse
}

func (o *OrganizationFindListOK) Error() string {
	return fmt.Sprintf("[GET /organization][%d] organizationFindListOK  %+v", 200, o.Payload)
}
func (o *OrganizationFindListOK) GetPayload() *models.OrganizationFindResponse {
	return o.Payload
}

func (o *OrganizationFindListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationFindResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationFindListBadRequest creates a OrganizationFindListBadRequest with default headers values
func NewOrganizationFindListBadRequest() *OrganizationFindListBadRequest {
	return &OrganizationFindListBadRequest{}
}

/* OrganizationFindListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationFindListBadRequest struct {
	Payload *models.Response
}

func (o *OrganizationFindListBadRequest) Error() string {
	return fmt.Sprintf("[GET /organization][%d] organizationFindListBadRequest  %+v", 400, o.Payload)
}
func (o *OrganizationFindListBadRequest) GetPayload() *models.Response {
	return o.Payload
}

func (o *OrganizationFindListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationFindListUnauthorized creates a OrganizationFindListUnauthorized with default headers values
func NewOrganizationFindListUnauthorized() *OrganizationFindListUnauthorized {
	return &OrganizationFindListUnauthorized{}
}

/* OrganizationFindListUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type OrganizationFindListUnauthorized struct {
	Payload *models.Response
}

func (o *OrganizationFindListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organization][%d] organizationFindListUnauthorized  %+v", 401, o.Payload)
}
func (o *OrganizationFindListUnauthorized) GetPayload() *models.Response {
	return o.Payload
}

func (o *OrganizationFindListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationFindListInternalServerError creates a OrganizationFindListInternalServerError with default headers values
func NewOrganizationFindListInternalServerError() *OrganizationFindListInternalServerError {
	return &OrganizationFindListInternalServerError{}
}

/* OrganizationFindListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type OrganizationFindListInternalServerError struct {
	Payload *models.Response
}

func (o *OrganizationFindListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organization][%d] organizationFindListInternalServerError  %+v", 500, o.Payload)
}
func (o *OrganizationFindListInternalServerError) GetPayload() *models.Response {
	return o.Payload
}

func (o *OrganizationFindListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
