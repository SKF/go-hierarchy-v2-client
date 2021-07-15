// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DuplicateNodeResponse duplicate node response
//
// swagger:model DuplicateNodeResponse
type DuplicateNodeResponse struct {

	// new node Id
	NewNodeID string `json:"newNodeId,omitempty"`
}

// Validate validates this duplicate node response
func (m *DuplicateNodeResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DuplicateNodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DuplicateNodeResponse) UnmarshalBinary(b []byte) error {
	var res DuplicateNodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
