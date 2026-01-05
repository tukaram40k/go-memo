package main

import (
	"fmt"
	"os"

	"go-memo/internal"
)

func main() {
	args := os.Args[1:]
	f, params := internal.GetParams(args)
	flags := internal.GetFlags(f)

	fmt.Printf("flags: %q\n", f)
	fmt.Printf("mode: %q\n", flags)
	fmt.Printf("params: %q\n\n", params)

	err := internal.Execute(flags, params)
	if err != nil {
		fmt.Printf("runtime error: %v\n", err)
	}
}
