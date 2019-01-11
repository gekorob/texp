package out

// The Style struct contains the basic styling informations for
// log messages, including separator and Severity label translation.
type Style struct {
	Color       string
	Indentation string
	SevLabel    string
	Separator   string
	Termination string
}

// The Formatter interface defines a contract for implementations to
// format the log stack basing on message Severity.
// according to the category of each entry.
type Formatter interface {
	StyleBySeverity(sev Severity) Style
}
