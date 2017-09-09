package sort

import "github.com/arrwhidev/algorithms-golang/util"

/**
Bubble Sort

- Gradually move higher elements to right, lower elements to left.
- Worst case: O(n^2)
- Best case: Ω(n)
*/

func Bubble(unsorted []int) []int {
	size := len(unsorted)
	didSwap := true
	for didSwap {
		didSwap = false
		for i := 1; i < size; i++ {
			if unsorted[i-1] > unsorted[i] {
				util.Swap(unsorted, i-1, i)
				didSwap = true
			}
		}
	}

	return unsorted
}
