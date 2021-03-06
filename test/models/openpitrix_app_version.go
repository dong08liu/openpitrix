// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixAppVersion openpitrix app version
// swagger:model openpitrixAppVersion
type OpenpitrixAppVersion struct {

	// active or not
	Active bool `json:"active,omitempty"`

	// app id
	AppID string `json:"app_id,omitempty"`

	// the time when app version create
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// description of app of specific version
	Description string `json:"description,omitempty"`

	// home of app of specific version
	Home string `json:"home,omitempty"`

	// icon of app of specific version
	Icon string `json:"icon,omitempty"`

	// keywords of app of specific version
	Keywords string `json:"keywords,omitempty"`

	// maintainers of app of specific version
	Maintainers string `json:"maintainers,omitempty"`

	// message path of app of specific version
	Message string `json:"message,omitempty"`

	// version name
	Name string `json:"name,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// owner path of app of specific version, concat string group_path:user_id
	OwnerPath string `json:"owner_path,omitempty"`

	// package name of app of specific version
	PackageName string `json:"package_name,omitempty"`

	// readme of app of specific version
	Readme string `json:"readme,omitempty"`

	// review id of app of specific version
	ReviewID string `json:"review_id,omitempty"`

	// screenshots of app of specific version
	Screenshots string `json:"screenshots,omitempty"`

	// sequence of app of specific version
	Sequence int64 `json:"sequence,omitempty"`

	// sources of app of specific version
	Sources string `json:"sources,omitempty"`

	// status of app of specific version eg.[draft|submitted|passed|rejected|active|in-review|deleted|suspended]
	Status string `json:"status,omitempty"`

	// record status changed time
	StatusTime strfmt.DateTime `json:"status_time,omitempty"`

	// type of app of specific version
	Type string `json:"type,omitempty"`

	// the time when app version update
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`

	// version id of app
	VersionID string `json:"version_id,omitempty"`
}

// Validate validates this openpitrix app version
func (m *OpenpitrixAppVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixAppVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixAppVersion) UnmarshalBinary(b []byte) error {
	var res OpenpitrixAppVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
