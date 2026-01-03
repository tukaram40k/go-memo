package internal

import (
	"os"
	"path/filepath"
)

const defaultConfig = "# go-memo configuration\npath = ~/Documents/go-memo"

// read or create .config file
func GetConfig() (string, error) {
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

	return string(content), nil
}

// creates full path to .config
func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".go-memo", ".config"), nil
}
