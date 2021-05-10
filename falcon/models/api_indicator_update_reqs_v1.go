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

// APIIndicatorUpdateReqsV1 api indicator update reqs v1
//
// swagger:model api.IndicatorUpdateReqsV1
type APIIndicatorUpdateReqsV1 struct {

	// bulk update
	// Required: true
	BulkUpdate *APIBulkUpdateReqV1 `json:"bulk_update"`

	// comment
	Comment string `json:"comment,omitempty"`

	// indicators
	// Required: true
	Indicators []*APIIndicatorUpdateReqV1 `json:"indicators"`
}

// Validate validates this api indicator update reqs v1
func (m *APIIndicatorUpdateReqsV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBulkUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndicators(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIIndicatorUpdateReqsV1) validateBulkUpdate(formats strfmt.Registry) error {

	if err := validate.Required("bulk_update", "body", m.BulkUpdate); err != nil {
		return err
	}

	if m.BulkUpdate != nil {
		if err := m.BulkUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bulk_update")
			}
			return err
		}
	}

	return nil
}

func (m *APIIndicatorUpdateReqsV1) validateIndicators(formats strfmt.Registry) error {

	if err := validate.Required("indicators", "body", m.Indicators); err != nil {
		return err
	}

	for i := 0; i < len(m.Indicators); i++ {
		if swag.IsZero(m.Indicators[i]) { // not required
			continue
		}

		if m.Indicators[i] != nil {
			if err := m.Indicators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("indicators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this api indicator update reqs v1 based on the context it is used
func (m *APIIndicatorUpdateReqsV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBulkUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndicators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIIndicatorUpdateReqsV1) contextValidateBulkUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.BulkUpdate != nil {
		if err := m.BulkUpdate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bulk_update")
			}
			return err
		}
	}

	return nil
}

func (m *APIIndicatorUpdateReqsV1) contextValidateIndicators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Indicators); i++ {

		if m.Indicators[i] != nil {
			if err := m.Indicators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("indicators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIIndicatorUpdateReqsV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIIndicatorUpdateReqsV1) UnmarshalBinary(b []byte) error {
	var res APIIndicatorUpdateReqsV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
