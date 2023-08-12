package queue

import "fmt"

type Queue[T any] struct {
	arr []T
}

func (q *Queue[T]) Enqueue(data T) {
	q.arr = append(q.arr, data)
}

func (q *Queue[T]) Dequeue() T {
	var ele T
	if len(q.arr) == 0 {
		fmt.Println("Nothing to dequeue!!")
		return ele
	}
	ele = q.arr[0]
	q.arr = q.arr[1:]
	return ele
}

func (q *Queue[T]) Front() T {
	var ele T
	if len(q.arr) == 0 {
		fmt.Println("Queue is empty..")
		return ele
	}
	return q.arr[0]
}

func (q *Queue[T]) Size() int {
	return len(q.arr)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}
