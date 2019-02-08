package format

import (
	"github.com/gekorob/texp/log"
)

// DefaultStyle is a simple style manager with all the style parameters set to default.
// Only SeverityLabel change according to the severity, the other parameters remain the same.
type DefaultStyle struct {
	trace           bool
	baseStyle       Style
	labelBySeverity map[log.Severity]string
}

// NewDefaultStyle factory method creates a brand new DefaultStyle manager.
func NewDefaultStyle(options ...func(*DefaultStyle)) *DefaultStyle {
	d := DefaultStyle{
		trace: true,
		baseStyle: Style{
			Color:       "\u001b[39m",
			Indentation: "",
			Separator:   ": ",
			Termination: "\n",
		},
		labelBySeverity: map[log.Severity]string{
			log.FATAL: "Fatal",
			log.ERROR: "Error",
			log.TRACE: "Trace",
			log.INFO:  "Info",
			log.TEST:  "Test",
		},
	}

	for _, o := range options {
		o(&d)
	}

	return &d
}

// BySeverity method returns the correct style for the selected Severity
func (d *DefaultStyle) BySeverity(s log.Severity) *Style {
	if s == log.TRACE && !d.trace {
		return nil
	}
	d.baseStyle.SeverityLabel = d.labelBySeverity[s]
	return &d.baseStyle
}

func (d *DefaultStyle) disableTrace() {
	d.trace = false
}

func WithNoTrace() func(*DefaultStyle) {
	return func(d *DefaultStyle) {
		d.disableTrace()
	}
}
