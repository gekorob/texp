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
	{
		N: "Equal to nil",
		T: func(expect texp.ExpBuilder) {
			expect("first").ToEqual(nil)
		},
		A: func(tM *mock.TMock) bool {
			return testFailInvoked(tM)
		},
	},
	{
		N: "Not Equal uint32 and uint64",
		T: func(expect texp.ExpBuilder) {
			expect(uint32(3)).ToEqual(uint64(3))
		},
		A: func(tM *mock.TMock) bool {
			return testFailInvoked(tM)
		},
	},
	{
		N: "Equal uint",
		T: func(expect texp.ExpBuilder) {
			expect(uint(3)).ToEqual(uint(3))
		},
		A: func(tM *mock.TMock) bool {
			return !testFailInvoked(tM)
		},
	},
	{
		N: "Equal slicesi, same order",
		T: func(expect texp.ExpBuilder) {
			s1 := []string{"first", "second", "third"}
			s2 := []string{"first", "second", "third"}
			expect(s1).ToEqual(s2)
		},
		A: func(tM *mock.TMock) bool {
			return !testFailInvoked(tM)
		},
	},
	{
		N: "Not Equal slices, same elements different order",
		T: func(expect texp.ExpBuilder) {
			s1 := []string{"first", "second", "third"}
			s2 := []string{"second", "first", "third"}
			expect(s1).ToEqual(s2)
		},
		A: func(tM *mock.TMock) bool {
			return testFailInvoked(tM)
		},
	},
	{
		N: "Not Equal maps",
		T: func(expect texp.ExpBuilder) {
			m1 := map[string]interface{}{"element": 5}
			m2 := map[string]interface{}{"element": 20}
			expect(m1).ToEqual(m2)
		},
		A: func(tM *mock.TMock) bool {
			return testFailInvoked(tM)
		},
	},
	{
		N: "Equal nil and nil",
		T: func(expect texp.ExpBuilder) {
			expect(nil).ToEqual(nil)
		},
		A: func(tM *mock.TMock) bool {
			return !testFailInvoked(tM)
		},
	},
	{
		N: "Not Equal nil map",
		T: func(expect texp.ExpBuilder) {
			var m1 map[string]int
			m2 := map[string]int{}
			expect(m1).ToEqual(m2)
		},
		A: func(tM *mock.TMock) bool {
			return testFailInvoked(tM)
		},
	},
	{
		N: "Equal nil map",
		T: func(expect texp.ExpBuilder) {
			var m1 map[string]int
			var m2 map[string]int
			expect(m1).ToEqual(m2)
		},
		A: func(tM *mock.TMock) bool {
			return !testFailInvoked(tM)
		},
	},
	{
		N: "Equal nil slice",
		T: func(expect texp.ExpBuilder) {
			var s1 []string
			var s2 []string
			expect(s1).ToEqual(s2)
		},
		A: func(tM *mock.TMock) bool {
			return !testFailInvoked(tM)
		},
	},
	{
		N: "Not Equal chan",
		T: func(expect texp.ExpBuilder) {
			var c1 chan string
			c2 := make(chan string)
			expect(c1).ToEqual(c2)
		},
		A: func(tM *mock.TMock) bool {
			return testFailInvoked(tM)
		},
	},
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
