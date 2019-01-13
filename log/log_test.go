package log_test

import (
	"fmt"
	"testing"

	"github.com/gekorob/texp/log"
)

func TestEmptyLogQueue(t *testing.T) {
	q := log.NewQueue()

	if count := q.Count(); count != 0 {
		t.Error("queue not empty or error creating it")
	}

	if fmt.Sprint(q) != "" {
		t.Error("wrong queue output")
	}
}
