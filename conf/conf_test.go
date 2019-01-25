package conf_test

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/gekorob/texp/conf"
	"github.com/gekorob/texp/format"
	"github.com/gekorob/texp/log"
)

type StyleMock struct {
	// BySeverity func(log.Severity) *format.Style
}

func (s *StyleMock) BySeverity(sev log.Severity) *format.Style {
	return &format.Style{
		Color:         "\u001b[39m",
		Indentation:   "",
		Separator:     ": ",
		Termination:   "\n",
		SeverityLabel: "Severity",
	}
}

func TestConfigCreation(t *testing.T) {
	c := conf.NewConfig()

	if !reflect.DeepEqual(c.Output(), os.Stdout) {
		t.Error("wrong default output writer")
	}

	if !reflect.DeepEqual(c.Style(), format.NewDefaultStyle()) {
		t.Error("wrong default style")
	}
}

func TestConfigWithOutputOptions(t *testing.T) {
	var b strings.Builder
	c := conf.NewConfig(conf.OutputTo(&b))

	if !reflect.DeepEqual(c.Output(), &b) {
		t.Error("error setting output option")
	}
}

func TestConfigWithStyleOption(t *testing.T) {
	mStyle := new(StyleMock)
	c := conf.NewConfig(conf.StyleWith(mStyle))

	if !reflect.DeepEqual(c.Style(), mStyle) {
		t.Error("error setting style option")
	}
}

func TestConfigWithAllOptions(t *testing.T) {
	var b strings.Builder
	mStyle := new(StyleMock)
	c := conf.NewConfig(conf.StyleWith(mStyle), conf.OutputTo(&b))

	if !reflect.DeepEqual(c.Style(), mStyle) {
		t.Error("error setting style option")
	}

	if !reflect.DeepEqual(c.Output(), &b) {
		t.Error("error setting output option")
	}
}

func TestConfigFromAnotherConfig(t *testing.T) {
	var b strings.Builder

	oC := conf.NewConfig(conf.OutputTo(&b))
	c := conf.FromConfig(oC)

	if !reflect.DeepEqual(c, oC) {
		t.Error("wrong configuration copy")
	}
}
