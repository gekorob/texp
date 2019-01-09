package out_test

import (
	"testing"

	"github.com/gekorob/texp/out"
)

func TestEmptyMsgQueue(t *testing.T) {
	mq := out.NewMsgQueue()

	if count := mq.Count(); count != 0 {
		t.Error("queue not empty or error creating it")
	}

	if _, found := mq.Pop(); found {
		t.Error("queue must be empty. expecting found false")
	}
}
