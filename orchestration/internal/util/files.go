package util

import (
	"os"
	"golang.org/x/sys/unix"
	)

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func IsWritable(path string) bool {
	return unix.Access(path, unix.W_OK) == nil
}