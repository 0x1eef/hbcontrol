package main

import (
	"flag"
	"os"

	"github.com/0x1eef/control/internal/usage"
)

var args []string
var help *bool
var namespace *string

func main() {
	if *help || len(args) == 0 {
		showUsage(*namespace)
		os.Exit(2)
	} else if len(args) == 2 {
		cmd, path := args[0], args[1]
		switch cmd {
		case "status":
			status(*namespace, path)
		default:
			showUsage(*namespace)
		}
	} else if len(args) == 3 {
		cmd, feature, path := args[0], args[1], args[2]
		switch cmd {
		case "enable":
			enable(*namespace, feature, path)
		case "disable":
			disable(*namespace, feature, path)
		case "restore":
			restore(*namespace, feature, path)
		default:
			showUsage(*namespace)
			os.Exit(2)
		}
	}
}

func init() {
	help = flag.Bool("h", false, "Show help")
	namespace = flag.String("n", "system", "Set namespace (either user or system)")
	flag.Parse()
	args = flag.Args()
}

func showUsage(ns string) {
	usage.PrintHeader()
	usage.PrintCommands()
	usage.PrintOptions()
	usage.PrintExamples()
	usage.PrintFeatures(ns)
}
