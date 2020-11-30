package uuid

import (
	"fmt"
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestV4(t *testing.T) {
	m := make(map[string]bool)
	for x := 1; x < 32; x++ {
		uuid := V4()
		s := uuid.ToFullString()
		if m[s] {
			t.Errorf("NewRandom returned duplicated UUID %s\n", s)
		}
		m[s] = true
		if v := uuid.Version(); v != 4 {
			t.Errorf("Random UUID of version %v\n", v)
		}
	}
}

func makeTestUUIDv4(versionNumber byte, variant byte) UUID {
	return []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, versionNumber, 0x0, variant, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
}

func TestIsUUIDv4(t *testing.T) {
	it := assert.New(t)

	valid := makeTestUUIDv4(0x40, 0x80)
	versionInvalid := makeTestUUIDv4(0xF0, 0x80)
	variantInvalid := makeTestUUIDv4(0x40, 0xF0)
	lengthInvalid := UUID([]byte{})

	it.True(valid.IsV4())
	it.False(variantInvalid.IsV4())
	it.False(versionInvalid.IsV4())
	it.False(lengthInvalid.IsV4())
}

func TestParseUUIDv4Valid(t *testing.T) {
	it := assert.New(t)

	validShort := V4().ToShortString()
	validParsedShort, err := Parse(validShort)
	it.Nil(err)
	it.True(validParsedShort.IsV4())
	it.Equal(validShort, validParsedShort.ToShortString())

	validFull := V4().ToFullString()
	validParsedFull, err := Parse(validFull)
	it.Nil(err)
	it.True(validParsedFull.IsV4())
	it.Equal(validFull, validParsedFull.ToFullString())

	validBracedShort := fmt.Sprintf("{%s}", validShort)
	validParsedBracedShort, err := Parse(validBracedShort)
	it.Nil(err)
	it.True(validParsedBracedShort.IsV4())
	it.Equal(validShort, validParsedBracedShort.ToShortString())

	validBracedFull := fmt.Sprintf("{%s}", validFull)
	validParsedBracedFull, err := Parse(validBracedFull)
	it.Nil(err)
	it.True(validParsedBracedFull.IsV4())
	it.Equal(validFull, validParsedBracedFull.ToFullString())
}

func TestParseUUIDv4Invalid(t *testing.T) {
	it := assert.New(t)

	_, err := Parse("fcae3946f75d+3258678bb5e6795a6d3")
	it.NotNil(err, "should handle invalid characters")

	_, err = Parse("4f2e28b7b8f94b9eba1d90c4452")
	it.NotNil(err, "should handle invalid length uuids")
}
