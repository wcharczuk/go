package graceful

import (
	"os"
	"syscall"
)

// DefaultSignals are the default os signals to capture to shut down.
var DefaultSignals = []os.Signal{
	os.Interrupt, syscall.SIGTERM,
}
