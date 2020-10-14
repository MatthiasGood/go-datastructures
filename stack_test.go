package datastructures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Size(t *testing.T) {
	t.Run("size of empty stack", func(t *testing.T) {
		stack := emptyStack()

		size := stack.Size()

		assert.Zero(t, size)
	})
	t.Run("size of non-empty stack", func(t *testing.T) {
		stack := nonEmptyStack()

		size := stack.Size()

		assert.NotZero(t, size)
	})
}

func TestStack_Empty(t *testing.T) {
	t.Run("check empty on empty stack", func(t *testing.T) {
		stack := emptyStack()

		empty := stack.Empty()

		assert.True(t, empty)
	})
	t.Run("check empty on non-empty stack", func(t *testing.T) {
		stack := nonEmptyStack()

		empty := stack.Empty()

		assert.False(t, empty)
	})
}

func TestStack_Top(t *testing.T) {
	t.Run("top element of empty stack", func(t *testing.T) {
		stack := emptyStack()

		_, err := stack.Top()

		assert.Error(t, err)
	})
	t.Run("top element of non-empty stack", func(t *testing.T) {
		stack := nonEmptyStack()

		top, _ := stack.Top()

		assert.Equal(t, top, topElement)
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("push element to empty stack", func(t *testing.T) {
		stack := emptyStack()
		newTop := &StackElement{Value: 4}

		stack.Push(newTop)

		assert.Equal(t, stack.Size(), 1)
		assert.False(t, stack.Empty())
		top, _ := stack.Top()
		assert.Equal(t, top, newTop)

	})
	t.Run("push element to non-empty stack", func(t *testing.T) {
		stack := nonEmptyStack()
		newTop := &StackElement{Value: 4}

		stack.Push(newTop)

		assert.Equal(t, stack.Size(), 4)
		assert.False(t, stack.Empty())
		top, _ := stack.Top()
		assert.Equal(t, top, newTop)
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("pop element of empty stack", func(t *testing.T) {
		stack := emptyStack()

		_, err := stack.Pop()

		assert.Error(t, err)
	})
	t.Run("pop element of non-empty stack", func(t *testing.T) {
		stack := nonEmptyStack()

		top, _ := stack.Pop()

		assert.Equal(t, top, topElement)
		assert.Equal(t, stack.Size(), 2)
		newTop, _ := stack.Top()
		assert.Equal(t, newTop, &StackElement{Value: 2})
	})
}

type StackElement struct {
	Value 	int
}

var topElement = &StackElement{Value: 3}

func emptyStack() *Stack {
	return &Stack{arr: make([]interface{}, 0)}
}

func nonEmptyStack() *Stack {
	return &Stack{arr: []interface{}{
		&StackElement{Value: 1},
		&StackElement{Value: 2},
		topElement,
	}}
}

