package stack

import "fmt"

type node struct {
	data string
	next *node
}

type SStack struct {
	head  *node
	count int
}

func newNode(data string) *node {
	return &node{data: data}
}

func (s *SStack) Push(data string) {
	newNode := newNode(data)
	newNode.next = s.head
	s.head = newNode
	s.count++
}

func (s *SStack) Pop() string {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty! nothing to pop..")
		return ""
	}
	data := s.head.data
	s.head = s.head.next
	s.count--
	return data
}

func (s *SStack) Top() string {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty! nothing to pop..")
		return ""
	}
	return s.head.data
}

func (s *SStack) IsEmpty() bool {
	return s.head == nil
}

func (s *SStack) Size() int {
	return s.count
}
