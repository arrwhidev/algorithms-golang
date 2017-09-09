package sorting

func Bubble(unsorted []int) []int {
	size := len(unsorted)
	swapCount := -1
	tmp := -1
	for swapCount != 0 {
		swapCount = 0
		for i := range unsorted {
			if (i != size-1) && (unsorted[i] > unsorted[i+1]) {
				tmp = unsorted[i]
				unsorted[i] = unsorted[i+1]
				unsorted[i+1] = tmp
				swapCount++
			}
		}
	}

	return unsorted
}
