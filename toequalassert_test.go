package texp_test

import (
	"testing"

	"github.com/gekorob/texp"
	"github.com/gekorob/texp/mock"
)

var eqTests = []struct {
	N string
	T func(texp.ExpBuilder)
	A func(*mock.TMock) bool
}{
	{
		N: "Equal ints",
		T: func(expect texp.ExpBuilder) {
			expect(1).ToEqual(1)
		},
		A: func(tM *mock.TMock) bool {
			return !testFailInvoked(tM)
		},
	},
	{
		N: "Not Equal ints",
		T: func(expect texp.ExpBuilder) {
			expect(10).ToEqual(1)
		},
		A: func(tM *mock.TMock) bool {
			return testFailInvoked(tM)
		},
	},
	{
		N: "Equal strings",
		T: func(expect texp.ExpBuilder) {
			expect("alpha").ToEqual("alpha")
		},
		A: func(tM *mock.TMock) bool {
			return !testFailInvoked(tM)
		},
	},
	{
		N: "Not Equal strings",
		T: func(expect texp.ExpBuilder) {
			expect("beta").ToEqual("alpha")
		},
		A: func(tM *mock.TMock) bool {
			return testFailInvoked(tM)
		},
	},
	{
		N: "Equal booleans",
		T: func(expect texp.ExpBuilder) {
			expect(true).ToEqual(true)
		},
		A: func(tM *mock.TMock) bool {
			return !testFailInvoked(tM)
		},
	},
	{
		N: "Not Equal booleans",
		T: func(expect texp.ExpBuilder) {
			expect(false).ToEqual(true)
		},
		A: func(tM *mock.TMock) bool {
			return testFailInvoked(tM)
		},
	},
	// {
	// 	N: "NewObjNotNil",
	// 	T: func(expect texp.ExpBuilder) {
	// 		type FakeObj struct {
	// 		}
	// 		expect(new(FakeObj)).ToBeNil()
	// 	},
	// 	A: func(tM *mock.TMock) bool {
	// 		return testFailInvoked(tM)
	// 	},
	// },
}

func TestEqual(t *testing.T) {
	for _, tt := range eqTests {
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
