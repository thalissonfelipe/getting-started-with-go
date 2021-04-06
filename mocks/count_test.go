package mocks

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeperSpy := &SleeperSpy{}

	Count(buffer, sleeperSpy)

	result := buffer.String()
	expected := `3
2
1
Go!`

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}

	if sleeperSpy.Calls != 4 {
		t.Errorf("there were not enough calls from the sleeper, result %d, expected 4", sleeperSpy.Calls)
	}
}
