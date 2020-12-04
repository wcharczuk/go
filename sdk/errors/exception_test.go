package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestNewOfString(t *testing.T) {
	its := assert.New(t)
	ex := As(New("this is a test"))
	its.Equal("this is a test", fmt.Sprintf("%v", ex))
	its.NotNil(ex.StackTrace)
	its.Nil(ex.Inner)
}

func TestNewOfError(t *testing.T) {
	its := assert.New(t)

	err := errors.New("This is an error")
	wrappedErr := New(err)
	its.NotNil(wrappedErr)
	typedWrapped := As(wrappedErr)
	its.NotNil(typedWrapped)
	its.Equal("This is an error", fmt.Sprintf("%v", typedWrapped))
}

func TestNewOfException(t *testing.T) {
	its := assert.New(t)

	ex := New(Class("This is an exception"))
	wrappedEx := New(ex)
	its.NotNil(wrappedEx)
	typedWrappedEx := As(wrappedEx)
	its.Equal("This is an exception", fmt.Sprintf("%v", typedWrappedEx))
	its.Equal(ex, typedWrappedEx)
}

func TestNewOfNil(t *testing.T) {
	its := assert.New(t)

	shouldBeNil := New(nil)
	its.Nil(shouldBeNil)
	its.Equal(nil, shouldBeNil)
	its.True(nil == shouldBeNil)
}

func TestNewOfTypedNil(t *testing.T) {
	its := assert.New(t)

	var nilError error
	its.Nil(nilError)
	its.Equal(nil, nilError)

	shouldBeNil := New(nilError)
	its.Nil(shouldBeNil)
	its.True(shouldBeNil == nil)
}

func TestNewOfReturnedNil(t *testing.T) {
	its := assert.New(t)

	returnsNil := func() error {
		return nil
	}

	shouldBeNil := New(returnsNil())
	its.Nil(shouldBeNil)
	its.True(shouldBeNil == nil)

	returnsTypedNil := func() error {
		return New(nil)
	}

	shouldAlsoBeNil := returnsTypedNil()
	its.Nil(shouldAlsoBeNil)
	its.True(shouldAlsoBeNil == nil)
}

func TestError(t *testing.T) {
	its := assert.New(t)

	ex := New(Class("this is a test"))
	message := ex.Error()
	its.NotEmpty(message)
}

func TestErrorOptions(t *testing.T) {
	its := assert.New(t)

	ex := New(Class("this is a test"), OptMessage("foo"))
	message := ex.Error()
	its.NotEmpty(message)

	typed := As(ex)
	its.NotNil(typed)
	its.Equal("foo", typed.Message)
}

func TestCallers(t *testing.T) {
	its := assert.New(t)

	callStack := func() StackTrace { return Callers(DefaultStartDepth) }()

	its.NotNil(callStack)
	callstackStr := callStack.String()
	its.True(strings.Contains(callstackStr, "TestCallers"), callstackStr)
}

func TestExceptionFormatters(t *testing.T) {
	its := assert.New(t)

	// test the "%v" formatter with just the exception class.
	class := &Exception{Class: Class("this is a test")}
	its.Equal("this is a test", fmt.Sprintf("%v", class))

	classAndMessage := &Exception{Class: Class("foo"), Message: "bar"}
	its.Equal("foo; bar", fmt.Sprintf("%v", classAndMessage))
}

func TestMarshalJSON(t *testing.T) {

	type ReadableStackTrace struct {
		Class   string   `json:"Class"`
		Message string   `json:"Message"`
		Inner   error    `json:"Inner"`
		Stack   []string `json:"StackTrace"`
	}

	its := assert.New(t)
	message := "new test error"
	ex := As(New(message))
	its.NotNil(ex)
	stackTrace := ex.StackTrace
	typed, isTyped := stackTrace.(StackPointers)
	its.True(isTyped)
	its.NotNil(typed)
	stackDepth := len(typed)

	jsonErr, err := json.Marshal(ex)
	its.Nil(err)
	its.NotNil(jsonErr)

	ex2 := &ReadableStackTrace{}
	err = json.Unmarshal(jsonErr, ex2)
	its.Nil(err)
	its.Len(stackDepth, ex2.Stack)
	its.Equal(message, ex2.Class)

	ex = As(New(fmt.Errorf(message)))
	its.NotNil(ex)
	stackTrace = ex.StackTrace
	typed, isTyped = stackTrace.(StackPointers)
	its.True(isTyped)
	its.NotNil(typed)
	stackDepth = len(typed)

	jsonErr, err = json.Marshal(ex)
	its.Nil(err)
	its.NotNil(jsonErr)

	ex2 = &ReadableStackTrace{}
	err = json.Unmarshal(jsonErr, ex2)
	its.Nil(err)
	its.Len(stackDepth, ex2.Stack)
	its.Equal(message, ex2.Class)
}

func TestJSON(t *testing.T) {
	its := assert.New(t)

	ex := New("this is a test",
		OptMessage("test message"),
		OptInner(New("inner exception", OptMessagef("inner test message"))),
	)

	contents, err := json.Marshal(ex)
	its.Nil(err)

	var verify Exception
	err = json.Unmarshal(contents, &verify)
	its.Nil(err)

	its.Equal(ErrClass(ex), ErrClass(verify))
	its.Equal(ErrMessage(ex), ErrMessage(verify))
	its.NotNil(verify.Inner)
	its.Equal(ErrClass(ErrInner(ex)), ErrClass(ErrInner(verify)))
	its.Equal(ErrMessage(ErrInner(ex)), ErrMessage(ErrInner(verify)))
}

func TestNest(t *testing.T) {
	its := assert.New(t)

	ex1 := As(New("this is an error"))
	ex2 := As(New("this is another error"))
	err := As(Nest(ex1, ex2))

	its.NotNil(err)
	its.NotNil(err.Inner)
	its.NotEmpty(err.Error())

	its.True(Is(ex1, Class("this is an error")))
	its.True(Is(ex1.Inner, Class("this is another error")))
}

func TestNestNil(t *testing.T) {
	its := assert.New(t)

	var ex1 error
	var ex2 error
	var ex3 error

	err := Nest(ex1, ex2, ex3)
	its.Nil(err)
	its.Equal(nil, err)
	its.True(nil == err)
}

func TestExceptionFormat(t *testing.T) {
	its := assert.New(t)

	e := &Exception{Class: fmt.Errorf("this is only a test")}
	output := fmt.Sprintf("%v", e)
	its.Equal("this is only a test", output)

	output = fmt.Sprintf("%+v", e)
	its.Equal("this is only a test", output)

	e = &Exception{
		Class: fmt.Errorf("this is only a test"),
		StackTrace: StackStrings([]string{
			"foo",
			"bar",
		}),
	}

	output = fmt.Sprintf("%+v", e)
	its.Equal("this is only a test\nfoo\nbar", output)
}

func TestExceptionPrintsInner(t *testing.T) {
	its := assert.New(t)

	ex := New("outer", OptInner(New("middle", OptInner(New("terminal")))))

	output := fmt.Sprintf("%v", ex)

	its.Contains(output, "outer")
	its.Contains(output, "middle")
	its.Contains(output, "terminal")

	output = fmt.Sprintf("%+v", ex)

	its.Contains(output, "outer")
	its.Contains(output, "middle")
	its.Contains(output, "terminal")
}
