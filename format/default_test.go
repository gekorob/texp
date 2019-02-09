package format_test

import (
	"reflect"
	"testing"

	"github.com/gekorob/texp/format"
	"github.com/gekorob/texp/log"
)

var expFatalS format.Style = format.Style{
	Color:         "\u001b[39m",
	Indentation:   "",
	SeverityLabel: "Fatal",
	Separator:     ": ",
	Termination:   "\n",
}

func TestDefaultFormatter(t *testing.T) {
	styleTests := []struct {
		name          string
		severity      log.Severity
		labelToVerify string
	}{
		{"Fatal", log.FATAL, "Fatal"},
		{"Error", log.ERROR, "Error"},
		{"Trace", log.TRACE, "Trace"},
		{"Test", log.TEST, "Test"},
		{"Info", log.INFO, "Info"},
	}
	df := format.NewDefaultStyle()

	for _, tt := range styleTests {
		t.Run(tt.name, func(t *testing.T) {
			expStyle := format.Style{
				Color:         "\u001b[39m",
				Indentation:   "",
				SeverityLabel: tt.labelToVerify,
				Separator:     ": ",
				Termination:   "\n",
			}
			sevS := df.BySeverity(tt.severity)
			if !reflect.DeepEqual(sevS, &expStyle) {
				t.Errorf("error getting style, exp: %v, got: %v", expStyle, sevS)
			}
		})
	}
}

func TestDefaultStyleWithoutTrace(t *testing.T) {
	s := format.NewDefaultStyle(format.WithNoTrace())

	m := s.BySeverity(log.TRACE)

	if m != nil {
		t.Errorf("expecting nil, got %v", m)
	}
}
