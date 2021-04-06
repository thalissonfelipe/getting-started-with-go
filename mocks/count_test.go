package mocks

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	buffer := &bytes.Buffer{}

	Count(buffer)

	result := buffer.String()
	expected := "3"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
