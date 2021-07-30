// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DetectsQuarantinedFile detects quarantined file
//
// swagger:model detects.QuarantinedFile
type DetectsQuarantinedFile struct {

	// id
	ID string `json:"id,omitempty"`

	// paths
	Paths string `json:"paths,omitempty"`

	// sha256
	Sha256 string `json:"sha256,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this detects quarantined file
func (m *DetectsQuarantinedFile) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this detects quarantined file based on context it is used
func (m *DetectsQuarantinedFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DetectsQuarantinedFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetectsQuarantinedFile) UnmarshalBinary(b []byte) error {
	var res DetectsQuarantinedFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}