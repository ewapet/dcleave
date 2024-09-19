package internal

type SliceStack[T any] struct {
	nextPushPosition int
	elements         []T
}

func NewSliceStack[T any](initialCapacity int) *SliceStack[T] {
	return &SliceStack[T]{
		nextPushPosition: 0,
		elements:         make([]T, 0, initialCapacity),
	}
}

func (s *SliceStack[T]) Copy() *SliceStack[T] {
	copiedStack := NewSliceStack[T](len(s.elements))
	copy(copiedStack.elements, s.elements)
	return copiedStack
}

func (s *SliceStack[T]) Push(value T) {
	if s.nextPushPosition >= len(s.elements) {
		s.elements = append(s.elements, value)
	} else {
		s.elements[s.nextPushPosition] = value
	}
	s.nextPushPosition++
}

func (s *SliceStack[T]) PushSlice(values []T) {
	for _, value := range values {
		s.Push(value)
	}
}

func (s *SliceStack[T]) Pop() T {
	if s.IsEmpty() {
		panic("Invalid operation - stack is empty")
	}
	// TODO: Maybe zero the current value
	s.nextPushPosition -= 1
	return s.elements[s.nextPushPosition]
}

func (s *SliceStack[T]) PopAll() T {
	if s.IsEmpty() {
		panic("Invalid operation - stack is empty")
	}
	lastElement := s.elements[0]
	s.elements = make([]T, 0, len(s.elements))
	s.nextPushPosition = 0
	return lastElement
}

func (s *SliceStack[T]) Peek() T {
	if s.IsEmpty() {
		panic("Invalid operation - the stack is empty")
	}
	return s.elements[s.nextPushPosition-1]
}

func (s *SliceStack[T]) PeekIndex(index int) T {
	if s.Len() <= index {
		panic("Invalid operation - index out of bounds")
	}
	indexToFind := s.Len() - index - 1
	return s.elements[indexToFind]
}

func (s *SliceStack[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *SliceStack[T]) Len() int {
	return s.nextPushPosition
}
