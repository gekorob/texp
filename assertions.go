package texp

import (
	"reflect"

	"github.com/gekorob/texp/format"
)

// ToBeTrue method match the true value of the sample
func (e *exp) ToBeTrue(msgs ...interface{}) *exp {
	if e.sample != true {
		e.log(format.ToString(msgs...))
		e.fail()
	}
	return e
}

func (e *exp) ToEqual(expValue interface{}, msgs ...interface{}) *exp {
	// if !reflect.DeepEqual(e.sample, expValue) {
	eq := reflect.DeepEqual(e.sample, expValue)
	if !eq {
		e.log(format.ToString(msgs...))
		e.fail()
	}
	return e
}

// ToBeNil method returns true if the sample can be
// considered Nil
func (e *exp) ToBeNil(msgs ...interface{}) *exp {
	if !isNil(e.sample) {
		e.log(format.ToString(msgs...))
		e.fail()
	}
	return e
}

func isNil(o interface{}) bool {
	if o == nil {
		return true
	}

	v := reflect.ValueOf(o)
	if canUseIsNilByKind(v.Kind()) && v.IsNil() {
		return true
	}

	return false
}

func canUseIsNilByKind(k reflect.Kind) bool {
	if k >= reflect.Chan && k <= reflect.Slice {
		return true
	}
	return false
}
