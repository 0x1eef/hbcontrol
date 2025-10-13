package usage

import (
	"fmt"
	"os"

	"github.com/0x1eef/majortom/control"
)

const width = 25
const columns = 3

func PrintHeader() {
	warn("usage: control [-hn] <command> <feature> <file>")
	warn("")
}

func PrintCommands() {
	warn("COMMANDS")
	warn("    enable      Enable a feature")
	warn("    disable     Disable a feature")
	warn("    restore     Restore a feature to the system default")
	warn("")
}

func PrintOptions() {
	warn("OPTIONS")
	warn("    -h          Show help")
	warn("    -n          Set the namespace (either 'user' or 'system')")
	warn("")
}

func PrintExamples() {
	warn("EXAMPLES")
	warn("    control enable mprotect /usr/bin/ls")
	warn("    control disable pageexec /usr/bin/ls")
	warn("    control restore segvguard /usr/bin/ls")
	warn("")
}

func PrintFeatures(ns string) {
	warn("FEATURES")
	ctx := control.New(control.Namespace(ns))
	if names, err := ctx.FeatureNames(); err != nil {
		warn("There was an unexpected error")
	} else {
		for i, name := range names {
			fmt.Printf("%-*s", width, name)
			if (i+1)%columns == 0 || i == len(names)-1 {
				fmt.Println()
			}
		}
	}
}

func warn(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}
