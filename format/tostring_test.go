package format_test

import (
	"testing"

	"github.com/gekorob/texp/format"
)

func TestToStringWithOneString(t *testing.T) {
	expStr := "simple string"

	got := format.ToString(expStr)

	if got != expStr {
		t.Errorf("wrong string. Expecting %s, got %s", expStr, got)
	}
}

func TestToStringWithNStrings(t *testing.T) {
	expStr := "firstsecondthird"

	got := format.ToString("first%s%s", "second", "third")

	if got != expStr {
		t.Errorf("wrong string. Expecting %s, got %s", expStr, got)
	}
}

func TestToStringWithNoString(t *testing.T) {
	expStr := ""

	got := format.ToString()

	if got != expStr {
		t.Errorf("wrong string. Expecting %s, got %s", expStr, got)
	}
}

func TestToStringWithParams(t *testing.T) {
	p := struct {
		a int
	}{a: 5}
	expStr := "{a:5}"

	got := format.ToString(p)

	if got != expStr {
		t.Errorf("wrong string. Expecting %s, got %s", expStr, got)
	}
}
