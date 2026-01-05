package internal

import "fmt"

func ProcessHelp() error {
	fmt.Println("Usage: go-memo [options] <command> [args]\n\nGo-memo is intended to be used as a tool for taking quick notes or memos and storing them locally.\n\nOptions:\n  -h, --help      Show this help message and exit\n  -v, --version   Show version information and exit\n  -q, --quiet     Suppress non-error output\n  -C <dir>        Run as if toolname was started in <dir>\n\nExamples:\n  toolname init\n  toolname build --quiet\n  toolname run\n\nFor more information, see the documentation or run:\n  toolname <command> --help")
	
	userDataPath, err := GetUserDataPath()
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Config file path:   \"~/.go-memo/.config\"")
	fmt.Printf("User data path:     %q\n", userDataPath)

	return err
}
