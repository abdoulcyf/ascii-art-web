package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFileToStr(dir, filename string) (string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Seek to the start of the file
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", fmt.Errorf("error seeking to the start of the file: %v", err)
	}

	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error scanning file: %v", err)
	}

	return content, nil
}
