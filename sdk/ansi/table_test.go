package ansi

import (
	"bytes"
	"fmt"
	"testing"

	"go.charczuk.com/sdk/assert"
)

func TestTable(t *testing.T) {
	its := assert.New(t)

	buffer := new(bytes.Buffer)

	err := Table(buffer, []string{"foo", "bar"}, [][]string{{"1", "2"}, {"3", "4"}})
	its.Nil(err)
	its.NotEmpty(buffer.String())
}

func TestTableEmptyErr(t *testing.T) {
	its := assert.New(t)

	buffer := new(bytes.Buffer)

	err := Table(buffer, nil, [][]string{{"1", "2"}, {"3", "4"}})
	its.NotNil(err)
}

func TestTableWriteErr(t *testing.T) {
	its := assert.New(t)

	failAfter := &failAfter{MaxBytes: 32}

	err := Table(failAfter, []string{"foo", "bar", "baz"}, [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	})
	its.NotNil(err)
	its.Equal("did fail", err.Error())
}

type failAfter struct {
	Written  []byte
	MaxBytes int
}

func (fa *failAfter) Write(contents []byte) (int, error) {
	fa.Written = append(fa.Written, contents...)
	if len(fa.Written) > fa.MaxBytes {
		return len(contents), fmt.Errorf("did fail")
	}
	return len(contents), nil
}

func TestTableForSlice(t *testing.T) {
	its := assert.New(t)

	objects := []struct {
		ID   int
		Name string
	}{
		{1, "Foo"},
		{2, "Bar"},
		{3, "Baz"},
	}

	output := new(bytes.Buffer)
	its.Nil(TableForSlice(output, objects))
	its.Equal(
		"┌──┬────┐\n│ID│Name│\n├──┼────┤\n│1 │Foo │\n│2 │Bar │\n│3 │Baz │\n└──┴────┘\n",
		output.String(),
	)
}

func TestTableForSliceUnicode(t *testing.T) {
	its := assert.New(t)

	objects := []struct {
		ID    string
		Count int
	}{
		{"モ foo", 1},
		{"ふ bar", 1},
		{"ス baz", 3},
	}

	output := new(bytes.Buffer)
	its.Nil(TableForSlice(output, objects))
	its.Equal(
		"┌──────┬─────┐\n│ID    │Count│\n├──────┼─────┤\n│モ foo│1    │\n│ふ bar│1    │\n│ス baz│3    │\n└──────┴─────┘\n",
		output.String(),
	)
}
