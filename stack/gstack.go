package stack

import "fmt"

type GStack struct {
	arr []interface{}
}

func (s *GStack) Push(data interface{}) {
	s.arr = append(s.arr, data)
}

func (s *GStack) Pop() interface{} {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty! Nothing to Pop()")
		return -1
	}
	ele := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return ele
}

func (s *GStack) Top() interface{} {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty!")
		return -1
	}
	return s.arr[len(s.arr)-1]
}

func (s *GStack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *GStack) Size() int {
	return len(s.arr)
}
