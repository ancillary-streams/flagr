// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateConstraintRequest create constraint request
//
// swagger:model createConstraintRequest
type CreateConstraintRequest struct {

	// operator
	// Required: true
	// Min Length: 1
	Operator *string `json:"operator"`

	// property
	// Required: true
	// Min Length: 1
	Property *string `json:"property"`

	// value
	// Required: true
	// Min Length: 1
	Value *string `json:"value"`
}

// Validate validates this create constraint request
func (m *CreateConstraintRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateConstraintRequest) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator); err != nil {
		return err
	}

	if err := validate.MinLength("operator", "body", *m.Operator, 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateConstraintRequest) validateProperty(formats strfmt.Registry) error {

	if err := validate.Required("property", "body", m.Property); err != nil {
		return err
	}

	if err := validate.MinLength("property", "body", *m.Property, 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateConstraintRequest) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.MinLength("value", "body", *m.Value, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create constraint request based on context it is used
func (m *CreateConstraintRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateConstraintRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateConstraintRequest) UnmarshalBinary(b []byte) error {
	var res CreateConstraintRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
