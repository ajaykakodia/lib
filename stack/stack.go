package stack

import "fmt"

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) Push(data T) {
	s.arr = append(s.arr, data)
}

func (s *Stack[T]) Pop() T {
	var ele T
	if s.IsEmpty() {
		fmt.Println("Stack is Empty! Nothing to Pop()")
		return ele
	}
	ele = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return ele
}

func (s *Stack[T]) Top() T {
	var ele T
	if s.IsEmpty() {
		fmt.Println("Stack is Empty!")
		return ele
	}
	return s.arr[len(s.arr)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.arr)
}
