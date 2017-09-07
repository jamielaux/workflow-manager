package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Retrier retrier
// swagger:model Retrier
type Retrier struct {

	// error equals
	ErrorEquals []ErrorEquals `json:"errorEquals"`

	// max attempts
	// Maximum: 10
	// Minimum: 0
	MaxAttempts *int64 `json:"maxAttempts,omitempty"`
}

// Validate validates this retrier
func (m *Retrier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorEquals(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaxAttempts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Retrier) validateErrorEquals(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorEquals) { // not required
		return nil
	}

	return nil
}

func (m *Retrier) validateMaxAttempts(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxAttempts) { // not required
		return nil
	}

	if err := validate.MinimumInt("maxAttempts", "body", int64(*m.MaxAttempts), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("maxAttempts", "body", int64(*m.MaxAttempts), 10, false); err != nil {
		return err
	}

	return nil
}
