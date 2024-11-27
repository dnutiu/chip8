package main

import "errors"

// Stack is a simple stack implementation that holds uint16 data.
type Stack[T any] struct {
	storage   []T
	stackSize uint32
}

// NewStack creates a new Stack instance.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		storage:   make([]T, 0, 128),
		stackSize: 0,
	}
}

// Push pushes data into the stack.
func (s *Stack[T]) Push(data T) {
	s.stackSize += 1
	s.storage = append(s.storage, data)
}

// Pop pops an element from the stack.
func (s *Stack[T]) Pop() (T, error) {
	var whatever T
	if s.stackSize == 0 {
		return whatever, errors.New("stack is empty")
	}
	data := s.storage[s.stackSize-1]
	s.stackSize -= 1
	return data, nil
}

// Peek peeks at the latest data stored on the stack.
func (s *Stack[T]) Peek() T {
	return s.storage[s.stackSize-1]
}

// Size returns the current size of the stack.
func (s *Stack[T]) Size() uint32 {
	return s.stackSize
}

// IsEmpty returns true if the stack is empty, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return s.stackSize == 0
}
