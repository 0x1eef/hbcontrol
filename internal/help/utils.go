package help

import (
	"fmt"
	"os"
)

func printf(msg string, args ...any) {
	fmt.Fprintf(os.Stdout, msg, args...)
}

func println(msg string, args ...any) {
	printf(msg+"\n", args...)
}

func warnf(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg, args...)
}

func warnln(msg string, args ...any) {
	warnf(msg+"\n", args...)
}