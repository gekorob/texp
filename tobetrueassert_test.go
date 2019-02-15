package texp_test

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/gekorob/texp"
	"github.com/gekorob/texp/conf"
	"github.com/gekorob/texp/format"
	"github.com/gekorob/texp/mock"
)

var trueTests = []struct {
	N string
	T func(texp.ExpBuilder)
	A func(*mock.TMock) bool
}{
	{
		N: "on equality",
		T: func(expect texp.ExpBuilder) {
			expect(10 == 10).ToBeTrue()
		},
		A: func(tM *mock.TMock) bool {
			return !testFailInvoked(tM)
		},
	},
	{
		N: "on false equality",
		T: func(expect texp.ExpBuilder) {
			expect(1 == 10).ToBeTrue()
		},
		A: func(tM *mock.TMock) bool {
			return testFailInvoked(tM)
		},
	},
}

func TestSimpleToBeTrue(t *testing.T) {
	var b strings.Builder
	tMock := mock.NewTMock()
	s := format.NewDefaultStyle(format.WithNoTrace())
	expect := texp.Expect(tMock, conf.OutputTo(&b), conf.StyleWith(s))

	expect(false).ToBeTrue()

	if c, ok := tMock.CallsTo("Fail"); ok && c != 1 {
		t.Error("Expecting call to Fail one time")
	}

	expLog := "Test: Test\nError: \n"
	if b.String() != expLog {
		t.Errorf("Expecting %s, got %s", expLog, b.String())
	}
}

func TestWithCustomParametrizedMessage(t *testing.T) {
	var b strings.Builder
	tMock := mock.NewTMock()
	s := format.NewDefaultStyle(format.WithNoTrace())
	expect := texp.Expect(tMock, conf.OutputTo(&b), conf.StyleWith(s))

	expect(false).ToBeTrue("expecting %v, got %v", true, false)

	expLog := "Test: Test\nError: expecting true, got false\n"
	if b.String() != expLog {
		t.Errorf("wrong custom error message with params:\n%s", b.String())
	}
}

// FAIL: TestFunctionName (0.00s)
// testfilename_test.go:##: log added to message
func TestErrorLogWithLinenumberAndFilename(t *testing.T) {
	var b strings.Builder
	tMock := mock.NewTMock()
	expect := texp.Expect(tMock, conf.OutputTo(&b))

	expect(false).ToBeTrue("")
	_, filename, linenum, _ := runtime.Caller(0)
	expLog := fmt.Sprintf("Test: Test\nTrace: %s:%v\nError: \n", path.Base(filename), linenum-1)
	if b.String() != expLog {
		t.Errorf("Expecting %s, got %s", expLog, b.String())
	}
}

func TestToBeTrueCases(t *testing.T) {
	for _, tt := range trueTests {
		t.Run(tt.N, func(t *testing.T) {
			tMock := mock.NewTMock()
			expect := texp.Expect(tMock, conf.OutputTo(new(strings.Builder)))
			tt.T(expect)
			if !tt.A(tMock) {
				t.Error("Wrong calls to Fail")
			}
		})
	}
}
