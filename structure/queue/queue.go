package queue

// Queue is an interface of queue
type Queue interface {
	Enqueue(v interface{})

	Dequeue() (interface{}, error)

	// Front returns the first element of queue if or nil if queue is empty
	Front() (interface{}, error)

	// Front returns the last element of queue if or nil if queue is empty
	Back() (interface{}, error)

	Len() int

	IsEmpty() bool
}
