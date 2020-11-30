package selector

import (
	"fmt"
	"strings"
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestCheckKey(t *testing.T) {
	its := assert.New(t)

	its.Nil(CheckKey("foo"))
	its.Nil(CheckKey("bar/foo"))
	its.Nil(CheckKey("bar.io/foo"))
	its.NotNil(CheckKey("_foo"))
	its.NotNil(CheckKey("-foo"))
	its.NotNil(CheckKey("foo-"))
	its.NotNil(CheckKey("foo_"))
	its.NotNil(CheckKey("bar/foo/baz"))

	its.NotNil(CheckKey(""), "should error on empty keys")

	its.NotNil(CheckKey("/foo"), "should error on empty dns prefixes")
	superLongDNSPrefixed := fmt.Sprintf("%s/%s", strings.Repeat("a", MaxDNSPrefixLen), strings.Repeat("a", MaxKeyLen))
	its.Nil(CheckKey(superLongDNSPrefixed), len(superLongDNSPrefixed))
	superLongDNSPrefixed = fmt.Sprintf("%s/%s", strings.Repeat("a", MaxDNSPrefixLen+1), strings.Repeat("a", MaxKeyLen))
	its.NotNil(CheckKey(superLongDNSPrefixed), len(superLongDNSPrefixed))
	superLongDNSPrefixed = fmt.Sprintf("%s/%s", strings.Repeat("a", MaxDNSPrefixLen+1), strings.Repeat("a", MaxKeyLen+1))
	its.NotNil(CheckKey(superLongDNSPrefixed), len(superLongDNSPrefixed))
	superLongDNSPrefixed = fmt.Sprintf("%s/%s", strings.Repeat("a", MaxDNSPrefixLen), strings.Repeat("a", MaxKeyLen+1))
	its.NotNil(CheckKey(superLongDNSPrefixed), len(superLongDNSPrefixed))
}

func TestCheckKeyK8S(t *testing.T) {
	its := assert.New(t)

	values := []string{
		// the "good" cases
		"simple",
		"now-with-dashes",
		"1-starts-with-num",
		"1234",
		"simple/simple",
		"now-with-dashes/simple",
		"now-with-dashes/now-with-dashes",
		"now.with.dots/simple",
		"now-with.dashes-and.dots/simple",
		"1-num.2-num/3-num",
		"1234/5678",
		"1.2.3.4/5678",
		"Uppercase_Is_OK_123",
		"example.com/Uppercase_Is_OK_123",
		"requests.storage-foo",
		strings.Repeat("a", 63),
		strings.Repeat("a", 253) + "/" + strings.Repeat("b", 63),
	}
	badValues := []string{
		// the "bad" cases
		"nospecialchars%^=@",
		"cantendwithadash-",
		"-cantstartwithadash-",
		"only/one/slash",
		"Example.com/abc",
		"example_com/abc",
		"example.com/",
		"/simple",
		strings.Repeat("a", 64),
		strings.Repeat("a", 254) + "/abc",
	}
	for _, val := range values {
		its.Nil(CheckKey(val))
	}
	for _, val := range badValues {
		its.NotNil(CheckKey(val))
	}
}

func TestCheckValue(t *testing.T) {
	its := assert.New(t)

	its.Nil(CheckValue(""), "should not error on empty values")
	its.Nil(CheckValue("foo"))
	its.Nil(CheckValue("bar_baz"))
	its.NotNil(CheckValue("_bar_baz"))
	its.NotNil(CheckValue("bar_baz_"))
	its.NotNil(CheckValue("_bar_baz_"))
}

func TestIsAlpha(t *testing.T) {
	its := assert.New(t)

	its.True(isAlpha('A'))
	its.True(isAlpha('a'))
	its.True(isAlpha('Z'))
	its.True(isAlpha('z'))
	its.True(isAlpha('0'))
	its.True(isAlpha('9'))
	its.True(isAlpha('함'))
	its.True(isAlpha('é'))
	its.False(isAlpha('-'))
	its.False(isAlpha('/'))
	its.False(isAlpha('~'))
}

func TestCheckLabels(t *testing.T) {
	its := assert.New(t)

	goodLabels := Labels{"foo": "bar", "foo.com/bar": "baz"}
	its.Nil(CheckLabels(goodLabels))
	badLabels := Labels{"foo": "bar", "_foo.com/bar": "baz"}
	its.NotNil(CheckLabels(badLabels))
}
