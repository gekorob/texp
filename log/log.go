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
