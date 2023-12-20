package main

import (
	"fmt"
	"lab6/sort"
	"lab6/utils"
)

func main() {
	arr := sort.IterativeSelectionSort(utils.RandomArray(10, 123456))
	fmt.Printf("IterativeSelectionSort: %v\n", arr)
	arr = sort.RecursiveSelectionSort(utils.RandomArray(10, 123456))
	fmt.Printf("RecursiveSelectionSort: %v\n", arr)
	arr = utils.RandomArray(10, 123456)
	sort.MergeSort(&arr)
	fmt.Printf("MergeSort: %v\n", arr)
}
