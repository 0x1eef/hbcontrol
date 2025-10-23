package main

import (
	"os"
)

func isRegularFile(path string) bool {
	stat, err := os.Lstat(path)
	if err != nil {
		return false
	} else {
		mode := stat.Mode()
		return mode.IsRegular()
	}
}

func follow(path string) string {
	if !options.followSymlinks {
		return path
	} else if isRegularFile(path) {
		return path
	} else {
		const max = 20
		p, err := os.Readlink(path)
		for i := 0; !isRegularFile(p); i++ {
			p, err = os.Readlink(p)
			if err != nil || i == max {
				break
			}
		}
		if err != nil {
			return path
		} else {
			return p
		}
	}
}
