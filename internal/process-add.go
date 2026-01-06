package internal

import (
	"fmt"
	"os"
)

func ProcessAdd(params string) error {
	userDataPath, err := GetUserDataPath()
	if err != nil {
		return err
	}

	userFilePath, err := GetUserFilePath(userDataPath)
	if err != nil {
		return err
	}

	if err := appendToUserFile(userFilePath, params); err != nil {
		return err
	}

	fmt.Println("saved!")

	return nil
}

func appendToUserFile(userFilePath string, content string) error {
	if content == "" {
		return fmt.Errorf("nothing to save!")
	}

	file, err := os.OpenFile(userFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening user data file: %w", err)
	}
	defer file.Close()

	content = fmt.Sprintf("- %s\n", content)

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("error appending to user data file: %w", err)
	}

	return nil
}
