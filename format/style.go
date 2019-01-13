package format

import (
	"github.com/gekorob/texp/log"
)

// The Style struct contains the basic styling informations for
// log messages, including separator and Severity label translation.
type Style struct {
	Color         string
	Indentation   string
	SeverityLabel string
	Separator     string
	Termination   string
}

// The Styler interface defines a contract for implementations to
// get the style for each entry of the log stack basing on message Severity
type Styler interface {
	BySeverity(sev log.Severity) *Style
}
