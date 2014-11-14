package ds

import "errors"

type Queue struct {
	data []interface{} // will try list vs slice.
}

func (q *Queue) Push(value interface{}) {
	q.data = append(q.data, value)
}

func (q *Queue) Pop() (interface{}, error) {
	if len(q.data) > 0 {
		s := q.data[0]
		q.data = q.data[1:]
		return s, nil
	}
	return nil, errors.New("no more items in queue.")
}

type IntQueue struct {
	q Queue
}

func (iq *IntQueue) Push(value int) {
	iq.q.Push(value)
}

func (iq *IntQueue) Pop() (int, error) {
	val, err := iq.q.Pop()
	if err != nil {
		return 0, err
	}
	return val.(int), nil
}

type StringQueue struct {
	q Queue
}

func (sq *StringQueue) Push(value string) {
	sq.q.Push(value)
}

func (sq *StringQueue) Pop() (string, error) {
	val, err := sq.q.Pop()
	if err != nil {
		return "", err
	}
	return val.(string), nil
}
