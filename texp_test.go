package texp_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/gekorob/texp"
	"github.com/gekorob/texp/conf"
	"github.com/gekorob/texp/mock"
)

func TestExpectCreation(t *testing.T) {
	expect := texp.Expect(t)

	if expect == nil {
		t.Fatal("expectation creation error")
	}

	if expT := expect(nil).T(); t != expT {
		t.Error("wrong test runner")
	}

	c := conf.NewConfig()

	if !reflect.DeepEqual(texp.GlobalConfig(), c) {
		t.Error("wrong global config")
	}

	if !reflect.DeepEqual(expect(nil).Config(), texp.GlobalConfig()) {
		t.Error("wrong expectation instance config")
	}
}

func TestExpectCreationWithOptions(t *testing.T) {
	var b strings.Builder

	expect := texp.Expect(t, conf.OutputTo(&b))
	c := conf.NewConfig(conf.OutputTo(&b))

	if !reflect.DeepEqual(expect(nil).Config(), c) {
		t.Error("wrong instance config")
	}

	if reflect.DeepEqual(expect(nil).Config(), texp.GlobalConfig()) {
		t.Error("error, instance and global config must be different")
	}
}

func testFailInvoked(tM *mock.TMock) bool {
	if nC, ok := tM.CallsTo("Fail"); ok && nC != 0 {
		return true
	}
	return false
}
