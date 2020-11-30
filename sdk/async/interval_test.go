package async

import (
	"context"
	"testing"
	"time"

	"go.charczuk.com/sdk/assert"
)

func TestIntervalWorker(t *testing.T) {
	its := assert.New(t)

	var didWork bool
	unbuffered := make(chan bool)
	w := NewInterval(func(_ context.Context) error {
		didWork = true
		<-unbuffered
		return nil
	}, time.Millisecond)

	its.Equal(time.Millisecond, w.Interval)

	go func() { _ = w.Start() }()
	<-w.NotifyStarted()

	its.True(w.IsStarted())
	unbuffered <- true
	close(unbuffered)
	its.Nil(w.Stop())
	its.True(w.IsStopped())
	its.True(didWork)
}
