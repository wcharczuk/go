package errors

// StackTraceProvider is a type that can return an exception class.
type StackTraceProvider interface {
	StackTrace() StackTrace
}
