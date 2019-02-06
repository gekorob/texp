package log_test

import (
	"reflect"
	"testing"

	"github.com/gekorob/texp/log"
)

func TestLogBySeverity(t *testing.T) {
	sevTests := []struct {
		name    string
		factory func() log.Message
		expMsg  log.Message
	}{
		{
			name:    "Info",
			factory: func() log.Message { return log.Info("content") },
			expMsg:  log.Message{Severity: log.INFO, Content: "content"},
		},
		{
			name:    "Test",
			factory: func() log.Message { return log.Test("content") },
			expMsg:  log.Message{Severity: log.TEST, Content: "content"},
		},
		{
			name:    "Trace",
			factory: func() log.Message { return log.Trace("content") },
			expMsg:  log.Message{Severity: log.TRACE, Content: "content"},
		},
		{
			name:    "Error",
			factory: func() log.Message { return log.Error("content") },
			expMsg:  log.Message{Severity: log.ERROR, Content: "content"},
		},
		{
			name:    "Fatal",
			factory: func() log.Message { return log.Fatal("content") },
			expMsg:  log.Message{Severity: log.FATAL, Content: "content"},
		},
	}

	for _, tt := range sevTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.factory()
			if !reflect.DeepEqual(tt.expMsg, got) {
				t.Errorf("expecting %v, got %v", tt.expMsg, got)
			}
		})
	}
}
