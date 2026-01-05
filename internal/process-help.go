package internal

import (
	"fmt"
)

func ProcessHelp() {
	fmt.Println("Usage: toolname [options] <command> [args]\n\nA brief description of what the tool does.\n\nCommands:\n  init            Initialize a new project\n  build           Build the project\n  run             Run the project\n  clean           Remove generated files\n\nOptions:\n  -h, --help      Show this help message and exit\n  -v, --version   Show version information and exit\n  -q, --quiet     Suppress non-error output\n  -C <dir>        Run as if toolname was started in <dir>\n\nExamples:\n  toolname init\n  toolname build --quiet\n  toolname run\n\nFor more information, see the documentation or run:\n  toolname <command> --help")
	
	userDataPath, err := GetUserDataPath()
	if err != nil {
		return
	}

	fmt.Println()
	fmt.Println("Config file path:   \"~/.go-memo/.config\"")
	fmt.Printf("User data path:     %q\n", userDataPath)
}
