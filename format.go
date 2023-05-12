package corelib

import (
	"fmt"
)

func Format(input string) string {
	return fmt.Sprintf("%s: %s, %s", pkg, version, input)
}
