package main

import (
	"github.com/mingrammer/cfmt"
)

func main() {
	opts := ParseOptions()
	if err := Run(opts); err != nil {
		cfmt.Warningln(err.Error())
	}
}
