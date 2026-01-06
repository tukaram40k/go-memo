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

	err := internal.Execute(flags, params)
	if err != nil {
		fmt.Printf("runtime error: %v\n", err)
	}
}
