package async

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func Test_Latch(t *testing.T) {
	its := assert.New(t)

	l := NewLatch()

	var didStart bool
	var didAbort bool
	var didGetWork bool

	work := make(chan bool)
	workComplete := make(chan bool)

	l.Starting()
	its.True(l.IsStarting())
	its.False(l.IsStarted())
	its.False(l.IsStopping())
	its.False(l.IsStopped())

	go func() {
		l.Started()
		didStart = true
		for {
			select {
			case <-work:
				didGetWork = true
				workComplete <- true
			case <-l.NotifyStopping():
				didAbort = true
				l.Stopped()
				return
			}
		}
	}()
	<-l.NotifyStarted()

	work <- true
	its.True(l.IsStarted())

	// wait for work to happen.
	<-workComplete

	// signal stop
	l.Stopping()
	<-l.NotifyStopped()

	its.True(didStart)
	its.True(didAbort)
	its.True(didGetWork)
	its.False(l.IsStopping())
	its.False(l.IsStarted())
	its.True(l.IsStopped())

	didAbort = false
	its.False(didAbort)
	its.False(l.IsStopping())
	its.False(l.IsStarted())
	its.True(l.IsStopped())
}
