package queue

import (
	"errors"
)

type queue[T any] struct {
	q     []T
	front int
	rear  int
}

func Newqueue[T any](qLength int) *queue[T] {
	if qLength == -1 {
		// qLength is -1 means the user did provide any especific length so we assign a big length.
		qLength = 1000
	}
	return &queue[T]{
		q:     make([]T, qLength),
		front: -1,
		rear:  -1,
	}
}

func (q *queue[T]) IsEmpty() bool {
	if q.front == -1 && q.rear == -1 {
		return true
	}
	return false
}

func (q *queue[T]) Enqueue(element T) (T, error) {
	if q.IsEmpty() {
		q.rear = 0
		q.front = 0
		q.q[q.rear] = element
		return element, nil
	}
	q.rear = (q.rear + 1) % len(q.q)
	if q.rear == q.front {
		return element, errors.New("queue is full.")
	}
	q.q[q.rear] = element
	return element, nil
}

func (q *queue[T]) Dequeue() (T, error) {
	var element T
	if q.IsEmpty() {
		return element, errors.New("queue is empty.")
	}
	element = q.q[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front + 1) % len(q.q)
	}
	return element, nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
