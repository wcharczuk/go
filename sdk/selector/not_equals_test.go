package selector

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestNotEquals(t *testing.T) {
	its := assert.New(t)

	valid := Labels{
		"foo": "far",
		"moo": "bar",
	}
	its.False(NotEquals{Key: "foo", Value: "far"}.Matches(valid))
	its.True(NotEquals{Key: "zoo", Value: "buzz"}.Matches(valid))
	its.True(NotEquals{Key: "foo", Value: "bar"}.Matches(valid))
	its.Equal("foo != bar", NotEquals{Key: "foo", Value: "bar"}.String())
}
