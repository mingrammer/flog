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

	fmt.Println(NewLog("apache_common", 0))
	// Output: 222.83.191.222 - Kozey7157 697 [2018-04-22T09:30:00Z] "DELETE /innovate/next-generation" 302 24570

	fmt.Println(NewLog("apache_combined", 0))
	// Output: 222.83.191.222 - Kozey7157 119 [2018-04-22T09:30:00Z] "DELETE /innovate/next-generation" 302 81317 "https://www.forwardholistic.biz/mission-critical/synergize/morph/sticky" "Mozilla/5.0 (Windows NT 5.01) AppleWebKit/5320 (KHTML, like Gecko) Chrome/40.0.875.0 Mobile Safari/5320"

	fmt.Println(NewLog("apache_error", 0))
	// Output: [2018-04-22T09:30:00Z] [quia:crit] [pid 4214:tid 6037] [client: 90.151.9.107] If we back up the program, we can get to the SSL sensor through the redundant SAS program!

	fmt.Println(NewLog("rfc3164", 0))
	// Output: 2018-04-22T09:30:00Z Daniel2872 info-mediaries[9030]: I'Ll reboot the auxiliary USB protocol, that should sensor the SAS capacitor!

	fmt.Println(NewLog("unknown", 0))
	// Output:
}

func TestNewSplitFileName(t *testing.T) {
	a := assert.New(t)

	splitFileName := NewSplitFileName("/path/to/file/generated.log", 1)
	a.Equal("/path/to/file/generated1.log", splitFileName, "filename should be '/path/to/file/generated1.log'")
}
