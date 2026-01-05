package internal

import (
	"fmt"
	"slices"
)

func Execute(flags []string, params []string) error {
	if slices.Contains(flags, "error") || len(flags) != 1 {
		return ProcessError(flags)
	}

	if flags[0] == "help" {
		return ProcessHelp()
	}

	if flags[0] == "version" {
		return ProcessVersion()
	}

	if flags[0] == "add" && len(params) > 0{
		return ProcessAdd(params[0])
	}

	if flags[0] == "list" {

	}

	if flags[0] == "remove" {

	}

	return fmt.Errorf("unknown error")
}
