package sort

import (
	"github.com/arrwhidev/algorithms-golang/util"
)

/**
Quick Sort

- This implementation uses the `median of three` method to choose a good pivot.
- This ensures that sorted input lists are handled in good time instead of O(n^2)
- Worst case: O(n^2)
- Best case: Ω(n log n)
*/

func Quick(unsorted []int) []int {
	if len(unsorted) < 2 {
		return unsorted
	}

	left := 0
	right := len(unsorted) - 1
	mid := left + (right-left)/2 // Avoids overflow when doing (left + right)
	MedianOfThree(unsorted, 0, mid, len(unsorted)-1)
	pivotIndex := len(unsorted) - 1

	// Move lower items to left of array.
	for i, el := range unsorted {
		if el < unsorted[pivotIndex] {
			util.Swap(unsorted, i, left)
			left++
		}
	}
	util.Swap(unsorted, left, pivotIndex)

	// Recurse the elements lower and higher than the pivot
	Quick(unsorted[:left])
	Quick(unsorted[left+1:])
	return unsorted
}

// Moves median value to end of array.
func MedianOfThree(a []int, lowIndex int, midIndex int, highIndex int) {
	if a[lowIndex] > a[midIndex] {
		util.Swap(a, lowIndex, midIndex)
	}

	if a[lowIndex] > a[highIndex] {
		util.Swap(a, lowIndex, highIndex)
	}

	if a[midIndex] < a[highIndex] {
		util.Swap(a, midIndex, highIndex)
	}
}
