package main

import (
	"fmt"
	"lab6/sort"
	"lab6/utils"
)

func main() {
	arr := utils.RandomArray(10, 123456)
	fmt.Printf("%v\n", arr)
	arr = sort.InsertionSort()
	fmt.Printf("%v\n", arr)
}
