package utils

import (
	"fmt"
	"log"
	"math/rand"
)

func ReadUntilZero() []int {
	var num int = -1
	var out []int

	for num != 0 {
		_, err := fmt.Scan(&num)

		if err != nil {
			log.Fatal(err)
			break
		}

		out = append(out, num)
	}

	return out
}

func RandomArray(n int, seed int64) []int {
	out := make([]int, n)
	r := rand.New(rand.NewSource(seed))

	for i := 0; i < n; i++ {
		out[i] = r.Int() % 144 // more readable
	}

	return out
}
