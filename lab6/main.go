package main

import (
	"fmt"
	"lab6/sort"
	"lab6/utils"
)

func main() {
	arr := sort.IterativeSelectionSort(utils.RandomArray(10, 123456))
	fmt.Printf("%v\n", arr)
	arr = sort.RecursiveSelectionSort(utils.RandomArray(10, 123456))
	fmt.Printf("%v\n", arr)
}
