package log

import "sync"

// Severity is a form of message level classification.
type Severity int

const (
	// Info is for general purpose messages.
	Info Severity = iota + 1
	// Test is for the test name.
	Test
	// Trace is used to follow the call stack.
	Trace
	// Error represents the error message.
	Error
	// Fatal is used for the blocking error message.
	Fatal
)

// The Message is a categorized text content.
type Message struct {
	Severity
	Content string
}

// Queue collects all the Messages pushed into it and manages them in a FIFO way.
// type Queue []Message
type Queue struct {
	mm    []Message
	count int
	idx   int
	rw    *sync.RWMutex
}

// NewQueue is a factory convenience method to initialize a
// new message queue.
func NewQueue() *Queue {
	return &Queue{
		mm:  make([]Message, 0, 2),
		idx: -1,
		rw:  &sync.RWMutex{},
	}
}

// Count gives the number of element currently in the queue.
func (q *Queue) Count() int {
	q.rw.RLock()
	defer q.rw.RUnlock()
	return len(q.mm)
}

// Push method add a message to the FIFO log queue
func (q *Queue) Push(msg Message) {
	q.rw.Lock()
	defer q.rw.Unlock()
	q.mm = append(q.mm, msg)
}

// Front method move to the first element of the queue and retrieve it
func (q *Queue) Front() (msg Message, ok bool) {
	q.rw.RLock()
	defer q.rw.RUnlock()

	if q.isEmpty() {
		ok = false
		return
	}
	q.idxToFirst()

	return q.mm[q.idx], true
}

func (q *Queue) Back() (msg Message, ok bool) {
	q.rw.RLock()
	defer q.rw.RUnlock()

	if q.isEmpty() {
		ok = false
		return
	}
	q.idxToLast()

	return q.mm[q.idx], true
}

// Next method move to the next messages sequentially one by one.
// Return true if there's an element to move to, otherwises gives you false.
func (q *Queue) Next() (msg Message, ok bool) {
	q.rw.RLock()
	defer q.rw.RUnlock()

	if !q.hasNext() {
		ok = false
		return
	}
	q.incIdx()
	msg = q.mm[q.idx]
	ok = true

	return
}

func (q *Queue) Prev() (msg Message, ok bool) {
	q.rw.RLock()
	defer q.rw.RUnlock()

	if !q.hasPrev() {
		ok = false
		return
	}
	q.decIdx()
	msg = q.mm[q.idx]
	ok = true

	return
}

func (q *Queue) isEmpty() bool {
	return q.Count() < 1
}

func (q *Queue) hasNext() bool {
	return q.idx < (q.Count() - 1)
}

func (q *Queue) hasPrev() bool {
	return q.idx > 0
}

func (q *Queue) idxToFirst() {
	q.idx = 0
}

func (q *Queue) idxToLast() {
	q.idx = len(q.mm) - 1
}

func (q *Queue) incIdx() {
	q.idx++
}

func (q *Queue) decIdx() {
	q.idx--
}

// TODO: can be useful to implement a Front method, a Back Method and a Prev
