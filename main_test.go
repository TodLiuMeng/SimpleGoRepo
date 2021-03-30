package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	ast := assert.New(t)
	ast.Equal("hello", "hello")
}



