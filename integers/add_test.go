package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("result '%d', expected '%d", result, expected)
	}
}

func ExampleAdd() {
	result := Add(1, 5)
	fmt.Println(result)
	// Output: 6
}
