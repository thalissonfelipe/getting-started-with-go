package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "this is just a test"}
	result := Search(dict, "test")
	expected := "this is just a test"

	compareStrings(t, result, expected)
}

func compareStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s', given '%s'", result, expected, "test")
	}
}
