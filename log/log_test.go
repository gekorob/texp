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
		t.Error("no next for an empty queue")
	}

	if _, ok := q.Prev(); ok != false {
		t.Error("no previous for an empty queue")
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

func TestPositionToFront(t *testing.T) {
	q := log.NewQueue()

	_, ok := q.Front()
	if ok != false {
		t.Error("unable to position at the beginning of an empty queue")
	}

	expFirstMsg := log.Message{Severity: log.Test, Content: "this is the first message"}
	q.Push(expFirstMsg)

	otherMsg := log.Message{Severity: log.Info, Content: "this is another message"}
	q.Push(otherMsg)

	msg, _ := q.Front()
	if !reflect.DeepEqual(expFirstMsg, msg) {
		t.Error("uncorrect positioning to the front")
	}

	msg, _ = q.Next()
	if !reflect.DeepEqual(otherMsg, msg) {
		t.Error("unexpected message after the fist one")
	}
}

func TestPositionToTheBack(t *testing.T) {
	q := log.NewQueue()

	_, ok := q.Back()
	if ok != false {
		t.Error("unable to position at the end of an empty queue")
	}

	otherMsg := log.Message{Severity: log.Info, Content: "this is another message"}
	q.Push(otherMsg)

	expLastMsg := log.Message{Severity: log.Test, Content: "this is the last message"}
	q.Push(expLastMsg)

	msg, _ := q.Back()
	if !reflect.DeepEqual(expLastMsg, msg) {
		t.Error("uncorrect positioning to the back")
	}

	msg, _ = q.Prev()
	if !reflect.DeepEqual(otherMsg, msg) {
		t.Error("unexpected message before the last one")
	}
}

func TestPrevAndNextWithOneElement(t *testing.T) {
	q := log.NewQueue()

	q.Push(log.Message{Severity: log.Test, Content: "element"})
	q.Front()

	if _, ok := q.Prev(); ok != false {
		t.Error("no prev before the front")
	}

	if _, ok := q.Next(); ok != false {
		t.Error("no next after the back")
	}
}

func TestIterationOnQueue(t *testing.T) {
	q := log.NewQueue()

	expFirstMsg := log.Message{Severity: log.Test, Content: "this is the first message"}
	expLastMsg := log.Message{Severity: log.Test, Content: "this is the last message"}

	q.Push(expFirstMsg)
	q.Push(expLastMsg)

	if q.Count() != 2 {
		t.Error("error adding messages to the queue")
	}

	msgs := make([]log.Message, 0, q.Count())
	for m, ok := q.Front(); ok; m, ok = q.Next() {
		msgs = append(msgs, m)
	}

	if len(msgs) != q.Count() {
		t.Error("error iterating on queue")
	}

	if !reflect.DeepEqual(msgs, []log.Message{expFirstMsg, expLastMsg}) {
		t.Error("unexpected msgs sequence")
	}
}
