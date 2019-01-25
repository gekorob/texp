package conf

import (
	"io"
	"os"

	"github.com/gekorob/texp/format"
)

// The Config structure contains the output writer and the style formatter
type Config struct {
	out   io.Writer
	style format.Styler
}

// NewConfig creates a config instance with default values
func NewConfig(options ...func(*Config)) *Config {
	c := Config{
		out:   os.Stdout,
		style: format.NewDefaultStyle(),
	}

	for _, option := range options {
		option(&c)
	}

	return &c
}

// FromConfig method creates a new Config based on the one passed
// as a parameter
func FromConfig(c *Config) *Config {
	return NewConfig(OutputTo(c.Output()), StyleWith(c.Style()))
}

// Output method retrieves the output writer set in the config object
func (c *Config) Output() io.Writer {
	return c.out
}

// Style method gets the style formatter set in the config object
func (c *Config) Style() format.Styler {
	return c.style
}

func (c *Config) setOutputTo(w io.Writer) {
	c.out = w
}

// OutputTo method allows to set the output writer
func OutputTo(w io.Writer) func(*Config) {
	return func(c *Config) {
		c.setOutputTo(w)
	}
}

func (c *Config) setStyleWith(s format.Styler) {
	c.style = s
}

// StyleWith method is for setting the style format
func StyleWith(s format.Styler) func(*Config) {
	return func(c *Config) {
		c.setStyleWith(s)
	}
}
