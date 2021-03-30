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
		name     string
		form     Form
		expected float64
	}{
		{name: "Rectangle", form: Rectangle{12, 6}, expected: 72.0},
		{name: "Circle", form: Circle{10}, expected: 314.1592653589793},
		{name: "Triangle", form: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range testsArea {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.form.Area()
			if result != tt.expected {
				t.Errorf("%#v: result %.2f, expected %.2f", tt.form, result, tt.expected)
			}
		})
	}
}
