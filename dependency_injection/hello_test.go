package dependency_injection

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	buffer := bytes.Buffer{}
	Hello(&buffer, "Felipe")

	result := buffer.String()
	expected := "Hello, Felipe"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
