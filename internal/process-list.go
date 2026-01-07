package internal

import (
	"fmt"
	"os"
	"strings"
)

func ProcessList() error {
	userDataPath, err := GetUserDataPath()
	if err != nil {
		return err
	}

	userFilePath, err := GetUserFilePath(userDataPath)
	if err != nil {
		return err
	}

	if err := printEachLine(userFilePath); err != nil {
		return err
	}

	return nil
}

func printEachLine(userFilePath string) error {
	content, err := os.ReadFile(userFilePath)
	if err != nil {
		return err
	}

	lines := strings.SplitSeq(string(content), "\n")

	for line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		fmt.Println(line)
	}

	return nil
}