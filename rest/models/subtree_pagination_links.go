// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubtreePaginationLinks subtree pagination links
//
// swagger:model SubtreePaginationLinks
type SubtreePaginationLinks struct {

	// next
	Next string `json:"next,omitempty"`

	// self
	Self string `json:"self,omitempty"`
}

// Validate validates this subtree pagination links
func (m *SubtreePaginationLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubtreePaginationLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubtreePaginationLinks) UnmarshalBinary(b []byte) error {
	var res SubtreePaginationLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}