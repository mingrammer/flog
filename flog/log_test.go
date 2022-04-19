package flog

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
	// Output: 222.83.191.222 - - [22/Apr/2018:09:30:00 +0000] "DELETE /innovate/next-generation HTTP/1.1" 406 7610
}

func ExampleNewApacheCombinedLog() {
	rand.Seed(11)

	monkey.Patch(time.Now, func() time.Time { return stopped })
	defer monkey.Unpatch(time.Now)

	created := time.Now()
	fmt.Println(NewApacheCombinedLog(created))
	// Output: 222.83.191.222 - - [22/Apr/2018:09:30:00 +0000] "DELETE /innovate/next-generation HTTP/1.1" 406 97484 "https://www.humanscalable.io/synergize/morph/sticky" "Mozilla/5.0 (Windows NT 5.01) AppleWebKit/5320 (KHTML, like Gecko) Chrome/40.0.875.0 Mobile Safari/5320"
}

func ExampleNewApacheErrorLog() {
	rand.Seed(11)

	monkey.Patch(time.Now, func() time.Time { return stopped })
	defer monkey.Unpatch(time.Now)

	created := time.Now()
	fmt.Println(NewApacheErrorLog(created))
	// Output: [Sun Apr 22 09:30:00 2018] [quia:crit] [pid 4214:tid 6037] [client 90.151.9.107:14075] Copying the protocol won't do anything, we need to copy the redundant SAS program!
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
	// Output: 222.83.191.222 - - [22/Apr/2018:09:30:00 +0000] "DELETE /innovate/next-generation HTTP/1.1" 406 7610
}

func ExampleNewJSONLogFormat() {
	rand.Seed(11)

	monkey.Patch(time.Now, func() time.Time { return stopped })
	defer monkey.Unpatch(time.Now)

	created := time.Now()
	fmt.Println(NewJSONLogFormat(created))
	// Output: {"host":"222.83.191.222", "user-identifier":"-", "datetime":"22/Apr/2018:09:30:00 +0000", "method": "DELETE", "request": "/innovate/next-generation", "protocol":"HTTP/1.1", "status":406, "bytes":7610, "referer": "https://www.humanscalable.io/synergize/morph/sticky"}
}
