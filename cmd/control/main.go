package main

import (
	"flag"
	"fmt"
	"os"
)

var args []string
var help *bool
var namespace *string

func main() {
	if *help || len(args) == 0 || len(args) > 3 {
		usage()
		os.Exit(1)
	}
	cmd, feature, path := args[0], args[1], args[2]
	switch cmd {
	case "enable":
		enable(*namespace, feature, path)
	case "disable":
		disable(*namespace, feature, path)
	case "restore":
		restore(*namespace, feature, path)
	}
}

func usage() {
	warn("usage: control [-hn] command file\n")
	warn("\nCOMMANDS\n")
	warn("    enable\tEnable a feature\n")
	warn("    disable\tDisable a feature\n")
	warn("    restore\tRestore a feature to the system default\n")
	warn("\nOPTIONS\n")
	warn("    -h\tShow help\n")
	warn("    -n\tSet the namespace (either user or system)\n")
	warn("\nEXAMPLES\n")
	warn("    control enable mprotect /usr/bin/ls\n")
	warn("    control disable pageexec /usr/bin/ls\n")
	warn("    control restore segvguard /usr/bin/ls\n")
}

func warn(message string) {
	fmt.Fprintf(os.Stderr, message)
}

func init() {
	help = flag.Bool("h", false, "Show help")
	namespace = flag.String("n", "system", "Set namespace (either user or system)")
	flag.Parse()
	args = flag.Args()
}
