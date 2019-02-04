package texp_test

import (
	"strings"
	"testing"

	"github.com/gekorob/texp"
	"github.com/gekorob/texp/conf"
	"github.com/gekorob/texp/mock"
)

func TestSimpleToBeTrue(t *testing.T) {
	var b strings.Builder
	tMock := mock.NewTMock()

	expect := texp.Expect(tMock, conf.OutputTo(&b))

	expect(false).ToBeTrue()

	if c, ok := tMock.CallsTo("Fail"); ok && c != 1 {
		t.Error("Expecting call to Fail one time")
	}

	expLog := "Test: Test\nError: \n"
	if b.String() != expLog {
		t.Errorf("Expecting %s, got %s", expLog, b.String())
	}
}
