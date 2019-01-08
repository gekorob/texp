package texp

import (
	"testing"

	"github.com/gekorob/texp/out"
)

var decorator out.Formatter

type exp struct {
	t     *testing.T
	s     interface{}
	failF func()
}

// Expect builder returns a function to setup test expectations ala RSpec
func Expect(t *testing.T) func(interface{}) *exp {
	e := &exp{
		t:     t,
		failF: t.Fail,
	}

	return func(s interface{}) *exp {
		e.s = s
		return e
	}
}

func (e *exp) T() *testing.T {
	return e.t
}

func (e *exp) logAndFail() bool {

	e.failF()
	return false
}

func (e *exp) ToBeTrue() bool {
	if e.s != true {
		return e.logAndFail()
	}
	return true
}
