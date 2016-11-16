package godeg

import (
	"fmt"
	"runtime/debug"
)

// Dprintf to debug print in red
func Dprintf(f string, args ...interface{}) {
	fmt.Printf("\033[48;5;1m")
	fmt.Printf(f, args...)
	fmt.Printf("\033[0m\n")
}

// DSprintf debug print with stack
func DSprintf(f string, args ...interface{}) {
	Dprintf(f, args...)
	debug.PrintStack()
}
