package internal

import "slices"

func Execute(flags []string, params []string) {
	if slices.Contains(flags, "error") || len(flags) > 1 {
		ProcessError(flags)
	}

	if flags[0] == "help" {

	}

	if flags[0] == "version" {

	}

	if flags[0] == "add" {

	}

	if flags[0] == "list" {

	}

	if flags[0] == "remove" {

	}
}
