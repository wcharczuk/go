package selector

import (
	"encoding/json"
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestErrorJSON(t *testing.T) {
	// assert that the error can be serialized as json.
	its := assert.New(t)

	testErr := Error("this is only a test")

	contents, err := json.Marshal(testErr)
	its.Nil(err)
	its.Equal("\"this is only a test\"", string(contents))
}
