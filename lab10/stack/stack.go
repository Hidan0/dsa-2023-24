package stack

type listNode struct {
	val  int
	next *listNode
	prev *listNode
}

func newNode(val int) *listNode {
	return &listNode{val, nil, nil}
}

type Stack struct {
	head *listNode
	tail *listNode
}

func NewStack() *Stack {
	return &Stack{nil, nil}
}

func (l *Stack) Push(val int) {
	n := newNode(val)
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.tail.next = n
	n.prev = l.tail
	l.tail = n
}

func (l *Stack) Pop() int {
	popped := l.tail

	last := popped.prev

	if last == nil {
		l.head = nil
		l.tail = nil
		return popped.val
	}

	last.next = nil
	l.tail = last

	return popped.val
}
