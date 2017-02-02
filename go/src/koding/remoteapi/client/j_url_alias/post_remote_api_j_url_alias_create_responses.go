package j_url_alias

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJURLAliasCreateReader is a Reader for the PostRemoteAPIJURLAliasCreate structure.
type PostRemoteAPIJURLAliasCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJURLAliasCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJURLAliasCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJURLAliasCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJURLAliasCreateOK creates a PostRemoteAPIJURLAliasCreateOK with default headers values
func NewPostRemoteAPIJURLAliasCreateOK() *PostRemoteAPIJURLAliasCreateOK {
	return &PostRemoteAPIJURLAliasCreateOK{}
}

/*PostRemoteAPIJURLAliasCreateOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJURLAliasCreateOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJURLAliasCreateOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUrlAlias.create][%d] postRemoteApiJUrlAliasCreateOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJURLAliasCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJURLAliasCreateUnauthorized creates a PostRemoteAPIJURLAliasCreateUnauthorized with default headers values
func NewPostRemoteAPIJURLAliasCreateUnauthorized() *PostRemoteAPIJURLAliasCreateUnauthorized {
	return &PostRemoteAPIJURLAliasCreateUnauthorized{}
}

/*PostRemoteAPIJURLAliasCreateUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJURLAliasCreateUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJURLAliasCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUrlAlias.create][%d] postRemoteApiJUrlAliasCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJURLAliasCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
