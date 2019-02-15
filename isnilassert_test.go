package texp_test

import (
	"testing"

	"github.com/gekorob/texp"
	"github.com/gekorob/texp/mock"
)

var niltests = []struct {
	N string
	T func(texp.ExpBuilder)
	A func(*mock.TMock) bool
}{
	{
		N: "NilIsNil",
		T: func(expect texp.ExpBuilder) {
			expect(nil).ToBeNil()
		},
		A: func(tM *mock.TMock) bool {
			return !testFailInvoked(tM)
		},
	},
	{
		N: "EmptyStructPointerNil",
		T: func(expect texp.ExpBuilder) {
			type FakeObj struct {
			}
			expect((*struct{})(nil)).ToBeNil()
		},
		A: func(tM *mock.TMock) bool {
			return !testFailInvoked(tM)
		},
	},
	{
		N: "NewObjNotNil",
		T: func(expect texp.ExpBuilder) {
			type FakeObj struct {
			}
			expect(new(FakeObj)).ToBeNil()
		},
		A: func(tM *mock.TMock) bool {
			return testFailInvoked(tM)
		},
	},
}

func TestIsNil(t *testing.T) {
	for _, tt := range niltests {
		t.Run(tt.N, func(t *testing.T) {
			tMock := mock.NewTMock()
			expect := texp.Expect(tMock)
			tt.T(expect)
			if !tt.A(tMock) {
				t.Error("Wrong calls to Fail")
			}
		})
	}
}
