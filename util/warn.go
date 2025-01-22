package util

import (
	"fmt"
	"os"
)

func Warn(format string, arg ...any) {
	fmt.Fprintf(os.Stderr, format, arg...)
}
