package selector

import (
	"strings"
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestParserIsWhitespace(t *testing.T) {
	its := assert.New(t)

	l := &Parser{}
	its.True(l.isWhitespace(' '))
	its.True(l.isWhitespace('\n'))
	its.True(l.isWhitespace('\r'))
	its.True(l.isWhitespace('\t'))

	its.False(l.isWhitespace('a'))
	its.False(l.isWhitespace('z'))
	its.False(l.isWhitespace('A'))
	its.False(l.isWhitespace('Z'))
	its.False(l.isWhitespace('1'))
	its.False(l.isWhitespace('-'))
}

func TestParserIsAlpha(t *testing.T) {
	its := assert.New(t)

	l := &Parser{}
	its.True(l.isAlpha('a'))
	its.True(l.isAlpha('z'))
	its.True(l.isAlpha('A'))
	its.True(l.isAlpha('Z'))
	its.True(l.isAlpha('1'))

	its.False(l.isAlpha('-'))
	its.False(l.isAlpha(' '))
	its.False(l.isAlpha('\n'))
	its.False(l.isAlpha('\r'))
	its.False(l.isAlpha('\t'))
}

func TestParserSkipWhitespace(t *testing.T) {
	its := assert.New(t)

	l := &Parser{s: "foo    != bar    ", pos: 3}
	its.Equal(" ", string(l.current()))
	l.skipWhiteSpace()
	its.Equal(7, l.pos)
	its.Equal("!", string(l.current()))
	l.pos = 14
	its.Equal(" ", string(l.current()))
	l.skipWhiteSpace()
	its.Equal(len(l.s), l.pos)
}

func TestParserReadWord(t *testing.T) {
	its := assert.New(t)

	l := &Parser{s: "foo != bar"}
	its.Equal("foo", l.readWord())
	its.Equal(" ", string(l.current()))

	l = &Parser{s: "foo,"}
	its.Equal("foo", l.readWord())
	its.Equal(",", string(l.current()))

	l = &Parser{s: "foo"}
	its.Equal("foo", l.readWord())
	its.True(l.done())
}

func TestParserReadOp(t *testing.T) {
	its := assert.New(t)

	l := &Parser{s: "!= bar"}
	op, err := l.readOp()
	its.Nil(err)
	its.Equal("!=", op)
	its.Equal(" ", string(l.current()))

	l = &Parser{s: "!=bar"}
	op, err = l.readOp()
	its.Nil(err)
	its.Equal("!=", op)
	its.Equal("b", string(l.current()))

	l = &Parser{s: "!=bar"}
	op, err = l.readOp()
	its.Nil(err)
	its.Equal("!=", op)
	its.Equal("b", string(l.current()))

	l = &Parser{s: "!="}
	op, err = l.readOp()
	its.Nil(err)
	its.Equal("!=", op)
	its.True(l.done())

	l = &Parser{s: "= bar"}
	op, err = l.readOp()
	its.Nil(err)
	its.Equal("=", op)
	its.Equal(" ", string(l.current()))

	l = &Parser{s: "=bar"}
	op, err = l.readOp()
	its.Nil(err)
	its.Equal("=", op)
	its.Equal("b", string(l.current()))

	l = &Parser{s: "== bar"}
	op, err = l.readOp()
	its.Nil(err)
	its.Equal("==", op)
	its.Equal(" ", string(l.current()))

	l = &Parser{s: "==bar"}
	op, err = l.readOp()
	its.Nil(err)
	its.Equal("==", op)
	its.Equal("b", string(l.current()))

	l = &Parser{s: "in (foo)"}
	op, err = l.readOp()
	its.Nil(err)
	its.Equal("in", op)
	its.Equal(" ", string(l.current()))

	l = &Parser{s: "in(foo)"}
	op, err = l.readOp()
	its.Nil(err)
	its.Equal("in", op)
	its.Equal("(", string(l.current()))

	l = &Parser{s: "notin (foo)"}
	op, err = l.readOp()
	its.Nil(err)
	its.Equal("notin", op)
	its.Equal(" ", string(l.current()))

	l = &Parser{s: "notin(foo)"}
	op, err = l.readOp()
	its.Nil(err)
	its.Equal("notin", op)
	its.Equal("(", string(l.current()))
}

func TestParserReadCSV(t *testing.T) {
	its := assert.New(t)

	l := &Parser{s: "(bar, baz, biz)"}
	words, err := l.readCSV()
	its.Nil(err)
	its.Len(3, words, strings.Join(words, ","))
	its.Equal("bar", words[0])
	its.Equal("baz", words[1])
	its.Equal("biz", words[2])
	its.True(l.done())

	l = &Parser{s: "(bar,baz,biz)"}
	words, err = l.readCSV()
	its.Nil(err)
	its.Len(3, words, strings.Join(words, ","))
	its.Equal("bar", words[0])
	its.Equal("baz", words[1])
	its.Equal("biz", words[2])
	its.True(l.done())

	l = &Parser{s: "(bar, buzz, baz"}
	words, err = l.readCSV()
	its.NotNil(err)
	its.Empty(words)

	l = &Parser{s: "()"}
	words, err = l.readCSV()
	its.Nil(err)
	its.Empty(words)
	its.True(l.done())

	l = &Parser{s: "(), thing=after"}
	words, err = l.readCSV()
	its.Nil(err)
	its.Empty(words)
	its.Equal(",", string(l.current()))

	l = &Parser{s: "(foo, bar), buzz=light"}
	words, err = l.readCSV()
	its.Nil(err)
	its.Len(2, words)
	its.Equal("foo", words[0])
	its.Equal("bar", words[1])
	its.Equal(",", string(l.current()))

	l = &Parser{s: "(test, space are bad)"}
	words, err = l.readCSV()
	its.NotNil(err)
	its.Empty(words)
}

func TestParserHasKey(t *testing.T) {
	its := assert.New(t)
	l := &Parser{s: "foo"}
	valid, err := l.Parse()
	its.Nil(err)
	its.NotNil(valid)
	typed, isTyped := valid.(HasKey)
	its.True(isTyped)
	its.Equal("foo", string(typed))
}

func TestParserNotHasKey(t *testing.T) {
	its := assert.New(t)
	l := &Parser{s: "!foo"}
	valid, err := l.Parse()
	its.Nil(err)
	its.NotNil(valid)
	typed, isTyped := valid.(NotHasKey)
	its.True(isTyped)
	its.Equal("foo", string(typed))
}

func TestParserEquals(t *testing.T) {
	its := assert.New(t)

	l := &Parser{s: "foo = bar"}
	valid, err := l.Parse()
	its.Nil(err)
	its.NotNil(valid)
	typed, isTyped := valid.(Equals)
	its.True(isTyped)
	its.Equal("foo", typed.Key)
	its.Equal("bar", typed.Value)

	l = &Parser{s: "foo=bar"}
	valid, err = l.Parse()
	its.Nil(err)
	its.NotNil(valid)
	typed, isTyped = valid.(Equals)
	its.True(isTyped)
	its.Equal("foo", typed.Key)
	its.Equal("bar", typed.Value)
}

func TestParserDoubleEquals(t *testing.T) {
	its := assert.New(t)
	l := &Parser{s: "foo == bar"}
	valid, err := l.Parse()
	its.Nil(err)
	its.NotNil(valid)
	typed, isTyped := valid.(Equals)
	its.True(isTyped)
	its.Equal("foo", typed.Key)
	its.Equal("bar", typed.Value)
}

func TestParserNotEquals(t *testing.T) {
	its := assert.New(t)
	l := &Parser{s: "foo != bar"}
	valid, err := l.Parse()
	its.Nil(err)
	its.NotNil(valid)
	typed, isTyped := valid.(NotEquals)
	its.True(isTyped)
	its.Equal("foo", typed.Key)
	its.Equal("bar", typed.Value)
}

func TestParserIn(t *testing.T) {
	its := assert.New(t)
	l := &Parser{s: "foo in (bar, baz)"}
	valid, err := l.Parse()
	its.Nil(err)
	its.NotNil(valid)
	typed, isTyped := valid.(In)
	its.True(isTyped)
	its.Equal("foo", typed.Key)
	its.Len(2, typed.Values)
	its.Equal("bar", typed.Values[0])
	its.Equal("baz", typed.Values[1])
}

func TestParserLex(t *testing.T) {
	its := assert.New(t)
	l := &Parser{s: ""}
	_, err := l.Parse()
	its.Nil(err)
}
