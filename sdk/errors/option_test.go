package errors

import (
	"fmt"
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestOptMessage(t *testing.T) {
	its := assert.New(t)

	ex := new(Exception)

	OptMessage("a message", " bar")(ex)
	its.Equal("a message bar", ex.Message)
}

func TestOptMessagef(t *testing.T) {
	its := assert.New(t)

	ex := new(Exception)

	OptMessagef("a message %s", "bar")(ex)
	its.Equal("a message bar", ex.Message)
}

func TestOptStackTrace(t *testing.T) {
	its := assert.New(t)

	ex := new(Exception)

	OptStackTrace(StackStrings([]string{"first", "second"}))(ex)
	its.NotNil(ex.StackTrace)
	its.Equal([]string{"first", "second"}, ex.StackTrace.Strings())
}

func TestOptInner(t *testing.T) {
	its := assert.New(t)

	ex := new(Exception)

	OptInner(fmt.Errorf("this is only a test"))(ex)
	its.NotNil(ex.Inner)
}

func TestOptInnerClass(t *testing.T) {
	its := assert.New(t)

	ex := new(Exception)

	OptInnerClass(fmt.Errorf("this is only a test"))(ex)
	its.NotNil(ex.Inner)
	its.Nil(ErrStackTrace(ex.Inner))
}
