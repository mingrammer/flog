package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

func ExampleRandResourceURI() {
	gofakeit.Seed(11)
	fmt.Print(RandResourceURI())
	// Output: /target/seamless
}
