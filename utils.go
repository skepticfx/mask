package main

import (
	"fmt"
	"os"
)

// bail with an error message and return exit code = 1
func bail(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
