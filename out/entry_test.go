package out_test

import (
	"testing"

	"github.com/gekorob/texp/out"
)

func TestEmptyEntryStack(t *testing.T) {
	es := out.NewEntryStack()

	if count := es.Count(); count != 0 {
		t.Error("stack not empty or error creating it")
	}

	if entry := es.Pop(); entry != nil {
		t.Error("stack must be empty. expecting nil")
	}
}
