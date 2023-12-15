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

func (lst *LinkedList) Search(value int) bool {
	curr := lst.head
	for curr != nil && curr.value != value {
		curr = curr.next
	}

	return curr != nil
}
