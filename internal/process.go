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
		if len(params) > 0 && params[0] != "" {
			return fmt.Errorf("-version doesn't support arguments")
		}
		return ProcessVersion()
	}

	if flags[0] == "config" {
		if len(params) > 0 && params[0] != ""{
			return fmt.Errorf("-config doesn't support arguments")
		}
		return ProcessConfig()
	}

	if flags[0] == "add" {
		if len(params) < 1 {
			return fmt.Errorf("nothing to save!")
		}
		return ProcessAdd(params[0])
	}

	if flags[0] == "list" {
		if len(params) > 0 && params[0] != ""{
			return fmt.Errorf("-list doesn't support arguments")
		}
		return ProcessList()
	}

	if flags[0] == "remove" {
		return fmt.Errorf("operation not implemented")
	}

	return fmt.Errorf("unknown error")
}
