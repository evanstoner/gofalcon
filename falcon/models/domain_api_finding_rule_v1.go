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

// DomainAPIFindingRuleV1 domain API finding rule v1
//
// swagger:model domain.APIFindingRuleV1
type DomainAPIFindingRuleV1 struct {

	// authority
	Authority string `json:"authority,omitempty"`

	// cce
	Cce string `json:"cce,omitempty"`

	// edited
	// Required: true
	Edited *bool `json:"edited"`

	// group id
	GroupID string `json:"group_id,omitempty"`

	// group name
	GroupName string `json:"group_name,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// last edited timestamp
	LastEditedTimestamp string `json:"last_edited_timestamp,omitempty"`

	// mitre attack tactics
	MitreAttackTactics []*DomainAPIMitreAttackTacticV1 `json:"mitre_attack_tactics"`

	// name
	Name string `json:"name,omitempty"`

	// policy id
	PolicyID string `json:"policy_id,omitempty"`

	// policy name
	PolicyName string `json:"policy_name,omitempty"`

	// recommendation id
	RecommendationID string `json:"recommendation_id,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`
}

// Validate validates this domain API finding rule v1
func (m *DomainAPIFindingRuleV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdited(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMitreAttackTactics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIFindingRuleV1) validateEdited(formats strfmt.Registry) error {

	if err := validate.Required("edited", "body", m.Edited); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIFindingRuleV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainAPIFindingRuleV1) validateMitreAttackTactics(formats strfmt.Registry) error {
	if swag.IsZero(m.MitreAttackTactics) { // not required
		return nil
	}

	for i := 0; i < len(m.MitreAttackTactics); i++ {
		if swag.IsZero(m.MitreAttackTactics[i]) { // not required
			continue
		}

		if m.MitreAttackTactics[i] != nil {
			if err := m.MitreAttackTactics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mitre_attack_tactics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mitre_attack_tactics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this domain API finding rule v1 based on the context it is used
func (m *DomainAPIFindingRuleV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMitreAttackTactics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAPIFindingRuleV1) contextValidateMitreAttackTactics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MitreAttackTactics); i++ {

		if m.MitreAttackTactics[i] != nil {

			if swag.IsZero(m.MitreAttackTactics[i]) { // not required
				return nil
			}

			if err := m.MitreAttackTactics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mitre_attack_tactics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mitre_attack_tactics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainAPIFindingRuleV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAPIFindingRuleV1) UnmarshalBinary(b []byte) error {
	var res DomainAPIFindingRuleV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}