package corelib

import (
	"fmt"
)

func Format(input string) string {
	return fmt.Sprintf("%s@v%s: %s", pkg, version, input)
}
