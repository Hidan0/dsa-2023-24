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

func (s *Stack) Push(val int) {
	n := newNode(val)
	if s.head == nil {
		s.head = n
		s.tail = n
		return
	}

	s.tail.next = n
	n.prev = s.tail
	s.tail = n
}

func (s *Stack) Pop() int {
	popped := s.tail

	last := popped.prev

	if last == nil {
		s.head = nil
		s.tail = nil
		return popped.val
	}

	last.next = nil
	s.tail = last

	return popped.val
}

type SliceStack struct {
	data []string
}

func NewSStack() *SliceStack {
	return &SliceStack{}
}

func (s *SliceStack) Push(val string) {
	s.data = append(s.data, val)
}

func (s *SliceStack) Pop() string {
	out := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return out
}
