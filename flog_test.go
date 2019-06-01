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
	// 222.83.191.222 - - [22/04/2018:09:30:00 +0000] "DELETE /innovate/next-generation HTTP/1.1" 406 7610
	// 144.199.149.125 - waelchi7603 [22/04/2018:09:30:00 +0000] "PUT /revolutionary HTTP/1.1" 301 8089 "https://www.futureaggregate.io/users" "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_5 rv:4.0; en-US) AppleWebKit/536.38.2 (KHTML, like Gecko) Version/6.0 Safari/536.38.2"
	// [Sun Apr 22 09:30:00 2018] [eaque:error] [pid 3748:tid 2783] [client: 54.26.161.221] We need to transmit the open-source JSON capacitor!
	// <15>Apr 22 09:30:00 parisian2785 eius[350]: You can't quantify the capacitor without connecting the virtual XSS transmitter!
	// <12>2 2018-04-22T09:30:00.000Z internalmagnetic.io necessitatibus 3954 ID418 - Try to navigate the AGP feed, maybe it will quantify the optical monitor!
	// 7.101.102.40 - dibbert8018 [22/Apr/2018:09:30:00 +0000] "PATCH /empower/leading-edge/markets/whiteboard HTTP/1.0" 502 21224
	//
}

func TestNewSplitFileName(t *testing.T) {
	a := assert.New(t)

	splitFileName := NewSplitFileName("/path/to/file/generated.log", 1)
	a.Equal("/path/to/file/generated1.log", splitFileName, "filename should be '/path/to/file/generated1.log'")
}
