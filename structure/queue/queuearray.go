// 队列是一种线性结构，它遵循执行操作的特定顺序。顺序是先进先出 (FIFO)
package queue

import "errors"

type ListQueue []interface{}

func (q *ListQueue) Enqueue(v interface{}) {
	*q = append(*q, v)
}

func (q *ListQueue) Dequeue() (interface{}, error) {
	if !q.IsEmpty() {
		v := (*q)[0]
		*q = (*q)[1:]
		return v, nil
	}
	return nil, errors.New("Queue is empty")
}

// Front returns the first element of queue if or nil if queue is empty
func (q ListQueue) Front() (interface{}, error) {
	if !q.IsEmpty() {
		return q[0], nil
	}
	return nil, errors.New("Queue is empty")
}

// Front returns the last element of queue if or nil if queue is empty
func (q ListQueue) Back() (interface{}, error) {
	if l := q.Len(); l > 0 {
		return q[l-1], nil
	}
	return nil, errors.New("Queue is empty")
}

func (q ListQueue) Len() int {
	return len(q)
}

func (q ListQueue) IsEmpty() bool {
	return len(q) == 0
}
