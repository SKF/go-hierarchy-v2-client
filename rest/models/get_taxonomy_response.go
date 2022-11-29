// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetTaxonomyResponse get taxonomy response
//
// swagger:model GetTaxonomyResponse
type GetTaxonomyResponse struct {

	// terms
	Terms TaxonomyTerms `json:"terms,omitempty"`
}

// Validate validates this get taxonomy response
func (m *GetTaxonomyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTerms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTaxonomyResponse) validateTerms(formats strfmt.Registry) error {

	if swag.IsZero(m.Terms) { // not required
		return nil
	}

	if err := m.Terms.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("terms")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTaxonomyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTaxonomyResponse) UnmarshalBinary(b []byte) error {
	var res GetTaxonomyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
