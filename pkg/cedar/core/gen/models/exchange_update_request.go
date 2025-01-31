// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExchangeUpdateRequest exchange update request
//
// swagger:model ExchangeUpdateRequest
type ExchangeUpdateRequest struct {

	// exchanges
	// Required: true
	Exchanges []*Exchange `json:"Exchanges"`
}

// Validate validates this exchange update request
func (m *ExchangeUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExchanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeUpdateRequest) validateExchanges(formats strfmt.Registry) error {

	if err := validate.Required("Exchanges", "body", m.Exchanges); err != nil {
		return err
	}

	for i := 0; i < len(m.Exchanges); i++ {
		if swag.IsZero(m.Exchanges[i]) { // not required
			continue
		}

		if m.Exchanges[i] != nil {
			if err := m.Exchanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Exchanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Exchanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this exchange update request based on the context it is used
func (m *ExchangeUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExchanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExchangeUpdateRequest) contextValidateExchanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Exchanges); i++ {

		if m.Exchanges[i] != nil {
			if err := m.Exchanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Exchanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Exchanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeUpdateRequest) UnmarshalBinary(b []byte) error {
	var res ExchangeUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
