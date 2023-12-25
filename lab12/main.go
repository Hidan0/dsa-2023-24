package main

import (
	"fmt"
	"strings"
)

type bTreeNode[T any] struct {
	left  *bTreeNode[T]
	right *bTreeNode[T]
	val   T
}

func newBTNode[T any](val T) *bTreeNode[T] {
	return &bTreeNode[T]{left: nil, right: nil, val: val}
}

type BTree[T any] struct {
	root *bTreeNode[T]
}

func newBTree[T any]() *BTree[T] {
	return &BTree[T]{root: nil}
}

func (btn *bTreeNode[T]) _inorder(sb *strings.Builder) {
	if btn == nil {
		return
	}
	btn.left._inorder(sb)
	sb.WriteString(fmt.Sprintf("%v ", btn.val))
	btn.right._inorder(sb)
}

func (bt *BTree[T]) inorder() string {
	sb := strings.Builder{}
	bt.root._inorder(&sb)
	sb.WriteString("\n")
	return sb.String()
}

func (btn *bTreeNode[T]) _preorder(sb *strings.Builder) {
	if btn == nil {
		return
	}
	sb.WriteString(fmt.Sprintf("%v ", btn.val))
	btn.left._preorder(sb)
	btn.right._preorder(sb)
}

func (bt *BTree[T]) preorder() string {
	sb := strings.Builder{}
	bt.root._preorder(&sb)
	sb.WriteString("\n")
	return sb.String()
}

func (btn *bTreeNode[T]) _postorder(sb *strings.Builder) {
	if btn == nil {
		return
	}
	btn.left._postorder(sb)
	btn.right._postorder(sb)
	sb.WriteString(fmt.Sprintf("%v ", btn.val))
}

func (bt *BTree[T]) postorder() string {
	sb := strings.Builder{}
	bt.root._postorder(&sb)
	sb.WriteString("\n")
	return sb.String()
}

func (btn *bTreeNode[T]) _printSummary(spaces int, sb *strings.Builder) {
	// modified pre-order traversal
	if btn == nil {
		return
	}

	sb.WriteString(strings.Repeat(" ", spaces))
	sb.WriteString(fmt.Sprintf("*%v\n", btn.val))

	if btn.left == nil && btn.right != nil {
		sb.WriteString(strings.Repeat(" ", spaces+2) + "*\n")
	}
	btn.left._printSummary(spaces+2, sb)
	if btn.right == nil && btn.left != nil {
		sb.WriteString(strings.Repeat(" ", spaces+2) + "*\n")
	}
	btn.right._printSummary(spaces+2, sb)
}

func (bt *BTree[T]) printSummary() string {
	sb := strings.Builder{}
	bt.root._printSummary(0, &sb)
	return sb.String()
}

func main() {
	bt := newBTree[int]()
	bt.root = newBTNode(78)
	bt.root.left = newBTNode(54)
	bt.root.left.right = newBTNode(90)
	bt.root.left.right.left = newBTNode(19)
	bt.root.left.right.right = newBTNode(95)
	bt.root.right = newBTNode(21)
	bt.root.right.left = newBTNode(16)
	bt.root.right.right = newBTNode(19)
	bt.root.right.left.left = newBTNode(5)
	bt.root.right.right.left = newBTNode(56)
	bt.root.right.right.right = newBTNode(43)

	fmt.Println("[INORDER]: ", bt.inorder())
	fmt.Println("[PREORDER]: ", bt.preorder())
	fmt.Println("[POSTORDER]: ", bt.postorder())
	fmt.Print("[SUMMARY]:\n", bt.printSummary())
}
