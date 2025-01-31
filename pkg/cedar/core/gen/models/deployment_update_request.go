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

// DeploymentUpdateRequest deployment update request
//
// swagger:model DeploymentUpdateRequest
type DeploymentUpdateRequest struct {

	// deployments
	// Required: true
	Deployments []*Deployment `json:"Deployments"`
}

// Validate validates this deployment update request
func (m *DeploymentUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentUpdateRequest) validateDeployments(formats strfmt.Registry) error {

	if err := validate.Required("Deployments", "body", m.Deployments); err != nil {
		return err
	}

	for i := 0; i < len(m.Deployments); i++ {
		if swag.IsZero(m.Deployments[i]) { // not required
			continue
		}

		if m.Deployments[i] != nil {
			if err := m.Deployments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Deployments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Deployments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this deployment update request based on the context it is used
func (m *DeploymentUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeployments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentUpdateRequest) contextValidateDeployments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Deployments); i++ {

		if m.Deployments[i] != nil {
			if err := m.Deployments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Deployments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Deployments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentUpdateRequest) UnmarshalBinary(b []byte) error {
	var res DeploymentUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
