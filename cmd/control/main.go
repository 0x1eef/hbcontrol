package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/0x1eef/control/internal/help"
)

type flags struct {
	help      *bool
	version   *bool
	namespace *string
}

const VERSION = "0.1.1"

var args []string
var options flags

func main() {
	flag.Parse()
	args = flag.Args()

	if *options.help {
		showHelp(0)
	} else if *options.version {
		fmt.Printf("v%s\n", VERSION)
	} else if len(args) == 0 {
		showHelp(0)
	} else if len(args) == 2 {
		cmd, path := args[0], args[1]
		switch cmd {
		case "query":
			query(*options.namespace, path)
		default:
			showHelp(1)
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
			showHelp(1)
		}
	} else {
		showHelp(1)
	}
}

func showHelp(code int) {
	help.PrintHeader()
	help.PrintCommands()
	help.PrintOptions()
	help.PrintExamples()
	help.PrintFeatures(*options.namespace)
	os.Exit(code)
}

func init() {
	options.help = flag.Bool("h", false, "Show help")
	options.version = flag.Bool("v", false, "Print the current version")
	options.namespace = flag.String("n", "system", "Set namespace (either user or system)")
}
