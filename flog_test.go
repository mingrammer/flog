package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func ExampleNewLog() {
	rand.Seed(11)

	monkey.Patch(time.Now, func() time.Time { return stopped })
	defer monkey.Unpatch(time.Now)

	created := time.Now()
	fmt.Println(NewLog("apache_common", created))
	fmt.Println(NewLog("apache_combined", created))
	fmt.Println(NewLog("apache_error", created))
	fmt.Println(NewLog("rfc3164", created))
	fmt.Println(NewLog("rfc5424", created))
	fmt.Println(NewLog("common_log", created))
	fmt.Println(NewLog("unknown", created))
	// Output:
	// 222.83.191.222 - - [22/04/2018:09:30:00 +0000] "DELETE /innovate/next-generation HTTP/1.0" 302 24570
	// 174.144.199.149 - mosciski7760 [22/04/2018:09:30:00 +0000] "GET /networks/revolutionary HTTP/1.0" 204 70360 "https://www.directeyeballs.org/streamline" "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_7) AppleWebKit/5312 (KHTML, like Gecko) Chrome/37.0.821.0 Mobile Safari/5312"
	// [Sun Apr 22 09:30:00 2018] [tempore:warn] [pid 7401:tid 4039] [client: 184.155.77.136] Try to synthesize the SMS capacitor, maybe it will compress the online program!
	// <25>Apr 22 09:30:00 fay5424 ipsam[5455]: Try to reboot the SMS bandwidth, maybe it will synthesize the mobile transmitter!
	// <118>2 2018-04-22T09:30:00.000Z legacyb2b.com natus 8605 ID854 - Try to synthesize the JBOD matrix, maybe it will calculate the 1080p application!
	// 34.205.100.144 - - [22/Apr/2018:09:30:00 +0000] "DELETE /empower/revolutionary/rich HTTP/1.0" 416 19950
	//
}

func TestNewSplitFileName(t *testing.T) {
	a := assert.New(t)

	splitFileName := NewSplitFileName("/path/to/file/generated.log", 1)
	a.Equal("/path/to/file/generated1.log", splitFileName, "filename should be '/path/to/file/generated1.log'")
}
