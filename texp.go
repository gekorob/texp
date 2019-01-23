package texp

import (
	"io"
	"os"
	"testing"

	"github.com/gekorob/texp/format"
)

// TODO: refactor to a config structure...
var (
	gOut   io.Writer
	gStyle format.Styler
)

// TODO: need to figure out how to set this global values with
// flags or files
func init() {
	gOut = os.Stdout
	gStyle = format.NewDefaultStyle()
}

type exp struct {
	t      *testing.T
	sample interface{}
	out    io.Writer
	style  format.Styler

	failF func()
}

func (e *exp) setOut(w io.Writer) {
	e.out = w
}

func (e *exp) setStyle(s format.Styler) {
	e.style = s
}

func (e *exp) log() {
}

func (e *exp) fail() {
	e.failF()
}

func (e *exp) T() *testing.T {
	return e.t
}

type expBuilder func(interface{}) *exp

func (b expBuilder) DefaultOutput() io.Writer {
	return gOut
}

// Expect returns a builder function to setup test expectations ala RSpec
func Expect(t *testing.T, options ...func(*exp)) expBuilder {
	e := exp{
		t:     t,
		out:   gOut,
		style: gStyle,
		failF: t.Fail,
	}

	for _, option := range options {
		option(&e)
	}

	return func(s interface{}) *exp {
		e.sample = s
		return &e
	}
}

// OutTo configuration function sets the output writer for the expectation instance.
func OutTo(w io.Writer) func(*exp) {
	return func(e *exp) {
		e.setOut(w)
	}
}

// Style sets the styler for the expectation instance.
func Style(s format.Styler) func(*exp) {
	return func(e *exp) {
		e.setStyle(s)
	}
}

// TODO: can be useful to return the exp type to invoke other methods
// on it, like a chain

// The ToBeTrue method match the true value of the sample
func (e *exp) ToBeTrue() *exp {
	if e.sample != true {
		e.log()
		e.fail()
	}
	return e
}
