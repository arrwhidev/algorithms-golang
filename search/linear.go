package search

/**
Linear Search

- Go through entire array until item is found.
- Worst case: O(n)
- Best case: Ω(1)
*/

func Linear(array []int, target int) int {
	indexOfTarget := -1
	for i, el := range array {
		if el == target {
			indexOfTarget = i
			break
		}
	}
	return indexOfTarget
}
