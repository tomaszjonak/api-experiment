// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PartyDescription party description
// swagger:model PartyDescription
type PartyDescription struct {

	// account name
	AccountName string `json:"account_name,omitempty"`

	// account number
	// Required: true
	AccountNumber *string `json:"account_number"`

	// account number code
	AccountNumberCode string `json:"account_number_code,omitempty"`

	// account type
	AccountType int32 `json:"account_type,omitempty"`

	// address
	Address string `json:"address,omitempty"`

	// bank id
	// Required: true
	BankID *string `json:"bank_id"`

	// bank id code
	// Required: true
	BankIDCode *string `json:"bank_id_code"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this party description
func (m *PartyDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankIDCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyDescription) validateAccountNumber(formats strfmt.Registry) error {

	if err := validate.Required("account_number", "body", m.AccountNumber); err != nil {
		return err
	}

	return nil
}

func (m *PartyDescription) validateBankID(formats strfmt.Registry) error {

	if err := validate.Required("bank_id", "body", m.BankID); err != nil {
		return err
	}

	return nil
}

func (m *PartyDescription) validateBankIDCode(formats strfmt.Registry) error {

	if err := validate.Required("bank_id_code", "body", m.BankIDCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyDescription) UnmarshalBinary(b []byte) error {
	var res PartyDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
