package array_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	verifyCorrectOutput := func(t *testing.T, result int, expected int, numbers []int) {
		t.Helper()
		if result != expected {
			t.Errorf("result %d, expected %d, given %v", result, expected, numbers)
		}
	}

	t.Run("should return a sum of an array with any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sum(numbers)
		expected := 6
		verifyCorrectOutput(t, result, expected, numbers)
	})
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result %v, expected %v", result, expected)
	}
}

func TestSumAllRest(t *testing.T) {
	t.Run("should sum slices", func(t *testing.T) {
		result := SumAllRest([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v, expected %v", result, expected)
		}
	})

	t.Run("should sum empty slices safely", func(t *testing.T) {
		result := SumAllRest([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v, expected %v", result, expected)
		}
	})
}
