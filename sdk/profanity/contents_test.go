package profanity

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func Test_Contents_Contains_NoMatch(t *testing.T) {
	its := assert.New(t)

	contents := `
foo bar baz
bar foo baz
baz bar foo
`
	rule := Contents{
		Contains: ContainsFilter{
			Filter: Filter{
				Include: []string{"fuzzy wuzzy"},
			},
		},
	}

	res := rule.Check("test.go", []byte(contents))
	its.True(res.OK)
}

func Test_Contents_Contains_Include(t *testing.T) {
	its := assert.New(t)

	contents := `
foo bar baz
bar foo baz
baz bar foo
`
	rule := Contents{
		Contains: ContainsFilter{
			Filter: Filter{
				Include: []string{"foo baz"},
			},
		},
	}

	res := rule.Check("test.go", []byte(contents))
	its.False(res.OK)
	its.Equal("test.go", res.File)
	its.Equal(3, res.Line)
}

func Test_Contents_Contains_IncludeExclude(t *testing.T) {
	its := assert.New(t)

	contents := `
foo bar baz
bar foo moo
baz bar foo
buzz foo baz
`
	rule := Contents{
		Contains: ContainsFilter{
			Filter: Filter{
				Include: []string{"foo baz"},
				Exclude: []string{"buzz"},
			},
		},
	}

	res := rule.Check("test.go", []byte(contents))
	its.True(res.OK)
}

func Test_Contents_Glob(t *testing.T) {
	its := assert.New(t)

	contents := `
foo bar baz
bar foo baz
baz bar foo
`
	rule := Contents{
		Glob: GlobFilter{
			Filter: Filter{
				Include: []string{"bar*"},
			},
		},
	}

	res := rule.Check("test.go", []byte(contents))
	its.False(res.OK)
	its.Equal("test.go", res.File)
	its.Equal(3, res.Line)
}
