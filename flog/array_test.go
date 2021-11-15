package flog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainString(t *testing.T) {
	a := assert.New(t)

	arr := []string{"a", "b", "c", "d", "e"}
	a.Equal(true, containString(arr, "a"), "a is in arr")
	a.Equal(true, containString(arr, "e"), "e is in arr")
	a.Equal(false, containString(arr, "f"), "f is not in arr")
}
