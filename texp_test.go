package texp_test

import (
	"testing"

	"github.com/gekorob/texp"
)

func TestExpectCreation(t *testing.T) {
	expect := texp.Expect(t)

	if expect == nil {
		t.Error("expectation creation error")
	}

	if expT := expect(nil).T(); t != expT {
		t.Error("wrong test runner")
	}
}
