package util

import "math/rand"

func RandArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(300)
	}
	return arr
}
