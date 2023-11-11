package main

import "fmt"

func hanoi(n, from, tmp, to int) {
	if n <= 0 {
		return
	}

	hanoi(n-1, from, to, tmp)
	fmt.Printf("%d -> %d\n", from, to)
	hanoi(n-1, tmp, from, to)
}

func main() {
	hanoi(3, 0, 1, 2)
}
