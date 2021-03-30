package iteration

import "testing"

func TestIteration(t *testing.T) {
	result := Iteration("a")
	expected := "aaaaa"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
