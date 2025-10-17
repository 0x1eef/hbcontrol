package main

import (
	"fmt"
	"os"
)

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
