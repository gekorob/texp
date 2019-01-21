package texp

import (
	"io"

	"github.com/gekorob/texp/format"
)

type config struct {
	out   io.Writer
	style format.Styler
}
