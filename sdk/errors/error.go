package errors

// Error is a meta interface for exceptions.
type Error interface {
	error
	GetMessage() string
	GetInner() error
	GetStackTrace() StackTrace
}
