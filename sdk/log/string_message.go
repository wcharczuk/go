package log

import "runtime"

var (
	_ Message = (*StringMessage)(nil)
)

// StringMessage is a simple message that represents a string.
type StringMessage struct {
	MessageCaller runtime.Frame
	Message       string
	MessageLabels map[string]string
}

// Caller implements Message.
func (sm StringMessage) Caller() runtime.Frame {
	return sm.MessageCaller
}

// Labels returns nothing.
func (sm StringMessage) Labels() map[string]string {
	return sm.MessageLabels
}

// String implements fmt.Stringer
func (sm StringMessage) String() string {
	return sm.Message
}
