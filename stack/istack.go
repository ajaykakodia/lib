package stack

import "fmt"

type IStack struct {
	arr []int
}

func (s *IStack) Push(data int) {
	s.arr = append(s.arr, data)
}

func (s *IStack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty! Nothing to Pop()")
		return -1
	}
	ele := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return ele
}

func (s *IStack) Top() int {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty!")
		return -1
	}
	return s.arr[len(s.arr)-1]
}

func (s *IStack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *IStack) Size() int {
	return len(s.arr)
}
