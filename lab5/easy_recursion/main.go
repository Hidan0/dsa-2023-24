package main

import "fmt"

// Es 1
func multiply(x, y int) int {
	if y <= 0 {
		return 0
	} else {
		return x + multiply(x, y-1)
	}
}

// Es 2
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largest(numbers []int) int {
	n := len(numbers)
	if n == 1 {
		return numbers[0]
	}
	return max(numbers[n-1], largest(numbers[:n-1]))
}

// Es 3
// Dei sassi sferici sono ammucchiati a formare una piramide, con un
// sasso in cima, posto al centro di un quadrato formato da 4 sassi (2
// per lato), posti a loro volta sopra un quadrato formato da 9 sassi (3
// per lato), e così via.
func pebbles(height int) int {
	if height <= 0 {
		return 0
	}

	return height*height + pebbles(height-1)
}

func main() {
	// == Es 1 ==
	fmt.Printf("5 * 5 = %d\n", multiply(5, 5))
	fmt.Printf("5 * 3 = %d\n", multiply(5, 3))
	// =========

	// == Es 2 ==
	nums := []int{1, 2, 5, 7, -2, 10, 9, 21, 3, 8}
	fmt.Printf("The largest number is %d\n", largest(nums))
	// =========
	// Trace of execution
	// largest([1, 2, 5, 7, -2, 10, 9, 21, 3, 8])
	// max(8, largest([1, 2, 5, 7, -2, 10, 9, 21, 3]))
	// max(8, max(3, largest([1, 2, 5, 7, -2, 10, 9, 21])))
	// max(8, max(3, max(21, largest([1, 2, 5, 7, -2, 10, 9]))))
	// max(8, max(3, max(21, max(9, largest([1, 2, 5, 7, -2, 10])))))
	// max(8, max(3, max(21, max(9, max(10, largest([1, 2, 5, 7, -2])))))))
	// max(8, max(3, max(21, max(9, max(10, max(-2, largest([1, 2, 5, 7]))))))))
	// max(8, max(3, max(21, max(9, max(10, max(-2, max(7, largest([1, 2, 5])))))))))
	// max(8, max(3, max(21, max(9, max(10, max(-2, max(7, max(5, largest([1, 2]))))))))))
	// max(8, max(3, max(21, max(9, max(10, max(-2, max(7, max(5, max(2, largest([1]))))))))))) -> largest([1]) is the first call that terminates
	// max(8, max(3, max(21, max(9, max(10, max(-2, max(7, max(5, max(2, 1)))))))))) -> max(2, 1) the arguments with which max is called the first time
	// max(8, max(3, max(21, max(9, max(10, max(-2, max(7, max(5, 2)))))))))
	// max(8, max(3, max(21, max(9, max(10, max(-2, max(7, 5))))))))
	// max(8, max(3, max(21, max(9, max(10, max(-2, 7)))))))
	// max(8, max(3, max(21, max(9, max(10, 7))))))
	// max(8, max(3, max(21, max(9, 10)))))
	// max(8, max(3, max(21, 10))))
	// max(8, max(3, 21)))
	// max(8, 21)) -> the arguments with which max is called the last time
	// 21

	// == Es 3 ==
	fmt.Printf("pebbles(1) = %d\n", pebbles(1))
	fmt.Printf("pebbles(2) = %d\n", pebbles(2))
	fmt.Printf("pebbles(3) = %d\n", pebbles(3))
	fmt.Printf("pebbles(4) = %d\n", pebbles(4))
	// =========
}
