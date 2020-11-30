package selector

import (
	"fmt"
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestMustParse(t *testing.T) {
	its := assert.New(t)

	its.Equal("x == a", MustParse("x==a").String())

	var err error
	func() {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%+v", r)
			}
		}()
		MustParse("x!!")
	}()
	its.NotNil(err)
}

func TestParseInvalid(t *testing.T) {
	its := assert.New(t)

	testBadStrings := []string{
		"x=a||y=b",
		"x==a==b",
		"!x=a",
		"x<a",
		"x>1",
		"x>1,z<5",
	}
	var err error
	for _, str := range testBadStrings {
		_, err = Parse(str)
		its.NotNil(err, str)
	}
}

func TestParseSemiValid(t *testing.T) {
	its := assert.New(t)

	testGoodStrings := []string{
		"",
		"x=a,y=b,z=c",
		"x!=a,y=b",
		"x=",
		"x= ",
		"x=,z= ",
		"x= ,z= ",
		"!x",
	}

	var err error
	for _, str := range testGoodStrings {
		_, err = Parse(str)
		its.Nil(err, str)
	}
}

func TestParseEquals(t *testing.T) {
	its := assert.New(t)

	valid := Labels{
		"foo": "bar",
		"moo": "lar",
	}
	invalid := Labels{
		"zoo": "mar",
		"moo": "lar",
	}

	selector, err := Parse("foo == bar")
	its.Nil(err)
	its.True(selector.Matches(valid))
	its.False(selector.Matches(invalid))
}

func TestParseNotEquals(t *testing.T) {
	its := assert.New(t)

	valid := Labels{
		"foo": "far",
		"moo": "lar",
	}
	invalidPresent := Labels{
		"foo": "bar",
		"moo": "lar",
	}
	invalidMissing := Labels{
		"zoo": "mar",
		"moo": "lar",
	}

	selector, err := Parse("foo != bar")
	its.Nil(err)
	its.True(selector.Matches(valid))
	its.True(selector.Matches(invalidMissing))
	its.False(selector.Matches(invalidPresent))
}

func TestParseIn(t *testing.T) {
	its := assert.New(t)

	valid := Labels{
		"foo": "far",
		"moo": "lar",
	}
	valid2 := Labels{
		"foo": "bar",
		"moo": "lar",
	}
	invalid := Labels{
		"foo": "mar",
		"moo": "lar",
	}
	invalidMissing := Labels{
		"zoo": "mar",
		"moo": "lar",
	}

	selector, err := Parse("foo in (bar,far)")
	its.Nil(err)
	its.True(selector.Matches(valid), selector.String())
	its.True(selector.Matches(valid2))
	its.True(selector.Matches(invalidMissing))
	its.False(selector.Matches(invalid), selector.String())
}

func TestParseGroup(t *testing.T) {
	its := assert.New(t)

	valid := Labels{
		"zoo":   "mar",
		"moo":   "lar",
		"thing": "map",
	}
	invalid := Labels{
		"zoo":   "mar",
		"moo":   "something",
		"thing": "map",
	}
	invalid2 := Labels{
		"zoo":    "mar",
		"moo":    "lar",
		"!thing": "map",
	}
	selector, err := Parse("zoo=mar, moo=lar, thing")
	its.Nil(err)
	its.True(selector.Matches(valid))
	its.False(selector.Matches(invalid))
	its.False(selector.Matches(invalid2))

	complicated, err := Parse("zoo in (mar,lar,dar),moo,!thingy")
	its.Nil(err)
	its.NotNil(complicated)
	its.True(complicated.Matches(valid))
}

func TestParseGroupComplicated(t *testing.T) {
	its := assert.New(t)
	valid := Labels{
		"zoo":   "mar",
		"moo":   "lar",
		"thing": "map",
	}
	complicated, err := Parse("zoo in (mar,lar,dar),moo,thing == map,!thingy")
	its.Nil(err)
	its.NotNil(complicated)
	its.True(complicated.Matches(valid))
}

func TestParseDocsExample(t *testing.T) {
	its := assert.New(t)
	selector, err := Parse("x in (foo,,baz),y,z notin ()")
	its.Nil(err)
	its.NotNil(selector)
}

func TestParseEqualsOperators(t *testing.T) {
	its := assert.New(t)

	selector, err := Parse("notin=in")
	its.Nil(err)

	typed, isTyped := selector.(Equals)
	its.True(isTyped)
	its.Equal("notin", typed.Key)
	its.Equal("in", typed.Value)
}

func TestParseValidate(t *testing.T) {
	its := assert.New(t)

	_, err := Parse("zoo=bar")
	its.Nil(err)

	_, err = Parse("_zoo=bar")
	its.NotNil(err)

	_, err = Parse("_zoo=_bar")
	its.NotNil(err)

	_, err = Parse("zoo=bar,foo=_mar")
	its.NotNil(err)
}

func TestParseRegressionCSVSymbols(t *testing.T) {
	its := assert.New(t)

	sel, err := Parse("foo in (bar-bar, baz.baz, buzz_buzz), moo=boo")
	its.Nil(err, "regression is values can have '-' in them")
	its.NotEmpty(sel.String())
}

func TestParseRegressionIn(t *testing.T) {
	its := assert.New(t)

	_, err := Parse("foo in bar, buzz)")
	its.NotNil(err)
}

func TestParseMultiByte(t *testing.T) {
	its := assert.New(t)

	selector, err := Parse("함=수,목=록") // number=number, number=rock
	its.Nil(err)
	its.NotNil(selector)

	typed, isTyped := selector.(And)
	its.True(isTyped)
	its.Len(2, typed)
}

func TestParseOptions(t *testing.T) {
	its := assert.New(t)

	selQuery := "bar=foo@bar"
	labels := Labels{
		"foo": "bar",
		"bar": "foo@bar",
	}

	sel, err := Parse(selQuery)
	its.NotNil(err)
	its.Nil(sel)

	sel, err = Parse(selQuery, SkipValidation)
	its.Nil(err)
	its.NotNil(sel)

	its.True(sel.Matches(labels))
}

func BenchmarkParse(b *testing.B) {
	valid := Labels{
		"zoo":   "mar",
		"moo":   "lar",
		"thing": "map",
	}

	for i := 0; i < b.N; i++ {
		selector, err := Parse("zoo in (mar,lar,dar),moo,!thingy")
		if err != nil {
			b.Fail()
		}
		if !selector.Matches(valid) {
			b.Fail()
		}
	}
}
