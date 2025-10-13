package main

import (
	"flag"
	"os"

	"github.com/0x1eef/control/internal/help"
)

type flags struct {
	help      *bool
	namespace *string
}

var args []string
var options flags

func main() {
	if *options.help || len(args) == 0 {
		showHelp()
		os.Exit(1)
	} else if len(args) == 2 {
		cmd, path := args[0], args[1]
		switch cmd {
		case "query":
			query(*options.namespace, path)
		default:
			showHelp()
		}
	} else if len(args) == 3 {
		cmd, feature, path := args[0], args[1], args[2]
		switch cmd {
		case "enable":
			enable(*options.namespace, feature, path)
		case "disable":
			disable(*options.namespace, feature, path)
		case "restore":
			restore(*options.namespace, feature, path)
		default:
			showHelp()
			os.Exit(1)
		}
	}
}

func showHelp() {
	help.PrintHeader()
	help.PrintCommands()
	help.PrintOptions()
	help.PrintExamples()
	help.PrintFeatures(*options.namespace)
}

func init() {
	options.help = flag.Bool("h", false, "Show help")
	options.namespace = flag.String("n", "system", "Set namespace (either user or system)")
	flag.Parse()
	args = flag.Args()
}
