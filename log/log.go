package log

// Severity is a form of message level classification.
type Severity int

const (
	// INFO is for general purpose messages.
	INFO Severity = iota + 1
	// TEST is for the test name.
	TEST
	// TRACE is used to follow the call stack.
	TRACE
	// ERROR represents the error message.
	ERROR
	// FATAL is used for the blocking error message.
	FATAL
)

// The Message is a categorized text content.
type Message struct {
	Severity
	Content string
}

func Info(content string) Message {
	return createMessage(INFO, content)
}

func Test(content string) Message {
	return createMessage(TEST, content)
}

func Trace(content string) Message {
	return createMessage(TRACE, content)
}

func Error(content string) Message {
	return createMessage(ERROR, content)
}

func Fatal(content string) Message {
	return createMessage(FATAL, content)
}

func createMessage(sev Severity, cnt string) Message {
	return Message{
		Severity: sev,
		Content:  cnt,
	}
}

//Queuer interface define the method for adding messages
type Queuer interface {
	Push(msg Message)
	Count() int
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
