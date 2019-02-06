package format

import (
	"fmt"
)

// ToString method tries to convert a list of params into a string,
// allowing the user to use the first param as a format string and the others
// as arguments
func ToString(msgs ...interface{}) string {
	if len(msgs) == 0 {
		return ""
	}
	m, ok := msgs[0].(string)
	if !ok {
		m = fmt.Sprintf("%+v", msgs[0])
	}

	return fmt.Sprintf(m, msgs[1:]...)
}
