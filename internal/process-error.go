package internal

import (
	"fmt"
	"os"
)

func ProcessError(flags []string) {
	fmt.Println("Unknown arguments")
	ProcessHelp()
	os.Exit(0)
}
