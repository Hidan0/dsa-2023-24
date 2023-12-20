package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lab10/stack"
)

func main() {
	f, err := os.Open("in.txt")
	if err != nil {
		fmt.Println("Can't open file")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, ",")
		expected := tmp[1]
		res := evaluate(tmp[0])

		fmt.Printf("[evaluate(\"%s\")] expected: %s, Got: %d\n", tmp[0], expected, res)
	}
}

func evaluate(expr string) int {
	stack := stack.NewStack()
	tokens := strings.Split(expr, " ")
	for _, token := range tokens {
		n, err := strconv.Atoi(token)
		if err == nil {
			stack.Push(n)
			continue
		}
		switch token {
		case "+":
			n2 := stack.Pop()
			n1 := stack.Pop()
			stack.Push(n1 + n2)
		case "-":
			n2 := stack.Pop()
			n1 := stack.Pop()
			stack.Push(n1 - n2)
		case "*":
			n2 := stack.Pop()
			n1 := stack.Pop()
			stack.Push(n1 * n2)
		case "/":
			n2 := stack.Pop()
			n1 := stack.Pop()
			stack.Push(n1 / n2)
		default:
			fmt.Println("Invalid token", token)
			return -1
		}
	}

	return stack.Pop()
}
