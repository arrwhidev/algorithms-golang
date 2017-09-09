package util

import "math/rand"

func RandArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(300)
	}
	return arr
}

func Swap(a []int, i1 int, i2 int) {
	tmp := a[i1]
	a[i1] = a[i2]
	a[i2] = tmp
}
