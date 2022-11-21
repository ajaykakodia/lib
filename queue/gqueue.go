package queue

import "fmt"

type GQueue struct {
	arr []interface{}
}

func (q *GQueue) Enqueue(data interface{}) {
	q.arr = append(q.arr, data)
}

func (q *GQueue) Dequeue() interface{} {
	if len(q.arr) == 0 {
		fmt.Println("Nothing to dequeue!!")
		return -1
	}
	ele := q.arr[0]
	q.arr = q.arr[1:]
	return ele
}

func (q *GQueue) Front() interface{} {
	if len(q.arr) == 0 {
		fmt.Println("Queue is empty..")
		return -1
	}
	return q.arr[0]
}

func (q *GQueue) Size() int {
	return len(q.arr)
}

func (q *GQueue) IsEmpty() bool {
	return q.Size() == 0
}
