package log

import (
	"sync"
)

// Severity is a form of message level classification.
type Severity int

const (
	// Info is for general purpose messages.
	Info Severity = iota
	// Test is for the test name.
	Test
	// Trace is used to follow the call stack.
	Trace
	// Error represents the error message.
	Error
	// Fatal is used for the blocking error message.
	Fatal
)

// The Msg is a categorized text content.
type Msg struct {
	Level   Severity
	Content string
}

// Queue collects all the Messages pushed into it and manages them in a FIFO way.
type Queue struct {
	mm    []Msg
	count int
	rw    *sync.RWMutex
}

// NewQueue factory method creates and initialize a new MsgQueue.
func NewQueue() *Queue {
	return &Queue{
		mm:    make([]Msg, 0, 2),
		count: 0,
		rw:    &sync.RWMutex{},
	}
}

// Count gives the number of element currently in the queue.
func (s *Queue) Count() int {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return s.count
}

func (s *Queue) String() string {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return ""
}
