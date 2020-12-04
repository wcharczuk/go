package errors

import (
	"encoding/json"
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestClassMarshalJSON(t *testing.T) {
	its := assert.New(t)

	err := Class("this is only a test")
	contents, marshalErr := json.Marshal(err)
	its.Nil(marshalErr)
	its.Equal(`"this is only a test"`, string(contents))
}
