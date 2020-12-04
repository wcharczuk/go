package errors

import (
	"fmt"
	"strings"
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestMulti(t *testing.T) {
	its := assert.New(t)

	ex0 := New(New("hi0"))
	ex1 := New(fmt.Errorf("hi1"))
	ex2 := New("hi2")

	m := Append(ex0, ex1, ex2)

	its.True(strings.HasPrefix(m.Error(), `3 errors occurred:`), m.Error()) //todo, make this test more strict
	its.Len(3, m.(Multi).WrappedErrors())
	its.NotNil(m.(Multi).Unwrap())
}
