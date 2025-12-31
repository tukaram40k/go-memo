package main

import (
	"fmt"
	"os"

	"go-memo/internal"
)

func main() {
	fmt.Println("Args:", os.Args[1:])
	args := os.Args[1:]
	fmt.Println(internal.GetParams(args))
}
