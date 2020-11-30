package bufferutil

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestPool(t *testing.T) {
	its := assert.New(t)

	pool := NewPool(1024)
	buf := pool.Get()
	its.NotNil(buf)
	its.Equal(1024, buf.Cap())
	its.Zero(buf.Len())
	pool.Put(buf)
}
