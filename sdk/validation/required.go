package validation

import (
	"errors"
	"fmt"
	"strings"
)

// Error classes.
var (
	ErrRequired = errors.New("required value was not provided")
)

// FieldRequired returns a validation error.
func FieldRequired(fieldName string) error {
	return &Error{
		Err:     ErrRequired,
		Message: fmt.Sprintf("field: %s", fieldName),
	}
}

// FieldsRequired returns a validation error.
func FieldsRequired(fieldNames ...string) error {
	return &Error{
		Err:     ErrRequired,
		Message: fmt.Sprintf("field(s): %s", strings.Join(fieldNames, ", ")),
	}
}
