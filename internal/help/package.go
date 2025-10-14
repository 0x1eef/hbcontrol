package help

import (
	"fmt"
	"os"

	"github.com/0x1eef/majortom/control"
)

const width = 25
const columns = 3

func PrintHeader() {
	warnln("usage: control [-hnv] <command> [feature] <file>")
	warnln("")
}

func PrintCommands() {
	warnln("COMMANDS")
	warnln("    enable      Enable a feature")
	warnln("    disable     Disable a feature")
	warnln("    restore     Restore a feature to the system default")
	warnln("    query       Query feature states")
	warnln("")
}

func PrintOptions() {
	warnln("OPTIONS")
	warnln("    -h          Show help")
	warnln("    -n          Set the namespace (either 'user' or 'system')")
	warnln("    -v          Print current version")
	warnln("")
}

func PrintExamples() {
	warnln("EXAMPLES")
	warnln("    control enable mprotect /bin/ls")
	warnln("    control disable pageexec /bin/ls")
	warnln("    control restore segvguard /bin/ls")
	warnln("    control query /bin/ls")
	warnln("")
}

func PrintFeatures(ns string) {
	warnln("FEATURES")
	if ctx, err := control.NewContext(control.Namespace(ns)); err != nil {
		warnln("%s", err)
	} else {
		defer ctx.Free()
		if names, err := ctx.FeatureNames(); err != nil {
			warnln("There was an unexpected error")
		} else {
			for i, name := range names {
				warnf("%-*s", width, name)
				if (i+1)%columns == 0 || i == len(names)-1 {
					fmt.Println()
				}
			}
		}
	}
}

func warnf(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg, args...)
}

func warnln(msg string, args ...any) {
	warnf(msg+"\n", args...)
}
