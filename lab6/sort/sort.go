package sort

import (
	"fmt"
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
