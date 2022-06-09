// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostDevice host device
//
// swagger:model host_device
type HostDevice struct {

	// Type of the device
	DeviceType string `json:"device_type,omitempty"`

	// Group id
	Gid int64 `json:"gid,omitempty"`

	// Major of the device
	Major int64 `json:"major,omitempty"`

	// Minor of the device
	Minor int64 `json:"minor,omitempty"`

	// Path of the device
	Path string `json:"path,omitempty"`

	// Owner id
	UID int64 `json:"uid,omitempty"`
}

// Validate validates this host device
func (m *HostDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this host device based on context it is used
func (m *HostDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostDevice) UnmarshalBinary(b []byte) error {
	var res HostDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
