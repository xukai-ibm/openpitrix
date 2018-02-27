// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteAppResponse openpitrix delete app response
// swagger:model openpitrixDeleteAppResponse
type OpenpitrixDeleteAppResponse struct {

	// app
	App *OpenpitrixApp `json:"app,omitempty"`
}

// Validate validates this openpitrix delete app response
func (m *OpenpitrixDeleteAppResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApp(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteAppResponse) validateApp(formats strfmt.Registry) error {

	if swag.IsZero(m.App) { // not required
		return nil
	}

	if m.App != nil {

		if err := m.App.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteAppResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteAppResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteAppResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
