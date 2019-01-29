package texp

import (
	"fmt"

	"github.com/gekorob/texp/conf"
	"github.com/gekorob/texp/log"
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

// TestingT is useful to make assertion on the test runner passed as
// parameter and to mock failure methods for test reasons
type TestingT interface {
	Fail()
	FailNow()
	Name() string
}

type exp struct {
	t      TestingT
	config *conf.Config
	logQ   *log.Queue
	sample interface{}
	failF  func()
}

func (e *exp) T() TestingT {
	return e.t
}

func (e *exp) Config() *conf.Config {
	return e.config
}

func (e *exp) log(msg string) {
	e.logQ.Push(log.Message{
		Severity: log.Test,
		Content:  e.t.Name(),
	})
	e.logQ.Push(log.Message{
		Severity: log.Error,
		Content:  msg,
	})

	for m, ok := e.logQ.Front(); ok; m, ok = e.logQ.Next() {
		fmt.Fprint(e.config.Output(), m.Content)
	}
}

func (e *exp) fail() {
	e.failF()
}

// Expect returns a builder function to setup test expectations ala RSpec
func Expect(t TestingT, options ...func(*conf.Config)) func(interface{}) *exp {
	e := exp{
		t:      t,
		logQ:   log.NewQueue(),
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

// The ToBeTrue method match the true value of the sample
func (e *exp) ToBeTrue(msgs ...string) *exp {
	if e.sample != true {
		e.log("")
		e.fail()
	}
	return e
}
