package main

import (
	"github.com/mingrammer/cfmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	opts := ParseOptions()
	if err := Run(opts); err != nil {
		cfmt.Warningln(err.Error())
	}
}
