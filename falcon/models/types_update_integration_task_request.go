// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TypesUpdateIntegrationTaskRequest types update integration task request
//
// swagger:model types.UpdateIntegrationTaskRequest
type TypesUpdateIntegrationTaskRequest struct {

	// integration task
	IntegrationTask *TypesIntegrationTask `json:"integration_task,omitempty"`
}

// Validate validates this types update integration task request
func (m *TypesUpdateIntegrationTaskRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegrationTask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesUpdateIntegrationTaskRequest) validateIntegrationTask(formats strfmt.Registry) error {
	if swag.IsZero(m.IntegrationTask) { // not required
		return nil
	}

	if m.IntegrationTask != nil {
		if err := m.IntegrationTask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration_task")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integration_task")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this types update integration task request based on the context it is used
func (m *TypesUpdateIntegrationTaskRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntegrationTask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TypesUpdateIntegrationTaskRequest) contextValidateIntegrationTask(ctx context.Context, formats strfmt.Registry) error {

	if m.IntegrationTask != nil {

		if swag.IsZero(m.IntegrationTask) { // not required
			return nil
		}

		if err := m.IntegrationTask.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration_task")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integration_task")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TypesUpdateIntegrationTaskRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TypesUpdateIntegrationTaskRequest) UnmarshalBinary(b []byte) error {
	var res TypesUpdateIntegrationTaskRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}