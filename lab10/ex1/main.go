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

		postFix := fromInfixToPostfix(tmp[0])
		res := evaluate(postFix)

		fmt.Printf("[\"%s\"] => [evaluate(\"%s\")] expected: %s, Got: %d\n", tmp[0], postFix, expected, res)
	}
}

func evaluate(expr string) int {
	stack := stack.NewStack()
	tokens := strings.Split(strings.TrimSpace(expr), " ")
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

func fromInfixToPostfix(expr string) string {
	stack := stack.NewSStack()
	tokens := strings.Split(strings.TrimSpace(expr), " ")
	sb := strings.Builder{}

	for _, token := range tokens {
		_, err := strconv.Atoi(token)
		if err == nil {
			sb.WriteString(token)
			sb.WriteString(" ")
			continue
		}

		switch token {
		case "(":
			continue
		case ")":
			sb.WriteString(stack.Pop())
			sb.WriteString(" ")
		case "+", "-", "*", "/":
			stack.Push(token)
		default:
			fmt.Println("Invalid token", token)
			return ""
		}
	}

	return sb.String()
}
