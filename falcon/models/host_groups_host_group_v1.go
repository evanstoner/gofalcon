// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HostGroupsHostGroupV1 A host group which targets a set of devices either through a dynamic assignment rule or a static set of hosts
//
// swagger:model host_groups.HostGroupV1
type HostGroupsHostGroupV1 struct {

	// The assignment rule of a group
	AssignmentRule string `json:"assignment_rule,omitempty"`

	// The email of the user which created the policy
	// Required: true
	CreatedBy *string `json:"created_by"`

	// The time at which the policy was created
	// Required: true
	// Format: date-time
	CreatedTimestamp *strfmt.DateTime `json:"created_timestamp"`

	// An additional description of the group or the devices it targets
	// Required: true
	Description *string `json:"description"`

	// The method by which this host group is managed
	// Enum: [static dynamic staticByID]
	GroupType string `json:"group_type,omitempty"`

	// The identifier of this host group
	// Required: true
	ID *string `json:"id"`

	// The email of the user which last modified the policy
	// Required: true
	ModifiedBy *string `json:"modified_by"`

	// The time at which the policy was last modified
	// Required: true
	// Format: date-time
	ModifiedTimestamp *strfmt.DateTime `json:"modified_timestamp"`

	// The name of the group
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this host groups host group v1
func (m *HostGroupsHostGroupV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostGroupsHostGroupV1) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("created_by", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *HostGroupsHostGroupV1) validateCreatedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("created_timestamp", "body", m.CreatedTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("created_timestamp", "body", "date-time", m.CreatedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HostGroupsHostGroupV1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

var hostGroupsHostGroupV1TypeGroupTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["static","dynamic","staticByID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostGroupsHostGroupV1TypeGroupTypePropEnum = append(hostGroupsHostGroupV1TypeGroupTypePropEnum, v)
	}
}

const (

	// HostGroupsHostGroupV1GroupTypeStatic captures enum value "static"
	HostGroupsHostGroupV1GroupTypeStatic string = "static"

	// HostGroupsHostGroupV1GroupTypeDynamic captures enum value "dynamic"
	HostGroupsHostGroupV1GroupTypeDynamic string = "dynamic"

	// HostGroupsHostGroupV1GroupTypeStaticByID captures enum value "staticByID"
	HostGroupsHostGroupV1GroupTypeStaticByID string = "staticByID"
)

// prop value enum
func (m *HostGroupsHostGroupV1) validateGroupTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hostGroupsHostGroupV1TypeGroupTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HostGroupsHostGroupV1) validateGroupType(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupType) { // not required
		return nil
	}

	// value enum
	if err := m.validateGroupTypeEnum("group_type", "body", m.GroupType); err != nil {
		return err
	}

	return nil
}

func (m *HostGroupsHostGroupV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *HostGroupsHostGroupV1) validateModifiedBy(formats strfmt.Registry) error {

	if err := validate.Required("modified_by", "body", m.ModifiedBy); err != nil {
		return err
	}

	return nil
}

func (m *HostGroupsHostGroupV1) validateModifiedTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("modified_timestamp", "body", m.ModifiedTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("modified_timestamp", "body", "date-time", m.ModifiedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HostGroupsHostGroupV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this host groups host group v1 based on context it is used
func (m *HostGroupsHostGroupV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostGroupsHostGroupV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostGroupsHostGroupV1) UnmarshalBinary(b []byte) error {
	var res HostGroupsHostGroupV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
