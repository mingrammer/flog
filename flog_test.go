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
	fmt.Println(NewLog("json", created))
	// Output:
	// 222.83.191.222 - - [22/Apr/2018:09:30:00 +0000] "DELETE /innovate/next-generation HTTP/1.1" 406 7610
	// 144.199.149.125 - waelchi7603 [22/Apr/2018:09:30:00 +0000] "PUT /revolutionary HTTP/1.1" 301 8089 "https://www.futureaggregate.io/users" "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_5 rv:4.0; en-US) AppleWebKit/536.38.2 (KHTML, like Gecko) Version/6.0 Safari/536.38.2"
	// [Sun Apr 22 09:30:00 2018] [eaque:error] [pid 3748:tid 2783] [client 54.26.161.221:31944] Backing up the program won't do anything, we need to compress the optical PCI bandwidth!
	// <94>Apr 22 09:30:00 ortiz5384 vel[1775]: If we copy the firewall, we can get to the PCI firewall through the redundant SQL port!
	// <23>3 2018-04-22T09:30:00.000Z humaniterate.io iusto 544 ID177 [exampleSDID@877061 iut="6" eventSource="Application" eventID="380001"][examplePriority@86525 class="high" method="PUT" uri="/empower/leading-edge/benchmark" status_code="400" time_millis="80" remote_host="190.67.164.175" remote_ip_addr="116.227.143.59"] We need to program the open-source ADP pixel!
	// 15.48.13.108 - - [22/Apr/2018:09:30:00 +0000] "POST /cross-platform/extensible/out-of-the-box/architectures HTTP/2.0" 100 16077
	//
	// {"host":"11.53.30.203", "user-identifier":"wilkinson1680", "datetime":"22/Apr/2018:09:30:00 +0000", "method": "POST", "request": "/bleeding-edge/morph", "protocol":"HTTP/2.0", "status":502, "bytes":10203, "referer": "https://www.globale-enable.net/leverage/integrated"}
}

func TestNewSplitFileName(t *testing.T) {
	a := assert.New(t)

	splitFileName := NewSplitFileName("/path/to/file/generated.log", 1)
	a.Equal("/path/to/file/generated1.log", splitFileName, "filename should be '/path/to/file/generated1.log'")
}
