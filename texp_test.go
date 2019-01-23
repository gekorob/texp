package texp_test

import (
	"os"
	"reflect"
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

	if !reflect.DeepEqual(expect.DefaultOutput(), os.Stdout) {
		t.Error("cannot read default output")
	}
}
