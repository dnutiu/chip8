package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_NewStack(t *testing.T) {
	// Given
	stack := NewStack[string]()

	// Then
	assert.NotNil(t, stack)
}

func TestStack_BasicOperations(t *testing.T) {
	// Given
	stack := NewStack[uint16]()

	// Then
	stack.Push(uint16(2))
	stack.Push(uint16(13))

	assert.Equal(t, uint32(2), stack.Size())
	assert.Equal(t, uint16(13), stack.Peek())

	data, err := stack.Pop()
	assert.Equal(t, uint16(13), data)
	assert.Equal(t, uint32(1), stack.Size())
	assert.NoError(t, err)
	data, err = stack.Pop()
	assert.Equal(t, uint16(2), data)
	assert.Equal(t, uint32(0), stack.Size())
	assert.NoError(t, err)

	data, err = stack.Pop()
	assert.Error(t, err)
}

func TestStack_IsEmpty(t *testing.T) {
	// Given
	stack := NewStack[int]()

	// Then
	assert.True(t, stack.IsEmpty())

	stack.Push(1)
	assert.False(t, stack.IsEmpty())

	_, _ = stack.Pop()
	assert.True(t, stack.IsEmpty())
}
