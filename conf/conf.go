package conf

import (
	"io"
	"os"

	"github.com/gekorob/texp/format"
)

type config struct {
	out   io.Writer
	style format.Styler
}

func NewConfig(options ...func(*config)) *config {
	c := config{
		out:   os.Stdout,
		style: format.NewDefaultStyle(),
	}

	for _, option := range options {
		option(&c)
	}

	return &c
}

func (c *config) Output() io.Writer {
	return c.out
}

func (c *config) Style() format.Styler {
	return c.style
}

func (c *config) setOutputTo(w io.Writer) {
	c.out = w
}

func OutputTo(w io.Writer) func(*config) {
	return func(c *config) {
		c.setOutputTo(w)
	}
}

func (c *config) setStyleWith(s format.Styler) {
	c.style = s
}

func StyleWith(s format.Styler) func(*config) {
	return func(c *config) {
		c.setStyleWith(s)
	}
}
