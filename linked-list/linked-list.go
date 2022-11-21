package linkedlist

import "fmt"

type ILinkedList struct {
	head *INode
}

type INode struct {
	data int
	next *INode
}

func NewILinkedList() *ILinkedList {
	return &ILinkedList{}
}

func NewINode(num int) *INode {
	return &INode{data: num}
}

// Print => Print all node of linked list
func (ll *ILinkedList) Print() {
	if ll.head == nil {
		fmt.Println("No node for print")
		return
	}
	currentNode := ll.head
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("none")
}

// PrintParticularNode => Print a particular node of linked list
func (ll *ILinkedList) PrintParticularNode(index int) {
	i := 0
	cn := ll.head
	for cn != nil {
		if i == index {
			fmt.Printf("Node at Index %d is %v\n", index, cn)
			return
		}
		i++
		cn = cn.next
	}
	fmt.Printf("No node to print at index %d\n", index)
}

func (ll *ILinkedList) Length() int {
	count := 0
	currentNode := ll.head
	for currentNode != nil {
		currentNode = currentNode.next
		count++
	}
	return count
}

// AddNode => Add new node at last Node of list
func (ll *ILinkedList) AddNode(num int) {
	newNode := NewINode(num)
	if ll.head == nil {
		ll.head = newNode
		return
	}
	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
}

// AddNodeAtIndex=> add node at particular index
func (ll *ILinkedList) AddNodeAtIndex(num, index int) {
	if index < 0 {
		fmt.Println("Please provide index greater than or equal to 0")
		return
	}
	newNode := NewINode(num)

	var previousNode *INode
	currentNode := ll.head

	for index > 0 && currentNode != nil {
		previousNode = currentNode
		currentNode = currentNode.next
		index--
	}

	if index > 0 && previousNode != nil {
		fmt.Println("Index out of range")
		return
	}
	newNode.next = currentNode
	if previousNode == nil {
		ll.head = newNode
		return
	}

	previousNode.next = newNode
}

// DeleteNode => Delete ode Iterative
func (ll *ILinkedList) DeleteNode(index int) {
	if index == 0 {
		ll.head = ll.head.next
		return
	}
	i := 0
	cn := ll.head
	for i < index-1 && cn != nil {
		i++
		cn = cn.next
	}
	if cn == nil {
		fmt.Println("No node to delete on index ", index)
		return
	}
	cn.next = cn.next.next
}
