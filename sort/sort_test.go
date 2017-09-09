package sort

import (
	"reflect"
	"testing"
)

func TestBubble_shouldSortArray(t *testing.T) {
	array := []int{5, 7, 3, 9, 6}
	expected := []int{3, 5, 6, 7, 9}
	result := Bubble(array)
	if !reflect.DeepEqual(expected, result) {
		t.Error("Expected array to be sorted")
	}
}

func TestSelection_shouldSortArray(t *testing.T) {
	array := []int{5, 7, 3, 9, 6}
	expected := []int{3, 5, 6, 7, 9}
	result := Selection(array)
	if !reflect.DeepEqual(expected, result) {
		t.Error("Expected array to be sorted")
	}
}
