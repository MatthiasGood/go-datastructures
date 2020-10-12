package datastructures

import (
	"errors"
)

// Node of a singly linked list
type Node struct {
	Value 	interface{}
	Next 	*Node
}

// Singly linked list
type LinkedList struct {
	Head 	*Node
	Tail 	*Node
}

// PushFront adds an element to the front of the linked list.
func (l *LinkedList) PushFront(val interface{}) {
	node := &Node{Value: val, Next: l.Head}
	l.Head = node
	if l.Tail == nil {
		l.Tail = l.Head
	}
}

// TopFront returns the first element of the linked list.
// It returns an error if the list is empty.
func (l *LinkedList) TopFront() (*Node, error) {
	if l.Empty() {
		return nil, errors.New("can't top from empty list")
	}
	return l.Head, nil
}

// PopFront removes the first element of the linked list.
// It returns an error if the list is empty.
func (l *LinkedList) PopFront() error {
	if l.Empty() {
		return errors.New("can't remove from empty list")
	}
	l.Head = l.Head.Next
	if l.Head == nil {
		l.Tail = nil
	}
	return nil
}

// PushBack adds an element to the back of the linked list.
func (l *LinkedList) PushBack(val interface{}) {
	node := &Node{Value: val, Next: nil}
	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}

// TopBack returns the last element of the linked list.
// It returns an error if the list is empty.
func (l *LinkedList) TopBack() (*Node, error) {
	if l.Empty() {
		return nil, errors.New("can't top from empty list")
	}
	return l.Tail, nil
}

// PopBack removes the last element of the linked list.
// It returns an error if the list is empty.
func (l *LinkedList) PopBack() error {
	if l.Empty() {
		return errors.New("can't remove from empty list")
	}
	if l.Head.Value == l.Tail.Value {
		l.Head = nil
		l.Tail = nil
		return nil
	}
	currentHead := l.Head
	for currentHead.Next.Next != nil {
		currentHead = currentHead.Next
	}
	l.Tail = currentHead
	currentHead.Next = nil
	return nil
}

// AddAfter adds an element to the linked list after the given node.
func (l *LinkedList) AddAfter(node *Node, val interface{}) {
	newNode := &Node{Value: val, Next: node.Next}
	node.Next = newNode
	if l.Tail.Value == node.Value {
		l.Tail = newNode
	}
}

// AddBefore adds an element to the linked list before the given node.
func (l *LinkedList) AddBefore(node *Node, val interface{}) {
	if l.Empty() || l.Head == node {
		l.PushFront(val)
	} else {
		var prev *Node
		cur := l.Head
		for cur != nil && cur != node {
			prev = cur
			cur = cur.Next
		}
		if cur != nil {
			prev.Next = &Node{Value: val, Next: cur}
		} else {
			l.PushBack(val)
		}
	}
}

// Empty checks if the linked list is empty.
func (l *LinkedList) Empty() bool {
	return l.Head == nil
}
