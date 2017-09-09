package search

import "testing"

func TestLinear_shouldReturnMinus1_whenElementNotFound(t *testing.T) {
	array := []int{5, 7, 3, 9, 6}
	result := Linear(array, 2)
	if result != -1 {
		t.Error("Expected -1 when element not found in array")
	}
}

func TestLinear_shouldReturnIndex_whenElementFound(t *testing.T) {
	array := []int{5, 7, 3, 9, 6}
	result := Linear(array, 9)
	if result != 3 {
		t.Error("Expected index 3 but got ", result)
	}
}
