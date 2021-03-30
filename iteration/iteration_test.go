package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	verifyCorrectOutput := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("should return a string with repetitions", func(t *testing.T) {
		result := Iteration("a", 5)
		expected := "aaaaa"
		verifyCorrectOutput(t, result, expected)
	})

	t.Run("should return a string with N repetitions", func(t *testing.T) {
		result := Iteration("a", 3)
		expected := "aaa"
		verifyCorrectOutput(t, result, expected)
	})
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration("a", 5)
	}
}

func ExampleIteration() {
	result := Iteration("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
