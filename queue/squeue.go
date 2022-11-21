package queue

import "fmt"

type Node struct {
	data string
	next *Node
}

type SQueue struct {
	head  *Node
	tail  *Node
	count int
}

func (q *SQueue) Enqueue(data string) {
	newNode := &Node{data: data}
	q.count++
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		return
	}
	q.tail.next = newNode
	q.tail = newNode
}

func (q *SQueue) Dequeue() string {
	if q.IsEmpty() {
		fmt.Print("Nothing to Dequeue!!")
		return ""
	}
	ele := q.head
	q.head = q.head.next
	return ele.data
}

func (q *SQueue) Front() string {
	if q.IsEmpty() {
		fmt.Println("Queue is empty...")
		return ""
	}
	return q.head.data
}

func (q *SQueue) Size() int {
	return q.count
}

func (q *SQueue) IsEmpty() bool {
	return q.head == nil
}
