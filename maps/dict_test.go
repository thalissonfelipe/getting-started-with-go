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
	t.Run("should add a new word", func(t *testing.T) {
		word := "test"
		value := "this is just a test"
		dict := Dict{}

		err := dict.Add(word, value)

		compareError(t, err, nil)
		compareDefinition(t, dict, word, value)
	})

	t.Run("should return an error when the word already exists", func(t *testing.T) {
		word := "test"
		value := "this is just a test"
		dict := Dict{word: value}

		err := dict.Add(word, value)

		compareError(t, err, ErrWordAlreadyExists)
		compareDefinition(t, dict, word, value)
	})
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

func compareDefinition(t *testing.T, dict Dict, word, value string) {
	t.Helper()

	result, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find the added word:", err)
	}

	if value != result {
		t.Errorf("result '%s',  expected '%s'", result, value)
	}
}
