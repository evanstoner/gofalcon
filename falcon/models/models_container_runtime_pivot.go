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

// ModelsContainerRuntimePivot models container runtime pivot
//
// swagger:model models.ContainerRuntimePivot
type ModelsContainerRuntimePivot struct {

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// cloud name
	// Required: true
	CloudName *string `json:"cloud_name"`

	// cluster name
	// Required: true
	ClusterName *string `json:"cluster_name"`

	// container runtime version
	// Required: true
	ContainerRuntimeVersion *string `json:"container_runtime_version"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// first seen
	// Required: true
	FirstSeen *string `json:"first_seen"`

	// last seen
	// Required: true
	LastSeen *string `json:"last_seen"`

	// node name
	// Required: true
	NodeName *string `json:"node_name"`

	// pod name
	// Required: true
	PodName []string `json:"pod_name"`
}

// Validate validates this models container runtime pivot
func (m *ModelsContainerRuntimePivot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerRuntimeVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstSeen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSeen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsContainerRuntimePivot) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *ModelsContainerRuntimePivot) validateCloudName(formats strfmt.Registry) error {

	if err := validate.Required("cloud_name", "body", m.CloudName); err != nil {
		return err
	}

	return nil
}

func (m *ModelsContainerRuntimePivot) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("cluster_name", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

func (m *ModelsContainerRuntimePivot) validateContainerRuntimeVersion(formats strfmt.Registry) error {

	if err := validate.Required("container_runtime_version", "body", m.ContainerRuntimeVersion); err != nil {
		return err
	}

	return nil
}

func (m *ModelsContainerRuntimePivot) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *ModelsContainerRuntimePivot) validateFirstSeen(formats strfmt.Registry) error {

	if err := validate.Required("first_seen", "body", m.FirstSeen); err != nil {
		return err
	}

	return nil
}

func (m *ModelsContainerRuntimePivot) validateLastSeen(formats strfmt.Registry) error {

	if err := validate.Required("last_seen", "body", m.LastSeen); err != nil {
		return err
	}

	return nil
}

func (m *ModelsContainerRuntimePivot) validateNodeName(formats strfmt.Registry) error {

	if err := validate.Required("node_name", "body", m.NodeName); err != nil {
		return err
	}

	return nil
}

func (m *ModelsContainerRuntimePivot) validatePodName(formats strfmt.Registry) error {

	if err := validate.Required("pod_name", "body", m.PodName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models container runtime pivot based on context it is used
func (m *ModelsContainerRuntimePivot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsContainerRuntimePivot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsContainerRuntimePivot) UnmarshalBinary(b []byte) error {
	var res ModelsContainerRuntimePivot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}