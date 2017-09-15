package sort

import (
	"testing"

	"github.com/arrwhidev/testutils-golang"
)

func TestBubble_shouldSortArray(t *testing.T) {
	input := []int{5, 7, 3, 9, 6, 50, 2, 8, 11, 1}
	expected := []int{1, 2, 3, 5, 6, 7, 8, 9, 11, 50}
	utils.AssertArraysEqual(t, expected, Bubble(input))
}

func TestSelection_shouldSortArray(t *testing.T) {
	input := []int{5, 7, 3, 9, 6, 50, 2, 8, 11, 1}
	expected := []int{1, 2, 3, 5, 6, 7, 8, 9, 11, 50}
	utils.AssertArraysEqual(t, expected, Selection(input))
}

func TestInsertion_shouldSortArray(t *testing.T) {
	input := []int{5, 7, 3, 9, 6, 50, 2, 8, 11, 1}
	expected := []int{1, 2, 3, 5, 6, 7, 8, 9, 11, 50}
	utils.AssertArraysEqual(t, expected, Insertion(input))
}

func TestMerge_shouldSortArray(t *testing.T) {
	input := []int{5, 7, 3, 9, 6, 50, 2, 8, 11, 1}
	expected := []int{1, 2, 3, 5, 6, 7, 8, 9, 11, 50}
	utils.AssertArraysEqual(t, expected, Merge(input))
}

func TestQuick_shouldSortArray(t *testing.T) {
	input := []int{5, 7, 3, 9, 6, 50, 2, 8, 11, 1}
	expected := []int{1, 2, 3, 5, 6, 7, 8, 9, 11, 50}
	utils.AssertArraysEqual(t, expected, Quick(input))
}

func TestMedianOfThree_shouldMoveMedianValueToEnd(t *testing.T) {
	slices := [][]int{
		[]int{3, 7, 9},
		[]int{7, 3, 9},
		[]int{3, 9, 7},
		[]int{9, 3, 7},
		[]int{7, 9, 3},
		[]int{9, 7, 3},
	}

	for _, slice := range slices {
		MedianOfThree(slice, 0, 1, 2)
		if slice[0] != 3 || slice[1] != 9 || slice[2] != 7 {
			t.Error("Failed!")
		}
	}
}
