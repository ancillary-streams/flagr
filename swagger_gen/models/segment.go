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

// Segment segment
//
// swagger:model segment
type Segment struct {

	// constraints
	Constraints []*Constraint `json:"constraints"`

	// description
	// Required: true
	// Min Length: 1
	Description *string `json:"description"`

	// distributions
	Distributions []*Distribution `json:"distributions"`

	// id
	// Read Only: true
	// Minimum: 1
	ID int64 `json:"id,omitempty"`

	// rank
	// Required: true
	// Minimum: 0
	Rank *int64 `json:"rank"`

	// rollout percent
	// Required: true
	// Maximum: 100
	// Minimum: 0
	RolloutPercent *int64 `json:"rolloutPercent"`
}

// Validate validates this segment
func (m *Segment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDistributions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRolloutPercent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Segment) validateConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.Constraints) { // not required
		return nil
	}

	for i := 0; i < len(m.Constraints); i++ {
		if swag.IsZero(m.Constraints[i]) { // not required
			continue
		}

		if m.Constraints[i] != nil {
			if err := m.Constraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Segment) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", *m.Description, 1); err != nil {
		return err
	}

	return nil
}

func (m *Segment) validateDistributions(formats strfmt.Registry) error {
	if swag.IsZero(m.Distributions) { // not required
		return nil
	}

	for i := 0; i < len(m.Distributions); i++ {
		if swag.IsZero(m.Distributions[i]) { // not required
			continue
		}

		if m.Distributions[i] != nil {
			if err := m.Distributions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("distributions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Segment) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", m.ID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Segment) validateRank(formats strfmt.Registry) error {

	if err := validate.Required("rank", "body", m.Rank); err != nil {
		return err
	}

	if err := validate.MinimumInt("rank", "body", *m.Rank, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Segment) validateRolloutPercent(formats strfmt.Registry) error {

	if err := validate.Required("rolloutPercent", "body", m.RolloutPercent); err != nil {
		return err
	}

	if err := validate.MinimumInt("rolloutPercent", "body", *m.RolloutPercent, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("rolloutPercent", "body", *m.RolloutPercent, 100, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this segment based on the context it is used
func (m *Segment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDistributions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Segment) contextValidateConstraints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Constraints); i++ {

		if m.Constraints[i] != nil {
			if err := m.Constraints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Segment) contextValidateDistributions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Distributions); i++ {

		if m.Distributions[i] != nil {
			if err := m.Distributions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("distributions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Segment) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Segment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Segment) UnmarshalBinary(b []byte) error {
	var res Segment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
