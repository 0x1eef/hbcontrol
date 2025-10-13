package usage

import (
	"fmt"
	"os"

	"github.com/0x1eef/majortom/control"
)

const width = 25
const columns = 3

func PrintHeader() {
	warnlnf("usage: control [-hn] <command> <feature> <file>")
	warnlnf("")
}

func PrintCommands() {
	warnlnf("COMMANDS")
	warnlnf("    enable      Enable a feature")
	warnlnf("    disable     Disable a feature")
	warnlnf("    restore     Restore a feature to the system default")
	warnlnf("")
}

func PrintOptions() {
	warnlnf("OPTIONS")
	warnlnf("    -h          Show help")
	warnlnf("    -n          Set the namespace (either 'user' or 'system')")
	warnlnf("")
}

func PrintExamples() {
	warnlnf("EXAMPLES")
	warnlnf("    control enable mprotect /usr/bin/ls")
	warnlnf("    control disable pageexec /usr/bin/ls")
	warnlnf("    control restore segvguard /usr/bin/ls")
	warnlnf("")
}

func PrintFeatures(ns string) {
	warnlnf("FEATURES")
	ctx := control.New(control.Namespace(ns))
	if names, err := ctx.FeatureNames(); err != nil {
		warnlnf("There was an unexpected error")
	} else {
		for i, name := range names {
			warnf("%-*s", width, name)
			if (i+1)%columns == 0 || i == len(names)-1 {
				fmt.Println()
			}
		}
	}
}

func warnf(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg, args...)
}

func warnlnf(msg string, args ...any) {
	warnf(msg + "\n", args...)
}
