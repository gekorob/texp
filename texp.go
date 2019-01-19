package texp

import (
	"testing"
)

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

func (e *exp) logAndFail() {
	e.failF()
}

// TODO: can be useful to return the exp type to invoke other methods
// on it, like a chain

// The ToBeTrue method match the true value of the sample
func (e *exp) ToBeTrue() *exp {
	if e.s != true {
		e.logAndFail()
	}
	return e
}
