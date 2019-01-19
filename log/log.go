package log

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

//Queuer interface define the method for adding messages
type Queuer interface {
	Push(msg Message)
}

// FwIterator is a contract to define capabilities to position at the
// beginning and move forward
type FwIterator interface {
	Front() (Message, bool)
	Next() (Message, bool)
}

// RevIterator defines methods to position at the end and move backward
type RevIterator interface {
	Back() (Message, bool)
	Prev() (Message, bool)
}
