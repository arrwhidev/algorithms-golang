package main

import (
	"fmt"

	"github.com/arrwhidev/algorithms-golang/sort"
	"github.com/arrwhidev/algorithms-golang/util"
)

func main() {
	unsorted := util.RandArray(10)
	bubble(unsorted)
	selection(unsorted)
}

func bubble(unsorted []int) {
	fmt.Println("Executing Bubble Sort:")
	sorted := sort.Bubble(unsorted)
	fmt.Println("--> ", sorted)
}

func selection(unsorted []int) {
	fmt.Println("Executing Selection Sort:")
	sorted := sort.Selection(unsorted)
	fmt.Println("--> ", sorted)
}
