package io

import (
	"bufio"
	"os"
)

func ReadFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var fileContent string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContent += scanner.Text() + "\n"
	}

	return fileContent, scanner.Err()
}
