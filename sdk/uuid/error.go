package uuid

// Error is a type alias to string.
type Error string

// Error implements error.
func (e Error) Error() string {
	return string(e)
}

// ParseError is an error that occured during parsing.
type ParseError struct {
	Message string
	Err     error
}

// Unwrap implements unwrapper.
func (pe ParseError) Unwrap() error {
	return pe.Err
}

// Error implements error.
func (pe ParseError) Error() string {
	return "uuid parse error: " + pe.Message + ": " + pe.Err.Error()
}
