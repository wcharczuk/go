package validation

// Error is an error from validation.
type Error struct {
	Err     error
	Message string
}

// Unwrap implements error unwrapping.
func (ve Error) Unwrap() error {
	return ve.Err
}

func (ve Error) Error() string {
	if ve.Message != "" {
		return "validation error: " + ve.Err.Error() + ": " + ve.Message
	}
	return "validation error: " + ve.Err.Error()
}
