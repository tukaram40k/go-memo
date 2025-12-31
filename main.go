package main

import (
	"fmt"
	"os"

	"go-memo/internal"
)

func main() {
	args := os.Args[1:]
	flags, params := internal.GetParams(args)
	mode := internal.GetMode(flags)

	fmt.Println("flags:", flags)
	fmt.Println("mode:", mode)
	fmt.Println("params:", params)
}
