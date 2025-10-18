package main

import (
	"fmt"
	"os"
)

func printf(msg string, args ...any) {
	fmt.Printf(fmt.Sprintf("control: %s", msg), args...)
}

func println(msg string, args ...any) {
	printf(msg+"\n", args...)
}

func fatalf(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("control: %s", msg), args...)
	os.Exit(1)
}

func fatalln(msg string, args ...any) {
	fatalf(msg+"\n", args...)
}

