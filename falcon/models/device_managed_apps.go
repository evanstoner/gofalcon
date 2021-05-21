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

// DeviceManagedApps device managed apps
//
// swagger:model device.ManagedApps
type DeviceManagedApps struct {

	// airlock
	Airlock *DeviceManagedApp `json:"airlock,omitempty"`

	// automox
	Automox *DeviceManagedApp `json:"automox,omitempty"`

	// identity protection
	IdentityProtection *DeviceManagedApp `json:"identity-protection,omitempty"`

	// netskope
	Netskope *DeviceManagedApp `json:"netskope,omitempty"`
}

// Validate validates this device managed apps
func (m *DeviceManagedApps) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAirlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutomox(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityProtection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetskope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceManagedApps) validateAirlock(formats strfmt.Registry) error {
	if swag.IsZero(m.Airlock) { // not required
		return nil
	}

	if m.Airlock != nil {
		if err := m.Airlock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("airlock")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceManagedApps) validateAutomox(formats strfmt.Registry) error {
	if swag.IsZero(m.Automox) { // not required
		return nil
	}

	if m.Automox != nil {
		if err := m.Automox.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("automox")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceManagedApps) validateIdentityProtection(formats strfmt.Registry) error {
	if swag.IsZero(m.IdentityProtection) { // not required
		return nil
	}

	if m.IdentityProtection != nil {
		if err := m.IdentityProtection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity-protection")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceManagedApps) validateNetskope(formats strfmt.Registry) error {
	if swag.IsZero(m.Netskope) { // not required
		return nil
	}

	if m.Netskope != nil {
		if err := m.Netskope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netskope")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this device managed apps based on the context it is used
func (m *DeviceManagedApps) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAirlock(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutomox(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentityProtection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetskope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceManagedApps) contextValidateAirlock(ctx context.Context, formats strfmt.Registry) error {

	if m.Airlock != nil {
		if err := m.Airlock.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("airlock")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceManagedApps) contextValidateAutomox(ctx context.Context, formats strfmt.Registry) error {

	if m.Automox != nil {
		if err := m.Automox.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("automox")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceManagedApps) contextValidateIdentityProtection(ctx context.Context, formats strfmt.Registry) error {

	if m.IdentityProtection != nil {
		if err := m.IdentityProtection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity-protection")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceManagedApps) contextValidateNetskope(ctx context.Context, formats strfmt.Registry) error {

	if m.Netskope != nil {
		if err := m.Netskope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netskope")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceManagedApps) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceManagedApps) UnmarshalBinary(b []byte) error {
	var res DeviceManagedApps
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
