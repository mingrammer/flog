package main

import (
	"fmt"
	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func ExampleNewLog() {
	rand.Seed(11)

	monkey.Patch(time.Now, func() time.Time { return stopped })
	defer monkey.Unpatch(time.Now)

	fmt.Print(NewLog("apache_common", 0))
	fmt.Print(NewLog("apache_combined", 0))
	fmt.Print(NewLog("apache_error", 0))
	fmt.Print(NewLog("rfc3164", 0))
	fmt.Print(NewLog("unknown", 0))
	// Output:
	// 222.83.191.222 - Kozey7157 697 [2018-04-22T09:30:00Z] "DELETE /innovate/next-generation" 302 24570
	// 174.144.199.149 - Mosciski7760 386 [2018-04-22T09:30:00Z] "GET /networks/revolutionary" 204 70360 "https://www.directeyeballs.org/streamline" "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_7) AppleWebKit/5312 (KHTML, like Gecko) Chrome/37.0.821.0 Mobile Safari/5312"
	// [2018-04-22T09:30:00Z] [tempore:warn] [pid 7401:tid 4039] [client: 184.155.77.136] Try to synthesize the SMS capacitor, maybe it will compress the online program!
	// 2018-04-22T09:30:00Z Wolff5542 Centralized[5680]: Quantifying the capacitor won't do anything, we need to parse the neural SAS program!
	//
}

func TestNewSplitFileName(t *testing.T) {
	a := assert.New(t)

	splitFileName := NewSplitFileName("/path/to/file/generated.log", 1)
	a.Equal("/path/to/file/generated1.log", splitFileName, "filename should be '/path/to/file/generated1.log'")
}
