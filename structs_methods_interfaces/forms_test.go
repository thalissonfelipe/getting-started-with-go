package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expected := 40.0

	if result != expected {
		t.Errorf("result %.2f expected %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	result := Area(rectangle)
	expected := 72.0

	if result != expected {
		t.Errorf("result %.2f expected %.2f", result, expected)
	}
}
