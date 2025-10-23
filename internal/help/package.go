package help

import (
	"os"

	"github.com/0x1eef/majortom/control"
)

const (
	width   = 25
	columns = 3
)

func PrintHeader() {
	println("usage: hbcontrol [-hnv] <command> [feature] <file>")
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
	println("    hbcontrol enable mprotect /bin/ls")
	println("    hbcontrol disable pageexec /bin/ls")
	println("    hbcontrol sysdef segvguard /bin/ls")
	println("    hbcontrol query /bin/ls")
	println("    hbcontrol status /bin/ls")
	println("")
}

func PrintFeatures(ns string) {
	println("FEATURES")
	ctx, err := control.NewContext(control.Namespace(ns))
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
		printf("%-*s", width, name)
		if (i+1)%columns == 0 || i == len(names)-1 {
			println("")
		}
	}
}
