package async

import (
	"context"
	"fmt"
	"strconv"
	"sync/atomic"
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestBatch(t *testing.T) {
	its := assert.New(t)

	workItems := 32

	items := make(chan interface{}, workItems)
	for x := 0; x < workItems; x++ {
		items <- "hello" + strconv.Itoa(x)
	}

	var processed int32
	action := func(_ context.Context, v interface{}) error {
		atomic.AddInt32(&processed, 1)
		return fmt.Errorf("this is only a test")
	}

	errors := make(chan error, workItems)
	NewBatch(
		items,
		action,
		OptBatchErrors(errors),
		OptBatchParallelism(4),
	).Process(context.Background())

	its.Equal(workItems, processed)
	its.Equal(workItems, len(errors))
}

func TestBatchPanic(t *testing.T) {
	its := assert.New(t)

	workItems := 32

	items := make(chan interface{}, workItems)
	for x := 0; x < workItems; x++ {
		items <- "hello" + strconv.Itoa(x)
	}

	var processed int32
	action := func(_ context.Context, v interface{}) error {
		if result := atomic.AddInt32(&processed, 1); result == 1 {
			panic("this is only a test")
		}
		return nil
	}

	errors := make(chan error, workItems)
	NewBatch(items, action, OptBatchErrors(errors)).Process(context.Background())

	its.Equal(workItems, processed)
	its.Equal(1, len(errors))
}
