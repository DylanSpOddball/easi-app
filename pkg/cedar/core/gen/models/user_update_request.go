// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserUpdateRequest user update request
//
// swagger:model UserUpdateRequest
type UserUpdateRequest struct {

	// users
	// Required: true
	Users []*User `json:"Users"`

	// application
	// Required: true
	// Enum: [all alfabet]
	Application *string `json:"application"`
}

// Validate validates this user update request
func (m *UserUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserUpdateRequest) validateUsers(formats strfmt.Registry) error {

	if err := validate.Required("Users", "body", m.Users); err != nil {
		return err
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var userUpdateRequestTypeApplicationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","alfabet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userUpdateRequestTypeApplicationPropEnum = append(userUpdateRequestTypeApplicationPropEnum, v)
	}
}

const (

	// UserUpdateRequestApplicationAll captures enum value "all"
	UserUpdateRequestApplicationAll string = "all"

	// UserUpdateRequestApplicationAlfabet captures enum value "alfabet"
	UserUpdateRequestApplicationAlfabet string = "alfabet"
)

// prop value enum
func (m *UserUpdateRequest) validateApplicationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userUpdateRequestTypeApplicationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserUpdateRequest) validateApplication(formats strfmt.Registry) error {

	if err := validate.Required("application", "body", m.Application); err != nil {
		return err
	}

	// value enum
	if err := m.validateApplicationEnum("application", "body", *m.Application); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user update request based on the context it is used
func (m *UserUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserUpdateRequest) contextValidateUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Users); i++ {

		if m.Users[i] != nil {
			if err := m.Users[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserUpdateRequest) UnmarshalBinary(b []byte) error {
	var res UserUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
