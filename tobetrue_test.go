package texp_test

import (
	"strings"
	"testing"

	"github.com/gekorob/texp"
	"github.com/gekorob/texp/conf"
)

func TestSimpleToBeTrue(t *testing.T) {
	var b strings.Builder
	expect := texp.Expect(t, conf.OutputTo(&b))

	// t.Error("error")
	expect(true).ToBeTrue()
	// log.Print(b.String())
}
