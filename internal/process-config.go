package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const defaultConfig = "# go-memo configuration\npath = ~/Documents/go-memo"
const defaultUserFileName = "go-memo-data.txt"
const defaultUserFileContent = "# go-memo user data file\n"

// read or create .config file
func GetUserDataPath() (string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return "", err
	}

	// check if file doesnt exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// create directory with rwxr-xr-x
    if err := os.Mkdir(filepath.Dir(configPath), 0755); err != nil {
      return "", err
    }
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

// create/locate user data file and return path to it
func GetUserFilePath(userDataPath string) (string, error) {
	if userDataPath == "" {
		return "", fmt.Errorf("empty path to user file")
	}

	if strings.HasPrefix(userDataPath, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get home directory: %w", err)
		}
		userDataPath = filepath.Join(home, userDataPath[2:])
	}

	path := filepath.Join(userDataPath, defaultUserFileName)

	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("no user data directory found.\n")
		fmt.Printf("creating directory: %s\n", dir)

		// create directory with rwxr-xr-x
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", fmt.Errorf("failed to create directory: %w", err)
		}
	}

	// check if file exists
	_, err := os.Stat(path); if err == nil {
		return path, nil
	}

	if os.IsNotExist(err) {
		fmt.Println("no user data file found.")
		fmt.Printf("writing to new user data file at: %s\n\n", path)

		// create file with rw-r--r--
		err := os.WriteFile(path, []byte(defaultUserFileContent), 0644)
		if err != nil {
			return "", fmt.Errorf("failed to create user data file: %w", err)
		}

		fmt.Println("user data file created successfully.")
		return path, nil
	}

	return "", fmt.Errorf("cannot access user file: %w", err)
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
