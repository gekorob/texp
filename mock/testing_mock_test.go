package mock_test

import (
	"testing"

	"github.com/gekorob/texp/mock"
)

func TestMockOnExistingMethod(t *testing.T) {
	tMock := mock.NewTMock()

	tMock.Fail()
	tMock.FailNow()
	tMock.Fail()
	tMock.Name()

	if n, _ := tMock.CallsTo("Fail"); n != 2 {
		t.Errorf("wrong number of calls to Fail. Expected 2, got %v", n)
	}

	if n, _ := tMock.CallsTo("FailNow"); n != 1 {
		t.Errorf("wrong number of calls to FailNow. Expected 1, got %v", n)
	}

	if n, _ := tMock.CallsTo("Name"); n != 1 {
		t.Errorf("wrong number of calls to Name. Expected 1, got %v", n)
	}
}

func TestMockOnNotExistingMethod(t *testing.T) {
	tMock := mock.NewTMock()

	if _, ok := tMock.CallsTo("NotExistingMethod"); ok {
		t.Errorf("NotExisting Method call requests. Expecting ok=false, got ok=%v", ok)
	}
}
