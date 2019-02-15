package texp

import (
	"fmt"
	"io"
	"path"
	"runtime"

	"github.com/gekorob/texp/conf"
	"github.com/gekorob/texp/format"
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
	e.logQ.Push(log.Test(e.t.Name()))
	traceCalls(e.logQ)
	e.logQ.Push(log.Error(msg))
	write(e.config.Output(), e.logQ, e.config.Style())
}

func (e *exp) fail() {
	e.failF()
}

// ExpBuilder is the function that returns the expectation
// object to call assertions
type ExpBuilder func(interface{}) *exp

// Expect returns a builder function to setup test expectations ala RSpec
func Expect(t TestingT, options ...func(*conf.Config)) ExpBuilder {
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

func traceCalls(q log.Queuer) {
	type call struct {
		file    string
		linenum int
	}

	c := new(call)

	for i := 0; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok || file == "<autogenerated>" {
			break
		}

		f := runtime.FuncForPC(pc)
		if f == nil {
			break
		}
		name := f.Name()

		if name == "testing.tRunner" {
			break
		}

		c.file = file
		c.linenum = line
	}
	q.Push(log.Trace(fmt.Sprintf("%s:%d", path.Base(c.file), c.linenum)))
}

func write(w io.Writer, logIter log.FwIterator, s format.Styler) {
	for m, ok := logIter.Front(); ok; m, ok = logIter.Next() {
		style := s.BySeverity(m.Severity)
		if style != nil {
			fmt.Fprint(w, style.SeverityLabel, style.Separator, m.Content, style.Termination)
		}
	}
}
