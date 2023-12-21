package main

import (
	"bufio"
	"fmt"
	"lab10/stack"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("in.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		s := stack.New[string]()
		pos := 0
		for _, token := range tokens {
			token = strings.Trim(token, "<>")
			if token[0] == '/' {
				tag := s.Pop()
				if tag != token[1:] {
					fmt.Printf("[%s] Error at position %d. Expected <%s> found <%s>\n", line, pos, tag, token)
					s.Clear()
					break
				}
				continue
			}
			s.Push(token)
			pos += 1
		}

		if s.Size() > 0 {
			fmt.Printf("[%s] The following tags are not closed: ", line)
			for !s.IsEmpty() {
				fmt.Printf("<%s> ", s.Pop())
			}
			fmt.Println()
		}
	}
}
