package main

import (
	"fmt"
	"os"

	"github.com/0x1eef/majortom/control"
)

func enable(ns, feature, path string) {
	ctx := control.New(control.Namespace(ns))
	if err := ctx.Enable(feature, path); err != nil {
		fatalf("%s\n", err)
	} else {
		printf("ok\n")
	}
}

func disable(ns, feature, path string) {
	ctx := control.New(control.Namespace(ns))
	if err := ctx.Disable(feature, path); err != nil {
		fatalf("%s\n", err)
	} else {
		printf("ok\n")
	}
}

func restore(ns, feature, path string) {
	ctx := control.New(control.Namespace(ns))
	if err := ctx.Sysdef(feature, path); err != nil {
		fatalf("%s\n", err)
	} else {
		printf("ok\n")
	}
}

func query(ns, path string) {
	ctx := control.New(control.Namespace(ns))
	if names, err := ctx.FeatureNames(); err != nil {
		fatalf("%s\n", err)
	} else {
		for i, name := range names {
			if status, err := ctx.Status(name, path); err != nil {
				fatalf("%s\n", err)
			} else {
				if i == 0 {
					printf("%-25s %s\n", "Feature", "Status")
				}
				switch status {
				case "enabled":
					printf("%-25s enabled\n", name)
				case "disabled":
					printf("%-25s disabled\n", name)
				case "sysdef":
					printf("%-25s system default\n", name)
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
	os.Exit(2)
}
