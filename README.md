# TEXP

[![Go Report Card](https://goreportcard.com/badge/github.com/gekorob/texp)](https://goreportcard.com/report/github.com/gekorob/texp) [![GoDoc](https://godoc.org/github.com/gekorob/texp?status.png)](http://godoc.org/github.com/gekorob/texp) [![Build Status](https://travis-ci.org/gekorob/texp.svg?branch=master)](https://travis-ci.org/gekorob/texp)

## Testing by Expectations simple mini-framework

Easy expectations based testing mini-framework inspired by the wonderful Mat Ryer ["is"](https://github.com/matryer/is) project, but for RSpec nostalgic like me.
This project is a working progress, so expect changes and more new things.

### Basic usage

Using Texp with default settings is very simple

```go
import (
  ...
  "testing"

  "github.com/gekorob/texp"
  ...
)

func TestSomething(t *testing.T) {
  expect := texp.Expect(t)

  expect(2 == 2).ToBeTrue()
}

func TestSomethingOther(t *testing.T) {
  expect := texp.Expect(t)

  expect(3).ToEqual(2)
}
```

### Instance configuration

With Texp you can change the default configuration for each test, setting an output different from the StdOut or an alternative style (implementing the appropriate interface)

```go
import (
  ...
  "strings"
  "testing"

  "github.com/gekorob/texp"
  "github.com/gekorob/texp/conf"
  ...
)

func TestSomething(t *testing.T) {
  var b strings.Builder
  yourStyle := ....

  expect := texp.Expect(t, conf.OutputTo(&b, conf.StyleWith(yourStyle))

  expect(1 == 1).ToBeTrue()
}
```
