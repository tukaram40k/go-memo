package internal

import (
	"fmt"
	"os"
)

func ProcessAdd(params string) {
	userDataPath, err := GetUserDataPath()
	if err != nil {
		fmt.Printf("configuration error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("params: %q\n", params)
	fmt.Printf("user data path: %q\n", userDataPath)
}
