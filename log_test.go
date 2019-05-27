package main

import (
	"fmt"
	"math/rand"
	"time"

	"bou.ke/monkey"
)

var stopped = time.Date(2018, 04, 22, 9, 30, 0, 0, time.UTC)

func ExampleNewApacheCommonLog() {
	rand.Seed(11)

	monkey.Patch(time.Now, func() time.Time { return stopped })
	defer monkey.Unpatch(time.Now)

	created := time.Now()
	fmt.Println(NewApacheCommonLog(created))
	// Output: 222.83.191.222 - - [22/04/2018:09:30:00 +0000] "DELETE /innovate/next-generation HTTP/1.0" 302 24570
}

func ExampleNewApacheCombinedLog() {
	rand.Seed(11)

	monkey.Patch(time.Now, func() time.Time { return stopped })
	defer monkey.Unpatch(time.Now)

	created := time.Now()
	fmt.Println(NewApacheCombinedLog(created))
	// Output: 222.83.191.222 - - [22/04/2018:09:30:00 +0000] "DELETE /innovate/next-generation HTTP/1.0" 302 81317 "https://www.forwardholistic.biz/mission-critical/synergize/morph/sticky" "Mozilla/5.0 (Windows NT 5.01) AppleWebKit/5320 (KHTML, like Gecko) Chrome/40.0.875.0 Mobile Safari/5320"
}

func ExampleNewApacheErrorLog() {
	rand.Seed(11)

	monkey.Patch(time.Now, func() time.Time { return stopped })
	defer monkey.Unpatch(time.Now)

	created := time.Now()
	fmt.Println(NewApacheErrorLog(created))
	// Output: [Sun Apr 22 09:30:00 2018] [quia:crit] [pid 4214:tid 6037] [client: 90.151.9.107] If we back up the program, we can get to the SSL sensor through the redundant SAS program!
}

func ExampleNewRFC3164Log() {
	rand.Seed(11)

	monkey.Patch(time.Now, func() time.Time { return stopped })
	defer monkey.Unpatch(time.Now)

	created := time.Now()
	fmt.Println(NewRFC3164Log(created))
	// Output: <24>Apr 22 09:30:00 moen8727 quo[3160]: If we back up the program, we can get to the SSL sensor through the redundant SAS program!
}

func ExampleNewRFC5424Log() {
	rand.Seed(11)

	monkey.Patch(time.Now, func() time.Time { return stopped })
	defer monkey.Unpatch(time.Now)

	created := time.Now()
	fmt.Println(NewRFC5424Log(created))
	// Output: <24>3 2018-04-22T09:30:00.000Z futurefunctionalities.biz nisi 9030 ID160 - If we back up the program, we can get to the SSL sensor through the redundant SAS program!
}

func ExampleNewCommonLogFormat() {
	rand.Seed(11)

	monkey.Patch(time.Now, func() time.Time { return stopped })
	defer monkey.Unpatch(time.Now)

	created := time.Now()
	fmt.Println(NewCommonLogFormat(created))
	// Output: 222.83.191.222 - - [22/Apr/2018:09:30:00 +0000] "DELETE /innovate/next-generation HTTP/1.0" 302 24570
}
