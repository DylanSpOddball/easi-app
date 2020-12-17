// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Intake2 intake 2
//
// swagger:model Intake_2
type Intake2 struct {

	// business case
	// Required: true
	BusinessCase *BusinessCase `json:"BusinessCase"`
}

// Validate validates this intake 2
func (m *Intake2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessCase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Intake2) validateBusinessCase(formats strfmt.Registry) error {

	if err := validate.Required("BusinessCase", "body", m.BusinessCase); err != nil {
		return err
	}

	if m.BusinessCase != nil {
		if err := m.BusinessCase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BusinessCase")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Intake2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Intake2) UnmarshalBinary(b []byte) error {
	var res Intake2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
