package out

// Severity is a form of message level classification
type Severity int

const (
	// Info is for general purpose messages
	Info Severity = iota
	// Test is for the test name
	Test
	// Trace is used to follow the call stack
	Trace
	// Error represents the error message
	Error
	// Fatal is used for the blocking error message
	Fatal
)

// The Entry is a categorized text content
type Entry struct {
	Level   Severity
	Content string
}

type EntryStack struct {
	count   int
	entries []Entry
}

func NewEntryStack() *EntryStack {
	return &EntryStack{}
}

func (s *EntryStack) Count() int {
	return s.count
}

func (s *EntryStack) Pop() *Entry {
	if s.Count() == 0 {
		return nil
	}
	s.count--
	return &s.entries[s.count]
}
