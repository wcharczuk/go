package errors

import (
	"fmt"
	"testing"

	"go.charczuk.com/sdk/assert"
)

type classProvider struct {
	error
	ErrClass error
}

func (cp classProvider) Class() error {
	return cp.ErrClass
}

func TestErrClass(t *testing.T) {
	its := assert.New(t)

	its.Nil(ErrClass(nil))
	var unsetErr error
	its.Nil(ErrClass(unsetErr))

	its.Nil(ErrClass("foo"))

	err := New("this is a test")
	its.Equal("this is a test", ErrClass(err).Error())

	cp := classProvider{
		error:    fmt.Errorf("this is a provider test"),
		ErrClass: fmt.Errorf("the error class"),
	}
	its.Equal("the error class", ErrClass(cp).Error())
	its.Equal("this is a test", ErrClass(fmt.Errorf("this is a test")).Error())
}

func TestErrMessage(t *testing.T) {
	its := assert.New(t)

	its.Empty(ErrMessage(nil))
	its.Empty(ErrMessage(fmt.Errorf("foo bar baz")))
	its.Equal("this is a message", ErrMessage(New("error class", OptMessage("this is a message"))))
}

type stackProvider struct {
	error
	Stack StackTrace
}

func (sp stackProvider) StackTrace() StackTrace {
	return sp.Stack
}

func TestErrStackTrace(t *testing.T) {
	its := assert.New(t)

	err := New("this is a test")
	its.NotNil(ErrStackTrace(err))

	sp := stackProvider{
		error: fmt.Errorf("this is a provider test"),
		Stack: StackStrings([]string{"first", "second"}),
	}
	its.Equal([]string{"first", "second"}, ErrStackTrace(sp).Strings())

	its.Nil(ErrStackTrace(fmt.Errorf("this is also a test")))
}
