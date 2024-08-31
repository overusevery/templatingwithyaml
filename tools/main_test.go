package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fill(t *testing.T) {
	t.Run("fill in template", func(t *testing.T) {
		var out bytes.Buffer
		fill(&out)
		assert.Equal(t, mustReadFile("internal/test/example_output"), out.String())
	})
}

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
