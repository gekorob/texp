package texp

import (
	"testing"

	"github.com/gekorob/texp/conf"
)

// TODO: refactor to a config structure...
var gConfig *conf.Config

// TODO: need to figure out how to set this global values with
// flags or files
func init() {
	gConfig = conf.NewConfig()
}

// GlobalConfig returns the config object set at package level
func GlobalConfig() *conf.Config {
	return gConfig
}

type exp struct {
	t      *testing.T
	config *conf.Config
	sample interface{}
	failF  func()
}

func (e *exp) log() {
}

func (e *exp) fail() {
	e.failF()
}

func (e *exp) T() *testing.T {
	return e.t
}

func (e *exp) Config() *conf.Config {
	return e.config
}

// Expect returns a builder function to setup test expectations ala RSpec
func Expect(t *testing.T, options ...func(*conf.Config)) func(interface{}) *exp {
	e := exp{
		t:      t,
		config: conf.FromConfig(GlobalConfig()),
		failF:  t.Fail,
	}

	for _, option := range options {
		option(e.config)
	}

	return func(s interface{}) *exp {
		e.sample = s
		return &e
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
