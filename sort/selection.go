package sort

import "github.com/arrwhidev/algorithms-golang/util"

/**
Selection Sort

- Finds smallest element each iteration.
- Inserts it to the end of the sorted section of the array.
- Worst case: O(n^2)
- Best case: Ω(n^2)
*/

func Selection(unsorted []int) []int {
	size := len(unsorted)
	for sortedOffset := 0; sortedOffset < size-1; sortedOffset++ {
		indexOfSmallest := sortedOffset
		for i := sortedOffset + 1; i < size; i++ {
			if unsorted[i] < unsorted[indexOfSmallest] {
				indexOfSmallest = i
			}
		}
		if sortedOffset != indexOfSmallest {
			util.Swap(unsorted, sortedOffset, indexOfSmallest)
		}
	}
	return unsorted
}
