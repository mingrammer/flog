package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsString(t *testing.T) {
	a := assert.New(t)

	arr := []string{"a", "b", "c", "d", "e"}
	a.Equal(true, ContainsString(arr, "a"), "a is in arr")
	a.Equal(true, ContainsString(arr, "e"), "e is in arr")
	a.Equal(false,ContainsString(arr, "f"), "f is not in arr")
}
