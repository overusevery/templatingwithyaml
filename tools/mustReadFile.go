package main

import (
	"io"
	"log"
	"os"
)

// mustReadFile opens a file at the given path and reads its entire content as a string.
// If the file cannot be opened or read, the function logs the error and terminates the program.
func mustReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error readingfile:", err)
	}

	return string(content)
}
