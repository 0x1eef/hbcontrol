package main

import (
	"os"
)

const (
	maxLinkDepth = 25
)

func isRegularFile(path string) bool {
	stat, err := os.Lstat(path)
	if err != nil {
		return false
	}
	mode := stat.Mode()
	return mode.IsRegular()
}

func follow(path string) string {
	if !options.followSymlinks {
		return path
	}
	if isRegularFile(path) {
		return path
	}

	p, err := os.Readlink(path)
	for i := 0; !isRegularFile(p); i++ {
		p, err = os.Readlink(p)
		if err != nil || i == maxLinkDepth {
			break
		}
	}

	if err != nil {
		return path
	} else {
		return p
	}
}
