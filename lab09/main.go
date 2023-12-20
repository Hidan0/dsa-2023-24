package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type circularNode struct {
	val  int
	next *circularNode
	prev *circularNode
}

type CircularLinkedList struct {
	head *circularNode
}

func (lst *CircularLinkedList) insert(val int) {
	newNode := new(circularNode)
	newNode.val = val

	if lst.head == nil {
		lst.head = newNode
		lst.head.next = lst.head
		lst.head.prev = lst.head
		return
	}

	last := lst.head.prev
	last.next = newNode

	newNode.prev = last
	newNode.next = lst.head
	lst.head.prev = newNode
}

func (lst *CircularLinkedList) print() {
	fmt.Print("[ ")
	if lst.head == nil {
		fmt.Println("]")
		return
	}

	curr := lst.head
	stopAt := lst.head.prev
	ok := true

	for ok {
		if curr == stopAt {
			ok = false
		}
		fmt.Printf("%d ", curr.val)
		curr = curr.next
	}

	fmt.Println("]")
}

// Assumes that the list contains only one node with value `0`
func (lst *CircularLinkedList) printFromZero() {
	fmt.Print("[ ")
	if lst.head == nil {
		fmt.Println("]")
		return
	}

	curr := lst.head
	for curr.val != 0 {
		curr = curr.next
	}

	stopAt := curr.prev
	ok := true
	for ok {
		if curr == stopAt {
			ok = false
		}
		fmt.Printf("%d ", curr.val)
		curr = curr.next
	}

	fmt.Println("]")
}

func main() {
	f, err := os.Open("in.txt")
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lst := new(CircularLinkedList)

	for scanner.Scan() {
		line := scanner.Text()

		switch line {
		case "p":
			lst.printFromZero()
		default:
			n, _ := strconv.Atoi(line)
			lst.insert(n)
		}
	}
	fmt.Println("Done")
	fmt.Print("Normal print: ")
	lst.print()
}
