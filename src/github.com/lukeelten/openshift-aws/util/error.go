package util

import (
	"fmt"
	"os"
)

func ExitOnError(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s \n", msg)
		panic(err)
	}
}
