package main

import (
	"fmt"
	"os"

	// "go-memo/internal/core"
	internal "go-memo/internal/input-handling"
)

func main() {
	args := os.Args[1:]
	f, params := internal.GetParams(args)
	flags := internal.GetFlags(f)

	// need to pass all flags/args to mode processors
	// also need to extract validation of flag composition

	fmt.Printf("flags: %q\n", f)
	fmt.Printf("mode: %q\n", flags)
	fmt.Printf("params: %q\n", params)
}
