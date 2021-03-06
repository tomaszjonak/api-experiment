// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tomaszjonak/api-experiment/payments-client/models"
)

// GetPaymentsReader is a Reader for the GetPayments structure.
type GetPaymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetPaymentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPaymentsOK creates a GetPaymentsOK with default headers values
func NewGetPaymentsOK() *GetPaymentsOK {
	return &GetPaymentsOK{}
}

/*GetPaymentsOK handles this case with default header values.

Lists of payments
*/
type GetPaymentsOK struct {
	Payload []*models.Payment
}

func (o *GetPaymentsOK) Error() string {
	return fmt.Sprintf("[GET /payments][%d] getPaymentsOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsDefault creates a GetPaymentsDefault with default headers values
func NewGetPaymentsDefault(code int) *GetPaymentsDefault {
	return &GetPaymentsDefault{
		_statusCode: code,
	}
}

/*GetPaymentsDefault handles this case with default header values.

Error condition
*/
type GetPaymentsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get payments default response
func (o *GetPaymentsDefault) Code() int {
	return o._statusCode
}

func (o *GetPaymentsDefault) Error() string {
	return fmt.Sprintf("[GET /payments][%d] GetPayments default  %+v", o._statusCode, o.Payload)
}

func (o *GetPaymentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
