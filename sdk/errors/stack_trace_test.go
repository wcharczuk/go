package errors

import (
	"fmt"
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestGetStackTrace(t *testing.T) {
	its := assert.New(t)

	its.NotEmpty(GetStackTrace())
}

func TestStackStrings(t *testing.T) {
	its := assert.New(t)

	stack := []string{
		"foo",
		"bar",
		"baz",
	}

	stackStrings := StackStrings(stack)

	its.Equal("\nfoo\nbar\nbaz", fmt.Sprintf("%+v", stackStrings))
	its.Equal("[]string{\"foo\", \"bar\", \"baz\"}", fmt.Sprintf("%#v", stackStrings))
	its.Equal("\nfoo\nbar\nbaz", fmt.Sprintf("%v", stackStrings))
	its.Equal([]string{"foo", "bar", "baz"}, stackStrings)
}

func TestExceptionWithStackStrings(t *testing.T) {
	its := assert.New(t)

	stack := []string{
		"foo",
		"bar",
		"baz",
	}

	ex := As(New("foo", OptStackTrace(StackStrings(stack))))

	values := ex.Decompose()
	its.NotEmpty(values["StackTrace"])
	its.NotNil(ex.StackTrace)
}
