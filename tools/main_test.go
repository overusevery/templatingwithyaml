package main

import (
	"bytes"
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
