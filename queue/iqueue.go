package queue

import "fmt"

type IQueue struct {
	arr []int
}

func (q *IQueue) Enqueue(data int) {
	q.arr = append(q.arr, data)
}

func (q *IQueue) Dequeue() int {
	if len(q.arr) == 0 {
		fmt.Println("Nothing to dequeue!!")
		return -1
	}
	ele := q.arr[0]
	q.arr = q.arr[1:]
	return ele
}

func (q *IQueue) Front() int {
	if len(q.arr) == 0 {
		fmt.Println("Queue is empty..")
		return -1
	}
	return q.arr[0]
}

func (q *IQueue) Size() int {
	return len(q.arr)
}

func (q *IQueue) IsEmpty() bool {
	return q.Size() == 0
}
