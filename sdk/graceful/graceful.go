package graceful

import "context"

// Graceful is a server that can start and stop.
type Graceful interface {
	Start() error // this call must block
	Stop(context.Context) error
}
