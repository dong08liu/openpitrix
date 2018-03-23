// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDetachCredentialFromRuntimeEnvRequset openpitrix detach credential from runtime env requset
// swagger:model openpitrixDetachCredentialFromRuntimeEnvRequset
type OpenpitrixDetachCredentialFromRuntimeEnvRequset struct {

	// runtime env credential id
	RuntimeEnvCredentialID string `json:"runtime_env_credential_id,omitempty"`

	// runtime env id
	RuntimeEnvID string `json:"runtime_env_id,omitempty"`
}

// Validate validates this openpitrix detach credential from runtime env requset
func (m *OpenpitrixDetachCredentialFromRuntimeEnvRequset) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDetachCredentialFromRuntimeEnvRequset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDetachCredentialFromRuntimeEnvRequset) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDetachCredentialFromRuntimeEnvRequset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}