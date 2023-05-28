package main

import (
	"bufio"
	"os"
)

func openFile(ac int) (*os.File, error) {
	file, err := os.OpenFile("emails.txt", os.O_APPEND|ac, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			err := createFile()
			if err != nil {
				return nil, err
			}
			return file, nil
		}
		return nil, err
	}
	return file, nil
}

func createFile() error {
	file, err := os.Create("emails.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func isEmailAlreadyExists(newEmail string) bool {
	file, err := openFile(os.O_RDONLY)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == newEmail {
			return true
		}
	}
	return false
}
