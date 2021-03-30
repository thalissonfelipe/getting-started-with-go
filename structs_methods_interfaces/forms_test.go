package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	result := Perimeter(10.0, 10.0)
	expected := 40.0

	if result != expected {
		t.Errorf("result %.2f expected %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {
	result := Area(12.0, 6.0)
	expected := 72.0

	if result != expected {
		t.Errorf("result %.2f expected %.2f", result, expected)
	}
}