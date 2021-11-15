package flog

import (
	"fmt"
	"math/rand"
)

func ExampleRandResourceURI() {
	rand.Seed(11)
	fmt.Print(RandResourceURI())
	// Output: /target
}
