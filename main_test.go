package test_main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	otu := NewOverflowTest(t)

	t.Run("Runs test.cdc tx successfully", func(t *testing.T) {
		otu.HelloWorldTx()
	})

	t.Run("Runs test.cdc script successfully", func(t *testing.T) {
		result := otu.HelloWorldScript()
		assert.Equal(t, "Hello World.. from a script!", result)
	})
}
