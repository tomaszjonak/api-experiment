// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/tomaszjonak/api-experiment/rest-api/models"
)

// GetPaymentsOKCode is the HTTP code returned for type GetPaymentsOK
const GetPaymentsOKCode int = 200

/*GetPaymentsOK Lists payments

swagger:response getPaymentsOK
*/
type GetPaymentsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Payment `json:"body,omitempty"`
}

// NewGetPaymentsOK creates GetPaymentsOK with default headers values
func NewGetPaymentsOK() *GetPaymentsOK {

	return &GetPaymentsOK{}
}

// WithPayload adds the payload to the get payments o k response
func (o *GetPaymentsOK) WithPayload(payload []*models.Payment) *GetPaymentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get payments o k response
func (o *GetPaymentsOK) SetPayload(payload []*models.Payment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPaymentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Payment, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}