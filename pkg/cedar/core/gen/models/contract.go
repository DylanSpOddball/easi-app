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

// Contract contract
//
// swagger:model Contract
type Contract struct {

	// Contract number
	// Example: HHSM500201600052I
	// Required: true
	AwardID *string `json:"awardId"`

	// budget ids
	BudgetIds []*BudgetIds `json:"budgetIds"`

	// Is ADO Parent Contract, Yes/No
	// Example: yes
	ContractADO string `json:"contractADO,omitempty"`

	// contract deliverable Id
	// Example: 11-22-333
	ContractDeliverableID string `json:"contractDeliverableId,omitempty"`

	// Contract description
	// Example: Strategic partners acquisition readiness
	Description string `json:"description,omitempty"`

	// id
	// Example: 18-3-0
	// Required: true
	ID *string `json:"id"`

	// Parent contract number
	// Example: HHSM500201600052I
	// Required: true
	ParentAwardID *string `json:"parentAwardId"`

	// System which this budget funds
	// Example: 123-45-678
	SystemID string `json:"systemId,omitempty"`

	// tbm cost pool
	TbmCostPool []*ContractTbmCostPoolItems0 `json:"tbmCostPool"`

	// tbm it tower category
	TbmItTowerCategory []*ContractTbmItTowerCategoryItems0 `json:"tbmItTowerCategory"`
}

// Validate validates this contract
func (m *Contract) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBudgetIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentAwardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTbmCostPool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTbmItTowerCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Contract) validateAwardID(formats strfmt.Registry) error {

	if err := validate.Required("awardId", "body", m.AwardID); err != nil {
		return err
	}

	return nil
}

func (m *Contract) validateBudgetIds(formats strfmt.Registry) error {
	if swag.IsZero(m.BudgetIds) { // not required
		return nil
	}

	for i := 0; i < len(m.BudgetIds); i++ {
		if swag.IsZero(m.BudgetIds[i]) { // not required
			continue
		}

		if m.BudgetIds[i] != nil {
			if err := m.BudgetIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("budgetIds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("budgetIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Contract) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Contract) validateParentAwardID(formats strfmt.Registry) error {

	if err := validate.Required("parentAwardId", "body", m.ParentAwardID); err != nil {
		return err
	}

	return nil
}

func (m *Contract) validateTbmCostPool(formats strfmt.Registry) error {
	if swag.IsZero(m.TbmCostPool) { // not required
		return nil
	}

	for i := 0; i < len(m.TbmCostPool); i++ {
		if swag.IsZero(m.TbmCostPool[i]) { // not required
			continue
		}

		if m.TbmCostPool[i] != nil {
			if err := m.TbmCostPool[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tbmCostPool" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tbmCostPool" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Contract) validateTbmItTowerCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.TbmItTowerCategory) { // not required
		return nil
	}

	for i := 0; i < len(m.TbmItTowerCategory); i++ {
		if swag.IsZero(m.TbmItTowerCategory[i]) { // not required
			continue
		}

		if m.TbmItTowerCategory[i] != nil {
			if err := m.TbmItTowerCategory[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tbmItTowerCategory" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tbmItTowerCategory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this contract based on the context it is used
func (m *Contract) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBudgetIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTbmCostPool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTbmItTowerCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Contract) contextValidateBudgetIds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BudgetIds); i++ {

		if m.BudgetIds[i] != nil {
			if err := m.BudgetIds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("budgetIds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("budgetIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Contract) contextValidateTbmCostPool(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TbmCostPool); i++ {

		if m.TbmCostPool[i] != nil {
			if err := m.TbmCostPool[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tbmCostPool" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tbmCostPool" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Contract) contextValidateTbmItTowerCategory(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TbmItTowerCategory); i++ {

		if m.TbmItTowerCategory[i] != nil {
			if err := m.TbmItTowerCategory[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tbmItTowerCategory" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tbmItTowerCategory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Contract) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Contract) UnmarshalBinary(b []byte) error {
	var res Contract
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ContractTbmCostPoolItems0 contract tbm cost pool items0
//
// swagger:model ContractTbmCostPoolItems0
type ContractTbmCostPoolItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this contract tbm cost pool items0
func (m *ContractTbmCostPoolItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this contract tbm cost pool items0 based on context it is used
func (m *ContractTbmCostPoolItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContractTbmCostPoolItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractTbmCostPoolItems0) UnmarshalBinary(b []byte) error {
	var res ContractTbmCostPoolItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ContractTbmItTowerCategoryItems0 contract tbm it tower category items0
//
// swagger:model ContractTbmItTowerCategoryItems0
type ContractTbmItTowerCategoryItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this contract tbm it tower category items0
func (m *ContractTbmItTowerCategoryItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this contract tbm it tower category items0 based on context it is used
func (m *ContractTbmItTowerCategoryItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContractTbmItTowerCategoryItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractTbmItTowerCategoryItems0) UnmarshalBinary(b []byte) error {
	var res ContractTbmItTowerCategoryItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
