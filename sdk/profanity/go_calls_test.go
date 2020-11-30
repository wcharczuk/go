package profanity

import (
	"testing"

	"go.charczuk.com/sdk/assert"
)

func Test_GoCalls_passing(t *testing.T) {
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

func Test_GoCalls_println(t *testing.T) {
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

func Test_GoCalls_EmptyPackage(t *testing.T) {
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
