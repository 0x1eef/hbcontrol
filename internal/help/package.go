package help

import (
	"os"

	"git.hardenedbsd.org/0x1eef/ctrl"
)

func PrintHeader() {
	println("usage: ctrl [-hHnv] <command> [feature] <file>")
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
	println("    -H          Follow symlinks")
	println("    -n          Set the namespace (either 'user' or 'system')")
	println("    -v          Print current version")
	println("")
}

func PrintExamples() {
	println("EXAMPLES")
	println("    ctrl enable mprotect /bin/ls")
	println("    ctrl disable pageexec /bin/ls")
	println("    ctrl sysdef segvguard /bin/ls")
	println("    ctrl query /bin/ls")
	println("    ctrl status /bin/ls")
	println("")
}

func PrintFeatures(ns string) {
	println("FEATURES")
	ctx, err := ctrl.NewContext(ctrl.Namespace(ns))
	if err != nil {
		warnln("%s", err)
		os.Exit(1)
	}
	defer ctx.Free()

	names, err := ctx.FeatureNames()
	if err != nil {
		warnln("There was an unexpected error")
		os.Exit(1)
	}

	for i, name := range names {
		printf("%-*s", 25, name)
		if (i+1)%3 == 0 || i == len(names)-1 {
			println("")
		}
	}
}
