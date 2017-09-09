package sort

import (
	"reflect"
	"testing"
)

func TestBubble_shouldSortArray(t *testing.T) {
	input := []int{5, 7, 3, 9, 6}
	expected := []int{3, 5, 6, 7, 9}
	assertEqual(t, expected, Bubble(input))
}

func TestSelection_shouldSortArray(t *testing.T) {
	input := []int{5, 7, 3, 9, 6}
	expected := []int{3, 5, 6, 7, 9}
	assertEqual(t, expected, Selection(input))
}

func TestInsertion_shouldSortArray(t *testing.T) {
	input := []int{5, 7, 3, 9, 6}
	expected := []int{3, 5, 6, 7, 9}
	assertEqual(t, expected, Insertion(input))
}

func assertEqual(t *testing.T, expected []int, result []int) {
	if !reflect.DeepEqual(expected, result) {
		t.Error("Expected array to be sorted. Result:", result)
	}
}
