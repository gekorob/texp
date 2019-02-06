package format

import (
	"github.com/gekorob/texp/log"
)

// DefaultStyle is a simple style manager with all the style parameters set to default.
// Only SeverityLabel change according to the severity, the other parameters remain the same.
type DefaultStyle struct {
	baseStyle       Style
	labelBySeverity map[log.Severity]string
}

// NewDefaultStyle factory method creates a brand new DefaultStyle manager.
func NewDefaultStyle() *DefaultStyle {
	return &DefaultStyle{
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
}

// BySeverity method returns the correct style for the selected Severity
func (d *DefaultStyle) BySeverity(s log.Severity) *Style {
	d.baseStyle.SeverityLabel = d.labelBySeverity[s]
	return &d.baseStyle
}
