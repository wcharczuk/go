package selector

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestHasKey(t *testing.T) {
	its := assert.New(t)

	valid := Labels{
		"foo": "far",
	}
	its.True(HasKey("foo").Matches(valid))
	its.False(HasKey("zoo").Matches(valid))
	its.Equal("foo", HasKey("foo").String())
}
