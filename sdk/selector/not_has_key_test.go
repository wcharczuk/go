package selector

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestNotHasKey(t *testing.T) {
	its := assert.New(t)

	valid := Labels{
		"foo": "far",
	}
	its.False(NotHasKey("foo").Matches(valid))
	its.True(NotHasKey("zoo").Matches(valid))
	its.Equal("!foo", NotHasKey("foo").String())
}
