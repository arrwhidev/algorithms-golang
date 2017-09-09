package main

import (
	"fmt"

	"github.com/arrwhidev/algorithms-golang/sorting"
	"github.com/arrwhidev/algorithms-golang/util"
)

func main() {
	bubble()
	selection()
}

func bubble() {
	fmt.Println("Executing Bubble Sort...")
	unsorted := util.RandArray(10)
	sorted := sorting.Bubble(unsorted)
	fmt.Println("--> ", sorted)
}

func selection() {
	fmt.Println("Executing Selection Sort...")
	unsorted := util.RandArray(10)
	sorted := sorting.Selection(unsorted)
	fmt.Println("--> ", sorted)
}
