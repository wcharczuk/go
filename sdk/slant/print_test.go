package slant

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestPrint(t *testing.T) {
	its := assert.New(t)

	output, err := PrintString("WARDEN")
	its.Nil(err)
	its.NotEmpty(output)
}
