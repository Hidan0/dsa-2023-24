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

func numberOfMovesHanoi(n, from, tmp, to int) uint64 {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	numberOfMovesHanoi(n-1, from, to, tmp)
	return 2*numberOfMovesHanoi(n-1, tmp, from, to) + 1
}

func main() {
	hanoi(3, 0, 1, 2)

	fmt.Printf("To move 3 disks, we need %d moves\n", numberOfMovesHanoi(3, 0, 1, 2))
	fmt.Printf("To move 10 disks, we need %d moves\n", numberOfMovesHanoi(10, 0, 1, 2))
	// fmt.Printf("To move 64 disks, we need %d moves\n", numberOfMovesHanoi(64, 0, 1, 2)) // end of the world
}
