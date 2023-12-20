package lib

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

type LinkedList struct {
	head *node
}

func newNode(value int) *node {
	newNode := new(node)
	newNode.value = value
	newNode.next = nil
	return newNode
}

func (lst *LinkedList) PushNode(value int) {
	if lst.head == nil {
		lst.head = newNode(value)
		return
	}

	curr := lst.head
	next := curr.next
	for next != nil {
		curr = next
		next = curr.next
	}

	curr.next = newNode(value)
}

func (lst *LinkedList) Print() string {
	sb := strings.Builder{}

	curr := lst.head
	for ; curr != nil; curr = curr.next {
		sb.WriteString(fmt.Sprintf("%d -> ", curr.value))
	}
	sb.WriteString("nil")

	return sb.String()
}

func (lst *LinkedList) Size() int {
	if lst == nil {
		return 0
	}

	out := 0
	curr := lst.head
	for curr != nil {
		out += 1
		curr = curr.next
	}
	return out
}

func (lst *LinkedList) Drop() {
	lst.head = nil
}

func (lst *LinkedList) PrintReversed() {
	l := lst.head

	fmt.Print("[ ")
	lst.printReversed(l)
	fmt.Println("]")
}

func (lst *LinkedList) printReversed(l *node) {
	if l == nil {
		return
	}
	lst.printReversed(l.next)
	fmt.Printf("%d ", l.value)
}

func (lst *LinkedList) Search(value int) bool {
	curr := lst.head
	for curr != nil && curr.value != value {
		curr = curr.next
	}

	return curr != nil
}

func (lst *LinkedList) Remove(value int) bool {
	curr := lst.head
	var prev *node = nil

	for curr != nil && curr.value != value {
		prev = curr
		curr = curr.next
	}

	if curr == nil {
		return false
	}

	if prev == nil {
		lst.head = curr.next
	} else {
		prev.next = curr.next
	}
	return true
}
