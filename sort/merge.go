package sort

/**
Merge Sort

- Recursively split array into smaller lists until they contain a single element each.
- Recursively merge smaller sorted lists back up.
- Worst case: O(n log n)
- Best case: Ω(n log n)
*/

func Merge(unsorted []int) []int {
	if len(unsorted) < 2 {
		return unsorted
	}

	mid := len(unsorted) / 2
	return mergeTwo(Merge(unsorted[:mid]), Merge(unsorted[mid:]))
}

func mergeTwo(left, right []int) []int {
	mergedLength := len(left) + len(right)
	merged := make([]int, mergedLength, mergedLength)
	leftIndex, rightIndex := 0, 0

	for i := 0; i < mergedLength; i++ {
		if leftIndex > len(left)-1 && rightIndex <= len(right)-1 {
			merged[i] = right[rightIndex]
			rightIndex++
		} else if rightIndex > len(right)-1 && leftIndex <= len(left)-1 {
			merged[i] = left[leftIndex]
			leftIndex++
		} else if left[leftIndex] > right[rightIndex] {
			merged[i] = right[rightIndex]
			rightIndex++
		} else {
			merged[i] = left[leftIndex]
			leftIndex++
		}
	}

	return merged
}
