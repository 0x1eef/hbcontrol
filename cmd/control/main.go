package main

import (
	"flag"
	"fmt"
	"github.com/0x1eef/majortom/control"
	"os"
)

const width = 25
const columns = 3

var args []string
var help *bool
var namespace *string

func main() {
	if *help || len(args) == 0 || len(args) > 3 {
		usage()
		os.Exit(2)
	} else {
		cmd, feature, path := args[0], args[1], args[2]
		switch cmd {
		case "enable":
			enable(*namespace, feature, path)
		case "disable":
			disable(*namespace, feature, path)
		case "restore":
			restore(*namespace, feature, path)
		default:
			usage()
			os.Exit(2)
		}
	}
}

func usage() {
	warn("usage: control [-hn] <command> <feature> <file>")
	warn("")
	warn("COMMANDS")
	warn("    enable      Enable a feature")
	warn("    disable     Disable a feature")
	warn("    restore     Restore a feature to the system default")
	warn("")
	warn("OPTIONS")
	warn("    -h          Show help")
	warn("    -n          Set the namespace (either 'user' or 'system')")
	warn("")
	warn("EXAMPLES")
	warn("    control enable mprotect /usr/bin/ls")
	warn("    control disable pageexec /usr/bin/ls")
	warn("    control restore segvguard /usr/bin/ls")
	warn("")

	warn("FEATURES")
	ctx := control.New(control.Namespace(*namespace))
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

func init() {
	help = flag.Bool("h", false, "Show help")
	namespace = flag.String("n", "system", "Set namespace (either user or system)")
	flag.Parse()
	args = flag.Args()
}
