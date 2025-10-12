package main

import (
	"log"
	"github.com/0x1eef/majortom/control"
)

func enable(ns, feature, path string) {
	ctx := control.New(control.Namespace(ns))
	if err := ctx.Enable(feature, path); err != nil {
		log.Fatalf("control: %s\n", err)
	} else {
		log.Printf("control: ok\n")
	}
}

func disable(ns, feature, path string) {
	ctx := control.New(control.Namespace(ns))
	if err := ctx.Disable(feature, path); err != nil {
		log.Fatalf("control: %s\n", err)
	} else {
		log.Printf("control: ok\n")
	}
}

func restore(ns, feature, path string) {
	ctx := control.New(control.Namespace(ns))
	if err := ctx.Sysdef(feature, path); err != nil {
		log.Fatalf("control: %s\n", err)
	} else {
		log.Printf("control: ok\n")
	}
}
