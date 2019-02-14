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
func (d *DefaultStyle) BySeverity(sev log.Severity) *Style {
	if sev == log.TRACE && !d.trace {
		return nil
	}
	s := d.baseStyle
	s.SeverityLabel = d.labelBySeverity[sev]
	return &s
}

func (d *DefaultStyle) disableTrace() {
	d.trace = false
}

//WithNoTrace method, disable call stack tracing. Especially useful while developing and
// testing the assertions
func WithNoTrace() func(*DefaultStyle) {
	return func(d *DefaultStyle) {
		d.disableTrace()
	}
}
