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

// QuickscanproLaunchScanResource quickscanpro launch scan resource
//
// swagger:model quickscanpro.LaunchScanResource
type QuickscanproLaunchScanResource struct {

	// created timestamp
	// Required: true
	// Format: date-time
	CreatedTimestamp *strfmt.DateTime `json:"created_timestamp"`

	// id
	// Required: true
	ID *string `json:"id"`

	// sha256
	// Required: true
	Sha256 *string `json:"sha256"`
}

// Validate validates this quickscanpro launch scan resource
func (m *QuickscanproLaunchScanResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSha256(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuickscanproLaunchScanResource) validateCreatedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("created_timestamp", "body", m.CreatedTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("created_timestamp", "body", "date-time", m.CreatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QuickscanproLaunchScanResource) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *QuickscanproLaunchScanResource) validateSha256(formats strfmt.Registry) error {

	if err := validate.Required("sha256", "body", m.Sha256); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this quickscanpro launch scan resource based on context it is used
func (m *QuickscanproLaunchScanResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuickscanproLaunchScanResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuickscanproLaunchScanResource) UnmarshalBinary(b []byte) error {
	var res QuickscanproLaunchScanResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}