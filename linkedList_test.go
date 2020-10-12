package datastructures

import "testing"
import "github.com/stretchr/testify/assert"

func TestLinkedList_PushFront(t *testing.T) {
	t.Run("push node on front of empty list", func(t *testing.T) {
		list := emptyList()
		testVal := 1

		list.PushFront(testVal)

		assert.Equal(t, list.Head.Value, testVal)
		assert.Equal(t, list.Tail.Value, testVal)
	})
	t.Run("push node on front of non-empty list", func(t *testing.T) {
		list := singleList()
		testVal := 2

		list.PushFront(testVal)

		assert.Equal(t, list.Head.Value, testVal)
		assert.NotEqual(t, list.Tail.Value, testVal)
	})
}

func TestLinkedList_TopFront(t *testing.T) {
	t.Run("top front from empty list", func(t *testing.T) {
		list := emptyList()

		_, err := list.TopFront()

		assert.Error(t, err)
	})
	t.Run("top front from non-empty list", func(t *testing.T) {
		list := multiList()

		val, _ := list.TopFront()

		assert.Equal(t, list.Head, val)
	})
}

func TestLinkedList_PopFront(t *testing.T) {
	t.Run("pop node from front of empty list", func(t *testing.T) {
		list := emptyList()

		err := list.PopFront()

		assert.Error(t, err)
	})
	t.Run("pop node from front of list with a single node", func(t *testing.T) {
		list := singleList()

		err := list.PopFront()

		assert.NoError(t, err)
		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
	})
	t.Run("pop node from front of list with multiple nodes", func(t *testing.T) {
		list := multiList()

		err := list.PopFront()

		assert.NoError(t, err)
		assert.NotNil(t, list.Head)
		assert.Equal(t, list.Head.Value, second.Value)
	})
}

func TestLinkedList_PushBack(t *testing.T) {
	t.Run("push node on back of empty list", func(t *testing.T) {
		list := emptyList()
		testVal := 1

		list.PushBack(testVal)

		assert.Equal(t, list.Head.Value, testVal)
		assert.Equal(t, list.Tail.Value, testVal)
	})
	t.Run("push node on back of non-empty list", func(t *testing.T) {
		list := singleList()
		testVal := 2

		list.PushBack(testVal)

		assert.NotEqual(t, list.Head.Value, testVal)
		assert.Equal(t, list.Tail.Value, testVal)
	})
}

func TestLinkedList_TopBack(t *testing.T) {
	t.Run("top back from empty list", func(t *testing.T) {
		list := emptyList()

		_, err := list.TopBack()

		assert.Error(t, err)
	})
	t.Run("top back from non-empty list", func(t *testing.T) {
		list := multiList()

		val, _ := list.TopBack()

		assert.Equal(t, list.Tail, val)
	})
}

func TestLinkedList_PopBack(t *testing.T) {
	t.Run("pop node from back of empty list", func(t *testing.T) {
		list := emptyList()

		err := list.PopBack()

		assert.Error(t, err)
	})
	t.Run("pop node from back of list with a single node", func(t *testing.T) {
		list := singleList()

		err := list.PopBack()

		assert.NoError(t, err)
		assert.Nil(t, list.Head)
		assert.Nil(t, list.Tail)
	})
	t.Run("pop nod from back of list with multiple nodes", func(t *testing.T) {
		list := multiList()

		err := list.PopBack()

		assert.NoError(t, err)
		assert.NotNil(t, list.Head)
		assert.Equal(t, list.Tail.Value, tail.Value)
	})
}

func TestLinkedList_AddAfter(t *testing.T) {
	t.Run("add after node on list with a single node", func(t *testing.T) {
		list := singleList()
		testVal := 2

		list.AddAfter(tail, testVal)

		assert.Equal(t, list.Head.Value, tail.Value)
		assert.Equal(t, list.Tail.Value, testVal)
	})
	t.Run("add after node on list with multiple nodes", func(t *testing.T) {
		list := multiList()
		testVal := 6

		list.AddAfter(second, testVal)

		assert.Equal(t, second.Next.Value, testVal)
	})
}

func TestLikedList_AddBefore(t *testing.T) {
	t.Run("add before first node on list", func(t *testing.T) {
		list := singleList()
		testVal := 2

		list.AddBefore(tail, testVal)

		assert.Equal(t, list.Head.Value, testVal)
		assert.Equal(t, list.Tail.Value, tail.Value)
	})
	t.Run("add before node on list with multiple nodes", func(t *testing.T) {
		list := multiList()
		testVal := 6

		list.AddBefore(third, testVal)

		assert.Equal(t, second.Next.Value, testVal)
	})
	t.Run("add before non-existing node on list with multiple nodes", func(t *testing.T) {
		list := multiList()
		testVal := 9

		list.AddBefore(nil, testVal)

		assert.Equal(t, tail.Next.Value, testVal)
	})
}

var tail = &Node{
	Value: 4, Next: nil,
}
var third = &Node{
	Value: 3, Next: tail,
}
var second = &Node{
	Value: 2, Next: third,
}
var head = &Node{
	Value: 1, Next: second,
}

func emptyList() *LinkedList {
	return &LinkedList{Head:nil, Tail: nil}
}

func singleList() *LinkedList {
	return &LinkedList{Head: tail, Tail: tail}
}

func multiList() *LinkedList {
	return &LinkedList{Head: head, Tail: tail}
}