package main

import (
	"math/rand"
	"time"

	"github.com/mingrammer/cfmt"
	"github.com/mingrammer/flog/flog"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	opts := flog.ParseOptions()
	if err := flog.Run(opts); err != nil {
		cfmt.Warningln(err.Error())
	}
}
