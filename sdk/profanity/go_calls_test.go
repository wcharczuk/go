package profanity

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestCallsPassing(t *testing.T) {
	its := assert.New(t)

	file := `package main

import "foo/bar"

func doFoo() {
	return
}

func main() {
	thing := make(map[string]string)
	fmt.Println(foo.Bar)
	println(bar.Foo)
	doFoo()
}
`
	rule := GoCalls([]GoCall{
		{
			Package: "fmt",
			Func:    "Printf",
		},
	})

	res := rule.Check("main.go", []byte(file))
	its.Nil(res.Err)
	its.True(res.OK)
}

func TestCallsPrintln(t *testing.T) {
	its := assert.New(t)

	file := `package main

import "foo/bar"

func doFoo() {
	return
}

func main() {
	thing := make(map[string]string)
	fmt.Println(foo.Bar)
	println(bar.Foo)
	doFoo()
}
`
	rule := GoCalls([]GoCall{
		{
			Package: "fmt",
			Func:    "Println",
		},
	})

	res := rule.Check("main.go", []byte(file))
	its.Nil(res.Err)
	its.False(res.OK)
	its.Equal("main.go", res.File)
	its.Equal(11, res.Line)
}

func TestCallsEmptyPackage(t *testing.T) {
	its := assert.New(t)

	file := `package main

import "foo/bar"

func doFoo() {
	return
}

func main() {
	thing := make(map[string]string)
	fmt.Println(foo.Bar)
	println(bar.Foo)
	doFoo()
}
`

	rule := GoCalls([]GoCall{
		{
			Func: "println",
		},
	})

	res := rule.Check("main.go", []byte(file))
	its.Nil(res.Err)
	its.False(res.OK)
	its.Equal("main.go", res.File)
	its.Equal(12, res.Line)
}
