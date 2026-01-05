package internal

import "fmt"

func ProcessAdd(params string) error {
	userDataPath, err := GetUserDataPath()
	if err != nil {
		return err
	}

	fmt.Printf("params: %q\n", params)
	fmt.Printf("user data path: %q\n", userDataPath)

	return nil
}
