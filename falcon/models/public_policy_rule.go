// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PublicPolicyRule public policy rule
//
// swagger:model public.PolicyRule
type PublicPolicyRule struct {

	// base path
	BasePath string `json:"base_path,omitempty"`
}

// Validate validates this public policy rule
func (m *PublicPolicyRule) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this public policy rule based on context it is used
func (m *PublicPolicyRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PublicPolicyRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicPolicyRule) UnmarshalBinary(b []byte) error {
	var res PublicPolicyRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}