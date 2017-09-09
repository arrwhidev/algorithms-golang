package sort

func Selection(unsorted []int) []int {
	size := len(unsorted)
	smallestIndex := 0
	tmp := 0
	for sortedOffset := 0; sortedOffset < size; sortedOffset++ {
		smallestIndex = sortedOffset
		for i := sortedOffset + 1; i < size; i++ {
			if unsorted[i] < unsorted[smallestIndex] {
				smallestIndex = i
			}
		}
		tmp = unsorted[sortedOffset]
		unsorted[sortedOffset] = unsorted[smallestIndex]
		unsorted[smallestIndex] = tmp
	}
	return unsorted
}
