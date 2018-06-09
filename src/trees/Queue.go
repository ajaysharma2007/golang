package tree

import (
	"errors"
)

// Queue - Represent queue with the fixed capacity
type Queue struct {
	Head     int
	Tail     int
	Elements [3]*Node
}

// Enqueue - adds data to the end of the queue
func (queue *Queue) Enqueue(data *Node) (bool, error) {
	if queue.IsFull() {
		return false, errors.New("can't overflow the queue")
	}

	queue.Tail = (queue.Tail + 1) % len(queue.Elements)
	queue.Elements[queue.Tail] = data

	if queue.IsEmpty() {
		queue.Head = queue.Tail
	}
	return true, nil
}

// Dequeue - removes the element from the begining of the queue
func (queue *Queue) Dequeue() (*Node, error) {
	if queue.IsEmpty() {
		return nil, errors.New("can't return from the empty queue")
	}

	data := queue.Elements[queue.Head]
	if queue.Head == queue.Tail {
		queue.Head = -1
	} else {
		queue.Head = (queue.Head + 1) % len(queue.Elements)
	}
	return data, nil
}

// Peek - gives the view of the 1st element in the queue
func (queue *Queue) Peek() (*Node, error) {
	if queue.IsEmpty() {
		return nil, errors.New("can't peek into an empty queue")
	}

	return queue.Elements[queue.Head], nil
}

// IsEmpty - returns true if queue is empty
func (queue *Queue) IsEmpty() bool {
	return queue.Head == -1
}

// IsFull = returns true if queue is full
func (queue *Queue) IsFull() bool {
	nextIndex := (queue.Tail + 1) % len(queue.Elements)
	return nextIndex == queue.Head
}
