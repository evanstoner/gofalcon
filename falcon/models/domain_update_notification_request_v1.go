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

// DomainUpdateNotificationRequestV1 domain update notification request v1
//
// swagger:model domain.UpdateNotificationRequestV1
type DomainUpdateNotificationRequestV1 struct {

	// The unique ID of the user who is assigned to this notification
	// Required: true
	AssignedToUUID *string `json:"assigned_to_uuid"`

	// The ID of the notifications
	// Required: true
	ID *string `json:"id"`

	// The notification status. This can be one of: new, in-progress, closed-false-positive, closed-true-positive.
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this domain update notification request v1
func (m *DomainUpdateNotificationRequestV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignedToUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainUpdateNotificationRequestV1) validateAssignedToUUID(formats strfmt.Registry) error {

	if err := validate.Required("assigned_to_uuid", "body", m.AssignedToUUID); err != nil {
		return err
	}

	return nil
}

func (m *DomainUpdateNotificationRequestV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainUpdateNotificationRequestV1) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain update notification request v1 based on context it is used
func (m *DomainUpdateNotificationRequestV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainUpdateNotificationRequestV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainUpdateNotificationRequestV1) UnmarshalBinary(b []byte) error {
	var res DomainUpdateNotificationRequestV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}