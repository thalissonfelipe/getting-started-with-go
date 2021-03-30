package integers

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("result '%d', expected '%d", result, expected)
	}
}
