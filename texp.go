package texp

import (
	"testing"
)

var conf config

// var (
// 	out   io.Writer
// 	style format.Styler
// )

// Don't like it very much... prefer to implement a config obj...
// func init() {
// 	out = os.Stdout
// 	style = format.NewDefaultStyle()
// }

type exp struct {
	t *testing.T
	s interface{}

	failF func()
}

type expBuilder func(interface{}) *exp

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

func (e *exp) log() {
}

func (e *exp) fail() {
	e.failF()
}

// TODO: can be useful to return the exp type to invoke other methods
// on it, like a chain

// The ToBeTrue method match the true value of the sample
func (e *exp) ToBeTrue() *exp {
	if e.s != true {
		e.log()
		e.fail()
	}
	return e
}
