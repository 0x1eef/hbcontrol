package main

import (
	"errors"

	"github.com/0x1eef/majortom/control"
)

var (
	errIrregularFile = errors.New("not a regular file")
)

func enable(ns, feature, path string) {
	if ctx, err := control.NewContext(control.Namespace(ns)); err != nil {
		fatalln("%s", err)
	} else {
		defer ctx.Free()
		if !isRegularFile(path) {
			fatalln("%s", errIrregularFile)
		}
		if err := ctx.Enable(feature, path); err != nil {
			fatalln("%s", err)
		}
		println("ok")
	}
}

func disable(ns, feature, path string) {
	if ctx, err := control.NewContext(control.Namespace(ns)); err != nil {
		fatalln("%s", err)
	} else {
		defer ctx.Free()
		if !isRegularFile(path) {
			fatalln("%s", errIrregularFile)
		}
		if err := ctx.Disable(feature, path); err != nil {
			fatalln("%s", err)
		}
		println("ok")
	}
}

func sysdef(ns, feature, path string) {
	if ctx, err := control.NewContext(control.Namespace(ns)); err != nil {
		fatalln("%s", err)
	} else {
		defer ctx.Free()
		if !isRegularFile(path) {
			fatalln("%s", errIrregularFile)
		}
		if err := ctx.Sysdef(feature, path); err != nil {
			fatalln("%s", err)
		}
		println("ok")
	}
}

func query(ns, path string) {
	ctx, err := control.NewContext(control.Namespace(ns))
	if err != nil {
		fatalln("%s", err)
	} else if !isRegularFile(path) {
		fatalln("%s", errIrregularFile)
	}
	defer ctx.Free()

	names, err := ctx.FeatureNames()
	if err != nil {
		fatalln("%s", err)
	}

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
