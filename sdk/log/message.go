package log

import "runtime"

// Message is a log line and a label set that defines it.
type Message interface {
	Caller() runtime.Frame
	Labels() map[string]string
}
