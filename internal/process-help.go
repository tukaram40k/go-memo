package internal

import "fmt"

func ProcessHelp() error {
	fmt.Println("Usage: go-memo [options] [args]\n\nGo-memo is intended to be used as a tool for taking quick notes or memos and storing them locally.\n\nOptions:\n  -h, --help         Show this help message and exit\n  -v, --version      Show version information and exit\n  -c, --config       Configure the tool\n  -a, --add <args>   Add a new item to the list\n  -l, --list         List all entries in the list\n  -r, --remove       Remove all entries from the list\n\nExample usage:\n  go-memo -a The first item to be added\n  go-memo -l")

	userDataPath, err := GetUserDataPath()
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Config file path:   \"~/.go-memo/.config\"")
	fmt.Printf("User data path:     %q\n", userDataPath)

	return err
}
