package queue

import "errors"

// Queue is the base data structure.
// It's internal implementation is empty interface.
type Queue struct {
	data []interface{}
}

// New returns an initialized Queue.
func New() *Queue {
	return new(Queue)
}

// Pushes a new value onto the queue.
func (q *Queue) Push(value interface{}) {
	q.data = append(q.data, value)
}

// Pops some data from this queue.
func (q *Queue) Pop() (interface{}, error) {
	if len(q.data) > 0 {
		v := q.data[0]
		q.data = q.data[1:]
		return v, nil
	}
	return nil, errors.New("no more items in queue.")
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return len(q.data)
}

// Empty returns if the queue is empty or not.
func (q *Queue) Empty() bool {
	if len(q.data) > 0 {
		return false
	}
	return true
}

// Queue with an implementation of ints.
type IntQueue struct {
	q Queue
}

// New returns an initialized IntQueue.
func NewInt() *IntQueue {
	return new(IntQueue)
}

// Pushes a new int value onto the queue.
func (iq *IntQueue) Push(value int) {
	iq.q.Push(value)
}

// Pops some int data from this queue.
func (iq *IntQueue) Pop() (int, error) {
	val, err := iq.q.Pop()
	if err != nil {
		return 0, err
	}
	return val.(int), nil
}

// Size returns the number of elements in the queue.
func (iq *IntQueue) Size() int {
	return iq.q.Size()
}

// Empty returns if the queue is empty or not.
func (iq *IntQueue) Empty() bool {
	return iq.q.Empty()
}

// Queue with an implementation of strings.
type StringQueue struct {
	q Queue
}

// New returns an initialized StringQueue.
func NewString() *StringQueue {
	return new(StringQueue)
}

// Pushes a new string value onto the queue.
func (sq *StringQueue) Push(value string) {
	sq.q.Push(value)
}

// Pops some string data from this queue.
func (sq *StringQueue) Pop() (string, error) {
	val, err := sq.q.Pop()
	if err != nil {
		return "", err
	}
	return val.(string), nil
}

// Size returns the number of elements in the queue.
func (sq *StringQueue) Size() int {
	return sq.q.Size()
}

// Empty returns if the queue is empty or not.
func (sq *StringQueue) Empty() bool {
	return sq.q.Empty()
}
