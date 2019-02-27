package texp_test

import (
	"testing"

	"github.com/gekorob/texp"
	"github.com/gekorob/texp/mock"
)

func TestNotToBeTrue(t *testing.T) {
	tM := mock.NewTMock()
	expect := texp.Expect(tM)

	expect(false).Not().ToBeTrue()

	if testFailInvoked(tM) {
		t.Error("Unexpected error")
	}
}

func TestNotToBeNil(t *testing.T) {
	tM := mock.NewTMock()
	expect := texp.Expect(tM)

	expect(1).Not().ToBeNil()

	if testFailInvoked(tM) {
		t.Error("Unexpected error")
	}
}

func TestDisableNegAfterAssert(t *testing.T) {
	tM := mock.NewTMock()
	expect := texp.Expect(tM)

	expect(false).Not().ToBeTrue()

	expect(false).ToBeTrue()

	if !testFailInvoked(tM) {
		t.Error("Neutral behaviour not set after negative assertion")
	}
}
