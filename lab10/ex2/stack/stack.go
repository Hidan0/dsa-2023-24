package stack

// Another stack implementation

type node[T any] struct {
	value T
	next  *node[T]
}

type Stack[T any] struct {
	top *node[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (s *Stack[T]) Push(value T) {
	n := &node[T]{value, s.top}
	s.top = n
}

func (s *Stack[T]) Pop() T {
	n := s.top
	s.top = n.next
	return n.value
}

func (s *Stack[T]) Clear() {
	s.top = nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack[T]) Size() int {
	size := 0

	for n := s.top; n != nil; n = n.next {
		size++
	}

	return size
}
