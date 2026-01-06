package internal

import (
	"fmt"
)

func ProcessAdd(params string) error {
	userDataPath, err := GetUserDataPath(); if err != nil {
		return err
	}

	userFilePath, err := GetUserFilePath(userDataPath); if err != nil {
		return err
	}

	fmt.Printf("params: %q\n", params)
	fmt.Printf("user file path: %q\n", userFilePath)

	return nil
}
