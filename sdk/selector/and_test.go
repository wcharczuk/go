package selector

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestAnd(t *testing.T) {
	its := assert.New(t)

	valid := Labels{
		"foo": "far",
		"moo": "lar",
	}
	invalid := Labels{
		"foo": "far",
		"moo": "bar",
	}

	selector := And([]Selector{Equals{Key: "foo", Value: "far"}, Equals{Key: "moo", Value: "lar"}})
	its.True(selector.Matches(valid))
	its.False(selector.Matches(invalid))

	its.Equal("foo == far, moo == lar", selector.String())
}
