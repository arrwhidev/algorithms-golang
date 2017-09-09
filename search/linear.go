package search

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
