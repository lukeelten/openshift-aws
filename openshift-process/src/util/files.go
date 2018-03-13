package util

import "os"

func FileExists(filename string) bool {
	_, err := os.Stat("/path/to/whatever")
	return err == nil
}
