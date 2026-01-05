package internal

func GetParams(args []string) ([]string, []string) {
	var flags []string
	// arr with a single concatenated string
	params := []string{""}

	for _, arg := range args {
		if len(arg) > 1 && arg[0] == '-' {
			flags = append(flags, arg[1:])
		} else {
			params[0] += " " + arg
		}
	}

	if len(params[0]) > 0 && params[0][0] == ' ' {
		params[0] = params[0][1:]
	}
	return flags, params
}
