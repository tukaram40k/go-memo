package internal

import "slices"

func GetFlags(flags []string) []string {
	trimLeading(flags)
	var processedFlags []string

	for _, flag := range flags {
		if flag == "help" || flag == "h" && !slices.Contains(processedFlags, "help") {
			processedFlags = append(processedFlags, "help")
			continue
		}

		if flag == "version" || flag == "v" && !slices.Contains(processedFlags, "version") {
			processedFlags = append(processedFlags, "version")
			continue
		}

		if flag == "config" || flag == "c" && !slices.Contains(processedFlags, "config") {
			processedFlags = append(processedFlags, "config")
			continue
		}

		if flag == "add" || flag == "a" && !slices.Contains(processedFlags, "add") {
			processedFlags = append(processedFlags, "add")
			continue
		}

		if flag == "list" || flag == "l" && !slices.Contains(processedFlags, "list") {
			processedFlags = append(processedFlags, "list")
			continue
		}

		if flag == "remove" || flag == "r" && !slices.Contains(processedFlags, "remove") {
			processedFlags = append(processedFlags, "remove")
			continue
		}

		processedFlags = append(processedFlags, "error")
		break
	}

	return processedFlags
}

// remove leading '-'
func trimLeading(flags []string) {
	for i, flag := range flags {
		if len(flag) > 0 && flag[0] == '-' {
			flags[i] = flag[1:]
		}
	}
}
