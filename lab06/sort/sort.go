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

func merge(arr, x *[]int, start, middle, end int) {
	k := 0
	i1 := start
	i2 := middle

	for i1 < middle && i2 < end {
		if (*arr)[i1] <= (*arr)[i2] {
			(*x)[k] = (*arr)[i1]
			i1 += 1
		} else {
			(*x)[k] = (*arr)[i2]
			i2 += 1
		}
		k += 1
	}

	if i1 < middle {
		for j := i1; j < middle; j++ {
			(*x)[k] = (*arr)[j]
			k += 1
		}
	} else {
		for j := i2; j < end; j++ {
			(*x)[k] = (*arr)[j]
			k += 1
		}
	}
	for k := 0; k < end-start; k++ {
		(*arr)[start+k] = (*x)[k]
	}
}

func mergeSort(arr, x *[]int, start, end int) {
	if end-start > 1 {
		middle := (start + end) / 2
		mergeSort(arr, x, start, middle)
		mergeSort(arr, x, middle, end)
		merge(arr, x, start, middle, end)
	}
}

func MergeSort(arr *[]int) {
	x := make([]int, len(*arr))
	mergeSort(arr, &x, 0, len(*arr))
}
