package main

import (
	"flag"
	"os"

	"git.hardenedbsd.org/0x1eef/ctrl"
	"git.hardenedbsd.org/0x1eef/ctrl/internal/help"
)

type flags struct {
	help           bool
	version        bool
	followSymlinks bool
	namespace      string
}

var args []string
var options flags

func main() {
	args = flag.Args()
	if options.help {
		showHelp(0)
	} else if options.version {
		printf("v%s\n", ctrl.Version)
	} else if len(args) == 0 {
		showHelp(0)
	} else if len(args) == 2 {
		cmd, path := args[0], args[1]
		switch cmd {
		case "query", "status":
			query(options.namespace, follow(path))
		default:
			showHelp(1)
		}
	} else if len(args) == 3 {
		cmd, feature, path := args[0], args[1], args[2]
		switch cmd {
		case "enable":
			enable(options.namespace, feature, follow(path))
		case "disable":
			disable(options.namespace, feature, follow(path))
		case "sysdef":
			sysdef(options.namespace, feature, follow(path))
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
	help.PrintFeatures(options.namespace)
	os.Exit(code)
}

func init() {
	flag.BoolVar(&options.help, "h", false, "Show help")
	flag.BoolVar(&options.version, "v", false, "Print the current version")
	flag.BoolVar(&options.followSymlinks, "H", false, "Follow symlinks")
	flag.StringVar(&options.namespace, "n", "system", "Set namespace (either user or system)")
	flag.Parse()
}
