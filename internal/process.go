package internal

import "slices"

func Execute(flags []string, params []string) {
	if slices.Contains(flags, "error") || len(flags) > 1 {
		ProcessError(flags)
	}

	if flags[0] == "help" {
		ProcessHelp()
	}

	if flags[0] == "version" {
		ProcessVersion()
	}

	if flags[0] == "add" && len(params) > 0{
		ProcessAdd(params[0])
	}

	if flags[0] == "list" {

	}

	if flags[0] == "remove" {

	}
}
