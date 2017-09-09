package search

/**
Binary Search

- Input array must be sorted.
- Worst case: O(log n)
- Best case: Ω(1)
*/

func Binary(array []int, target int) int {
	result := -1
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := left + (right-left)/2 // Avoids overflow when doing (left + right)
		if target < array[mid] {
			right = mid - 1
		} else if target > array[mid] {
			left = mid + 1
		} else {
			result = mid
			break
		}
	}

	return result
}
