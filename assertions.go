package texp

import (
	"reflect"
)

// Not method is useful to express negative assertions
func (e *exp) Not() *exp {
	e.modF = negationModifierFunc
	return e
}

// ToBeTrue method match the true value of the sample
func (e *exp) ToBeTrue(msgs ...interface{}) *exp {
	if !e.modF(isTrue)(e.sample) {
		e.logAndFail(msgs...)
	}
	e.modF = neutralModifierFunc
	return e
}

// ToEqual method test the equality between sample and expectedValue
func (e *exp) ToEqual(expValue interface{}, msgs ...interface{}) *exp {
	if !reflect.DeepEqual(e.sample, expValue) {
		e.logAndFail(msgs...)
	}
	return e
}

// ToBeNil method returns true if the sample can be
// considered Nil
func (e *exp) ToBeNil(msgs ...interface{}) *exp {
	if !e.modF(isNil)(e.sample) {
		e.logAndFail(msgs...)
	}
	e.modF = neutralModifierFunc
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

func isTrue(o interface{}) bool {
	return o == true
}

func canUseIsNilByKind(k reflect.Kind) bool {
	if k >= reflect.Chan && k <= reflect.Slice {
		return true
	}
	return false
}
