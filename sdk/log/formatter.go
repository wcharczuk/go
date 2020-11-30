package log

import (
	"io"
)

// Formatter formats a given message to a given writer.
type Formatter interface {
	Format(io.Writer, Message)
}
