package selector

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestNotIn(t *testing.T) {
	its := assert.New(t)

	valid := Labels{
		"foo": "mar",
		"moo": "lar",
	}
	invalid := Labels{
		"foo": "far",
		"moo": "lar",
	}
	missing := Labels{
		"loo": "mar",
		"moo": "lar",
	}

	selector := NotIn{Key: "foo", Values: []string{"bar", "far"}}
	its.True(selector.Matches(valid))
	its.True(selector.Matches(missing))
	its.False(selector.Matches(invalid))
	its.Equal("foo notin (bar, far)", selector.String())
}
