package flog

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
	fmt.Println(NewLog("logfmt", created))
	// Output:
	// 222.83.191.222 - - [22/Apr/2018:09:30:00 +0000] "DELETE /innovate/next-generation HTTP/1.1" 406 7610
	// 144.199.149.125 - waelchi7603 [22/Apr/2018:09:30:00 +0000] "PUT /revolutionary HTTP/1.1" 301 8089 "https://www.futureaggregate.io/users" "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_5 rv:4.0; en-US) AppleWebKit/536.38.2 (KHTML, like Gecko) Version/6.0 Safari/536.38.2"
	// [Sun Apr 22 09:30:00 2018] [eaque:error] [pid 3748:tid 2783] [client 54.26.161.221:31944] Backing up the program won't do anything, we need to compress the optical PCI bandwidth!
	// <94>Apr 22 09:30:00 ortiz5384 vel[1775]: If we copy the firewall, we can get to the PCI firewall through the redundant SQL port!
	// <23>3 2018-04-22T09:30:00.000Z humaniterate.io iusto 544 ID177 - Use the optical RAM hard drive, then you can program the auxiliary feed!
	// 195.44.200.155 - kihn6187 [22/Apr/2018:09:30:00 +0000] "GET /revolutionary/e-markets/holistic/syndicate HTTP/2.0" 404 14503
	//
	// {"host":"13.108.182.26", "user-identifier":"bailey7205", "datetime":"22/Apr/2018:09:30:00 +0000", "method": "GET", "request": "/out-of-the-box/architectures/embrace", "protocol":"HTTP/1.0", "status":200, "bytes":5921, "referer": "http://www.dynamicexperiences.io/robust"}
	// host="169.19.25.250" user=rutherford8856 timestamp=2018-04-22T09:30:00.000Z method=HEAD request="/e-enable/virtual/partnerships" protocol=HTTP/2.0 status=400 bytes=16428 referer="http://www.investorstrategic.biz/strategic/vertical/scalable/ubiquitous"
}

func TestNewSplitFileName(t *testing.T) {
	a := assert.New(t)

	splitFileName := NewSplitFileName("/path/to/file/generated.log", 1)
	a.Equal("/path/to/file/generated1.log", splitFileName, "filename should be '/path/to/file/generated1.log'")
}
