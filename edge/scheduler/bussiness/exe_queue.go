package bussiness

import "time"

type node struct {
	data interface{}
	next *node
}

// Queue -
type Queue struct {
	head         *node
	end          *node
	len          int
	LastScanTime time.Time
}

// NewQueue -
func NewQueue() *Queue {
	q := &Queue{nil, nil, 0, time.Now()}
	return q
}

// Push -
func (q *Queue) Push(data interface{}) {
	n := &node{data: data, next: nil}
	if q.end == nil {
		q.head = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	q.len++
	return
}

// Pop -
func (q *Queue) Pop() (interface{}, bool) {
	if q.head == nil {
		return nil, false
	}
	data := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}
	q.len--
	return data, true
}

// GetSize -
func (q *Queue) GetSize() int {
	return q.len
}
