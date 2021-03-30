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
	testsArea := []struct {
		form     Form
		expected float64
	}{
		{form: Rectangle{12, 6}, expected: 72.0},
		{form: Circle{10}, expected: 314.1592653589793},
		{form: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range testsArea {
		result := tt.form.Area()
		if result != tt.expected {
			t.Errorf("result %.2f, expected %.2f", result, tt.expected)
		}
	}
}
