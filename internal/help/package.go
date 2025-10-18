package help

import (
	"github.com/0x1eef/majortom/control"
)

const (
	width   = 25
	columns = 3
)

func PrintHeader() {
	println("usage: control [-hnv] <command> [feature] <file>")
	println("")
}

func PrintCommands() {
	println("COMMANDS")
	println("    enable            Enable a feature")
	println("    disable           Disable a feature")
	println("    sysdef            Restore the system default")
	println("    [query|status]    Query feature states")
	println("")
}

func PrintOptions() {
	println("OPTIONS")
	println("    -h          Show help")
	println("    -n          Set the namespace (either 'user' or 'system')")
	println("    -v          Print current version")
	println("")
}

func PrintExamples() {
	println("EXAMPLES")
	println("    control enable mprotect /bin/ls")
	println("    control disable pageexec /bin/ls")
	println("    control sysdef segvguard /bin/ls")
	println("    control query /bin/ls")
	println("    control status /bin/ls")
	println("")
}

func PrintFeatures(ns string) {
	println("FEATURES")
	if ctx, err := control.NewContext(control.Namespace(ns)); err != nil {
		warnln("%s", err)
	} else {
		defer ctx.Free()
		if names, err := ctx.FeatureNames(); err != nil {
			warnln("There was an unexpected error")
		} else {
			for i, name := range names {
				printf("%-*s", width, name)
				if (i+1)%columns == 0 || i == len(names)-1 {
					println("")
				}
			}
		}
	}
}
