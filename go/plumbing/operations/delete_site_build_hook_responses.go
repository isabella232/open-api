// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// DeleteSiteBuildHookReader is a Reader for the DeleteSiteBuildHook structure.
type DeleteSiteBuildHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSiteBuildHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSiteBuildHookNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSiteBuildHookDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSiteBuildHookNoContent creates a DeleteSiteBuildHookNoContent with default headers values
func NewDeleteSiteBuildHookNoContent() *DeleteSiteBuildHookNoContent {
	return &DeleteSiteBuildHookNoContent{}
}

/*DeleteSiteBuildHookNoContent handles this case with default header values.

No content
*/
type DeleteSiteBuildHookNoContent struct {
}

func (o *DeleteSiteBuildHookNoContent) Error() string {
	return fmt.Sprintf("[DELETE /sites/{site_id}/build_hooks/{id}][%d] deleteSiteBuildHookNoContent ", 204)
}

func (o *DeleteSiteBuildHookNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSiteBuildHookDefault creates a DeleteSiteBuildHookDefault with default headers values
func NewDeleteSiteBuildHookDefault(code int) *DeleteSiteBuildHookDefault {
	return &DeleteSiteBuildHookDefault{
		_statusCode: code,
	}
}

/*DeleteSiteBuildHookDefault handles this case with default header values.

error
*/
type DeleteSiteBuildHookDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete site build hook default response
func (o *DeleteSiteBuildHookDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSiteBuildHookDefault) Error() string {
	return fmt.Sprintf("[DELETE /sites/{site_id}/build_hooks/{id}][%d] deleteSiteBuildHook default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSiteBuildHookDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}