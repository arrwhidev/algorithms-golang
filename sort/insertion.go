package sort

import "github.com/arrwhidev/algorithms-golang/util"

/**
Insertion Sort

- Walk through array, then compare backwards through 'sorted list'.
- Worst case: O(n^2)
- Best case: Ω(n)
*/

func Insertion(unsorted []int) []int {
	for i := 1; i < len(unsorted); i++ {
		for j := i; j > 0; j-- {
			if unsorted[j] < unsorted[j-1] {
				util.Swap(unsorted, j, j-1)
			}
		}
	}
	return unsorted
}
