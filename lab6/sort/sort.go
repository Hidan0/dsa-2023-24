package sort

import (
	"fmt"
	"lab6/utils"
	"log"
)

func InsertionSort() []int {
	var num int

	var arr []int

	for {
		_, err := fmt.Scan(&num)

		if err != nil {
			log.Fatal(err)
			break
		}

		if num == 0 {
			break
		}

		inserted := false

		for i := 0; i < len(arr); i++ {
			if num < arr[i] {
				arr = append(arr, 0)
				for k := len(arr) - 2; k >= i; k-- {
					arr[k+1] = arr[k]
				}
				arr[i] = num
				inserted = true
				break
			}
		}

		if !inserted {
			arr = append(arr, num)
			continue
		}
	}
	return arr
}

func findMax(arr *[]int, from, to int) int {
	max := from

	for i := from; i < to; i++ {
		if (*arr)[i] > (*arr)[max] {
			max = i
		}
	}

	return max
}

func IterativeSelectionSort(arr []int) []int {
	for k := len(arr) - 1; k >= 1; k-- {
		max := findMax(&arr, 0, k+1)
		utils.Swap(&arr, k, max)
	}
	return arr
}

func RecursiveSelectionSort(arr []int) []int {
	return selectionSort(arr, 0, len(arr))
}

func selectionSort(arr []int, from, to int) []int {
	if from >= to {
		return arr
	}
	max := findMax(&arr, from, to)
	utils.Swap(&arr, max, to-1)
	return selectionSort(arr, from, to-1)
}
