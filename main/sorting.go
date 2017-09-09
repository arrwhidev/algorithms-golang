package main

import (
	"datastructures-and-algorithms/sorting"
	"datastructures-and-algorithms/util"
	"fmt"
)

func main() {
	// bubble()
	selection()
}

func bubble() {
	fmt.Println("Executing Bubble Sort...")
	unsorted := util.RandArray(10)
	sorted := sorting.Bubble(unsorted)
	fmt.Println(sorted)
}

func selection() {
	fmt.Println("Executing Selection Sort...")
	unsorted := util.RandArray(10)
	sorted := sorting.Selection(unsorted)
	fmt.Println(sorted)
}
