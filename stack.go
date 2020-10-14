package datastructures

import "errors"

// Stack implemented with a slice
type Stack struct {
	arr []interface{}
}

// Size returns the number of elements within a stack
func (s *Stack) Size() int {
	return len(s.arr)
}

// Empty checks if there are any elements within a stack.
// It returns true, if the stack is empty, otherwise false.
func (s *Stack) Empty() bool {
	return len(s.arr) == 0
}

// Top returns the top element of the stack.
// It returns an error if the stack is empty.
func (s *Stack) Top() (interface{}, error) {
	if s.Empty() {
		return 0, errors.New("can't get top of empty stack")
	}
	return s.arr[s.Size() - 1], nil
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(obj interface{}) {
	s.arr = append(s.arr, obj)
}

// Push removes the top element of the stack and returns the removed element.
// It return s an error if the list is empty.
func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("can't get top of empty stack")
	}
	top, _ := s.Top()
	s.arr = s.arr[:s.Size()-1]
	return top, nil
}