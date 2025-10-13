package main

import (
	"fmt"
	"os"

	"github.com/0x1eef/majortom/control"
)

func enable(ns, feature, path string) {
	ctx := control.New(control.Namespace(ns))
	if err := ctx.Enable(feature, path); err != nil {
		fatalln("%s", err)
	} else {
		println("ok")
	}
}

func disable(ns, feature, path string) {
	ctx := control.New(control.Namespace(ns))
	if err := ctx.Disable(feature, path); err != nil {
		fatalln("%s", err)
	} else {
		println("ok")
	}
}

func restore(ns, feature, path string) {
	ctx := control.New(control.Namespace(ns))
	if err := ctx.Sysdef(feature, path); err != nil {
		fatalln("%s", err)
	} else {
		println("ok")
	}
}

func query(ns, path string) {
	ctx := control.New(control.Namespace(ns))
	if names, err := ctx.FeatureNames(); err != nil {
		fatalln("%s", err)
	} else {
		for i, name := range names {
			if status, err := ctx.Status(name, path); err != nil {
				fatalln("%s", err)
			} else {
				if i == 0 {
					println("%-25s %s", "Feature", "Status")
				}
				switch status {
				case "enabled":
					println("%-25s enabled", name)
				case "disabled":
					println("%-25s disabled", name)
				case "sysdef":
					println("%-25s system default", name)
				}
			}
		}
	}
}

func printf(msg string, args ...any) {
	fmt.Printf(fmt.Sprintf("control: %s", msg), args...)
}

func fatalf(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("control: %s", msg), args...)
	os.Exit(1)
}

func println(msg string, args ...any) {
	printf(msg+"\n", args...)
}

func fatalln(msg string, args ...any) {
	fatalf(msg+"\n", args...)
}
