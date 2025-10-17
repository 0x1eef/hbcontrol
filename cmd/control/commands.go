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
		defer ctx.Free()
		println("ok")
	}
}

func disable(ns, feature, path string) {
	if ctx, err := control.NewContext(control.Namespace(ns)); err != nil {
		fatalln("%s", err)
	} else {
		defer ctx.Free()
		println("ok")
	}
}

func sysdef(ns, feature, path string) {
	if ctx, err := control.NewContext(control.Namespace(ns)); err != nil {
		fatalln("%s", err)
	} else {
		defer ctx.Free()
		println("ok")
	}
}

func query(ns, path string) {
	ctx, err := control.NewContext(control.Namespace(ns))
	if err != nil {
		fatalln("%s", err)
	}

	names, err := ctx.FeatureNames()
	if err != nil {
		fatalln("%s", err)
	}
	defer ctx.Free()

	for i, name := range names {
		status, err := ctx.Status(name, path)
		if err != nil {
			fatalln("%s", err)
		}
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

func println(msg string, args ...any) {
	printf(msg+"\n", args...)
}

func fatalln(msg string, args ...any) {
	fatalf(msg+"\n", args...)
}

func printf(msg string, args ...any) {
	fmt.Printf(fmt.Sprintf("control: %s", msg), args...)
}

func fatalf(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("control: %s", msg), args...)
	os.Exit(1)
}
