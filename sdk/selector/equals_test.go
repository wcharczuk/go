package selector

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestEquals(t *testing.T) {
	its := assert.New(t)

	valid := Labels{
		"foo": "far",
		"moo": "bar",
	}
	its.True(Equals{Key: "foo", Value: "far"}.Matches(valid))
	its.False(Equals{Key: "zoo", Value: "buzz"}.Matches(valid))
	its.False(Equals{Key: "foo", Value: "bar"}.Matches(valid))

	its.Equal("foo == bar", Equals{Key: "foo", Value: "bar"}.String())
}
