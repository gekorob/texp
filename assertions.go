package texp

import "github.com/gekorob/texp/format"

// The ToBeTrue method match the true value of the sample
func (e *exp) ToBeTrue(msgs ...interface{}) *exp {
	if e.sample != true {
		e.log(format.ToString(msgs...))
		e.fail()
	}
	return e
}

func (e *exp) ToBeNil(msgs ...interface{}) *exp {
	if e.sample != nil {
		e.log(format.ToString(msgs...))
		e.fail()
	}
	return e
}
