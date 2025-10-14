package main

import (
	"fmt"
	"os"

	"github.com/0x1eef/majortom/control"
)

func enable(ns, feature, path string) {
	if ctx, err := control.NewContext(control.Namespace(ns)); err != nil {
		fatalln("%s", err)
	} else {
		println("ok")
		ctx.Free()
	}
}

func disable(ns, feature, path string) {
	if ctx, err := control.NewContext(control.Namespace(ns)); err != nil {
		fatalln("%s", err)
	} else {
		println("ok")
		ctx.Free()
	}
}

func restore(ns, feature, path string) {
	if ctx, err := control.NewContext(control.Namespace(ns)); err != nil {
		fatalln("%s", err)
	} else {
		println("ok")
		ctx.Free()
	}
}

func query(ns, path string) {
	if ctx, err := control.NewContext(control.Namespace(ns)); err != nil {
		fatalln("%s", err)
	} else {
		if names, err := ctx.FeatureNames(); err != nil {
			fatalln("%s", err)
		} else {
			defer ctx.Free()
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
