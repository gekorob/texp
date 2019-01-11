package out_test

import (
	"fmt"
	"testing"

	"github.com/gekorob/texp/out"
)

func TestEmptyLogQueue(t *testing.T) {
	lq := out.NewLogQueue()

	if count := lq.Count(); count != 0 {
		t.Error("queue not empty or error creating it")
	}

	if fmt.Sprint(lq) != "" {
		t.Error("wrong queue output")
	}
}
