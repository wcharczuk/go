package graceful

import (
	"context"
	"os"
)

// OptDefaultSignal returns an option that sets the shutdown signal to the defaults.
func OptDefaultSignal() Option {
	return func(o *Options) { o.Signal = Notify(DefaultSignals...) }
}

// OptSignal sets the shutdown signal.
func OptSignal(signal chan os.Signal) Option {
	return func(so *Options) { so.Signal = signal }
}

// OptSignalAll sets if we should handle any instances
// of the signal(s), or just the first.
func OptSignalAll(signalAll bool) Option {
	return func(so *Options) { so.SignalAll = signalAll }
}

// Option is a mutator for shutdown options.
type Option func(*Options)

// Options are the options for graceful shutdown.
type Options struct {
	Signal    chan os.Signal
	SignalAll bool
	Context   context.Context
	Cancel    func()
}
