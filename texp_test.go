package texp_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/gekorob/texp"
	"github.com/gekorob/texp/format"
)

func TestExpectCreation(t *testing.T) {
	expect := texp.Expect(t)

	if expect == nil {
		t.Error("expectation creation error")
	}

	if expT := expect(nil).T(); t != expT {
		t.Error("wrong test runner")
	}

	if !reflect.DeepEqual(expect.GlobalOutput(), os.Stdout) {
		t.Error("wrong global output")
	}

	if !reflect.DeepEqual(expect.GlobalStyle(), format.NewDefaultStyle()) {
		t.Error("wrong global style")
	}
}

// func TestExpectCreationWithOptions(t *testing.T) {
// 	var b strings.Builder
// 	mStyle := new(StyleMock)

// 	expect := texp.Expect(t, texp.OutTo(&b), texp.WithStyle(mStyle))

// 	if !reflect.DeepEqual(expect(nil).Out(), &b) {
// 		t.Error("wrong instance output")
// 	}

// 	if !reflect.DeepEqual(expect(nil).Style(), mStyle) {
// 		t.Error("wrong instance output")
// 	}
// }
