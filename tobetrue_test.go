package texp_test

import (
	"testing"

	"github.com/gekorob/texp"
)

func TestSimpleToBeTrue(t *testing.T) {
	expect := texp.Expect(t)

	expect(true).ToBeTrue()
}
