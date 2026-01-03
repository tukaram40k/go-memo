package internal

import "slices"

func GetFlags(flags []string) []string {
	trimLeading(flags)
	var processedFlags []string

	for _, flag := range flags {
		if !isAllowed(flag) {
			if !slices.Contains(processedFlags, "error") {
				processedFlags = append(processedFlags, "error")
			}
		} else {
			if !slices.Contains(processedFlags, flag) {
				processedFlags = append(processedFlags, flag)
			}
		}
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

// check if flag is allowed
func isAllowed(flag string) bool {
	allowedFlags := []string{"help", "version", "add", "list", "remove"}
	return slices.Contains(allowedFlags, flag)
}
