package log_test

import (
	"reflect"
	"testing"

	"github.com/gekorob/texp/log"
)

func TestEmptyLogQueue(t *testing.T) {
	q := log.NewQueue()

	if q.Count() != 0 {
		t.Error("queue not empty or error creating it")
	}

	if _, ok := q.Next(); ok != false {
		t.Error("empty queue must not have a next")
	}

}

func TestAddMessage(t *testing.T) {
	q := log.NewQueue()

	expMsg := log.Message{Severity: log.Error, Content: "error message added"}
	q.Push(expMsg)

	if q.Count() != 1 {
		t.Errorf("error in queue length, exp: 1, got: %v", q.Count())
	}

	msg, ok := q.Next()
	if !ok {
		t.Error("expecting to have a next element")
	}

	if !reflect.DeepEqual(expMsg, msg) {
		t.Error("unexpected message retreived")
	}
}
