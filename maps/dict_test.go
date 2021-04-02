package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dict{"test": "this is just a test"}

	t.Run("search a known word", func(t *testing.T) {
		result, _ := dict.Search("test")
		expected := "this is just a test"

		compareStrings(t, result, expected)
	})

	t.Run("search an unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")

		compareError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dict{}
	dict.Add("test", "this is just a test")

	result, err := dict.Search("test")
	expected := "this is just a test"

	compareError(t, err, nil)
	compareStrings(t, result, expected)
}

func compareStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s', given '%s'", result, expected, "test")
	}
}

func compareError(t *testing.T, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
