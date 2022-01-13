package queue

import "testing"

func TestQueueArray(t *testing.T) {

	t.Run("Test enqueue", func(t *testing.T) {
		q := &ListQueue{}
		q.Enqueue(1)
		front, err := q.Front()
		if err != nil {
			t.Errorf("Got error: %v", err)
		}
		if front != 1 {
			t.Errorf("Expect get %v but got: %v", 1, front)
		}
		if l := q.Len(); l != 1 {
			t.Errorf("Queue length should be: %v but got: %v", 1, l)
		}
	})

	t.Run("Test dequeue", func(t *testing.T) {
		q := &ListQueue{}
		q.Enqueue(1)
		q.Enqueue(2)

		// FIFO so the front should equal to 1
		front, err := q.Dequeue()
		if err != nil {
			t.Errorf("Got error: %v", err)
		}
		if front != 1 {
			t.Errorf("Expect get %v but got: %v", 1, front)
		}
		// remove second element
		q.Dequeue()
		if l := q.Len(); l != 0 {
			t.Errorf("Queue length should be: 0 but got: %v", l)
		}
	})
}
