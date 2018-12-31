package log

// Severity represents the log level
type Severity int

const (
	// Info level
	Info Severity = iota
	// Warning level
	Warning
	//Error level
	Error
	//Fatal level
	Fatal
)

// Entry contains the content to log and the severity level
type Entry struct {
	Level      Severity
	LogContent string
}

// The Formatter interface has a simple method to format several log entries
type Formatter interface {
	Format(ee []Entry) string
}

type Writer struct {
	f Formatter
}

func NewWriter(f Formatter) *Writer {
	return &Writer{
		f: f,
	}
}
