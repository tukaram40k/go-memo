package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const defaultConfig = "# go-memo configuration\npath = ~/Documents/go-memo"

// read or create .config file
func GetUserDataPath() (string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return "", err
	}

	// check if file doesnt exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// create directory with rwxr-xr-x
		os.Mkdir(filepath.Dir(configPath), 0755)
		// create file with rw-r--r--
		if err := os.WriteFile(configPath, []byte(defaultConfig), 0644); err != nil {
			return "", err
		}
	}

	// read existing file
	content, err := os.ReadFile(configPath)
	if err != nil {
		return "", err
	}

	userDataPath, err := extractUserDataPath(string(content))
	if err != nil {
		return "", err
	}

	return userDataPath, nil
}

// creates full path to .config
func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".go-memo", ".config"), nil
}

// extract path to user data from .config contents
func extractUserDataPath(content string) (string, error) {
	lines := strings.SplitSeq(content, "\n")

	for line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "path = ") {
			path := strings.TrimPrefix(line, "path = ")
			path = strings.TrimSpace(path)

			if path == "" {
				return "", fmt.Errorf("empty path in .config")
			}

			return path, nil
		}
	}

	return "", fmt.Errorf("path not found in .config")
}
