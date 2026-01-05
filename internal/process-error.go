package internal

import "fmt"

func ProcessError(flags []string) error {
	fmt.Println("Unknown arguments")
	return ProcessHelp()
}
