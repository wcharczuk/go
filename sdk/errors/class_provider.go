package errors

// ClassProvider is a type that supplies an error class.
type ClassProvider interface {
	Class() error
}
