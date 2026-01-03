package internal

import (
	"fmt"
	"os"
)

func ProcessHelp() {
	fmt.Println("Usage: toolname [options] <command> [args]\n\nA brief description of what the tool does.\n\nCommands:\n  init            Initialize a new project\n  build           Build the project\n  run             Run the project\n  clean           Remove generated files\n\nOptions:\n  -h, --help      Show this help message and exit\n  -v, --version   Show version information and exit\n  -q, --quiet     Suppress non-error output\n  -C <dir>        Run as if toolname was started in <dir>\n\nExamples:\n  toolname init\n  toolname build --quiet\n  toolname run\n\nFor more information, see the documentation or run:\n  toolname <command> --help")
	os.Exit(0)
}
