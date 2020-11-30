package selector

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestIn(t *testing.T) {
	its := assert.New(t)

	valid := Labels{
		"foo": "far",
		"moo": "lar",
	}
	valid2 := Labels{
		"foo": "bar",
		"moo": "lar",
	}
	missing := Labels{
		"loo": "mar",
		"moo": "lar",
	}
	invalid := Labels{
		"foo": "mar",
		"moo": "lar",
	}

	selector := In{Key: "foo", Values: []string{"bar", "far"}}
	its.True(selector.Matches(valid))
	its.True(selector.Matches(valid2))
	its.True(selector.Matches(missing))
	its.False(selector.Matches(invalid))

	its.Equal("foo in (bar, far)", selector.String())
}
