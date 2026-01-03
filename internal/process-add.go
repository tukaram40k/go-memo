package internal

import (
	"fmt"
	"os"
)

func ProcessAdd(params string) {
	config, err := GetConfig()
	if err != nil {
		fmt.Printf("configuration error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("params: %q\n", params)
	fmt.Printf("config file contents: %q\n", config)
}
