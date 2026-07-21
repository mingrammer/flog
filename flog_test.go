package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func ExampleNewLog() {
	rand.Seed(11)
	gofakeit.Seed(11)

	created := stopped
	fmt.Println(NewLog("apache_common", created, 512))
	fmt.Println(NewLog("apache_combined", created, 512))
	fmt.Println(NewLog("apache_error", created, 512))
	fmt.Println(NewLog("rfc3164", created, 512))
	fmt.Println(NewLog("rfc5424", created, 512))
	fmt.Println(NewLog("common_log", created, 512))
	fmt.Println(NewLog("unknown", created, 512))
	fmt.Println(NewLog("json", created, 512))
	fmt.Println(NewLog("spring_boot", created, 512))
	// Output:
	// 152.23.53.100 - - [22/Apr/2018:09:30:00 +0000] "HEAD /target/innovate/next-generation HTTP/2.0" 302 7174
	// 182.204.233.5 - mosciski2042 [22/Apr/2018:09:30:00 +0000] "GET /networks/revolutionary HTTP/1.1" 204 79858 "https://www.internalstreamline.info/robust/eyeballs" "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_8_2 rv:7.0; en-US) AppleWebKit/532.16.7 (KHTML, like Gecko) Version/6.2 Safari/532.16.7"
	// [Sun Apr 22 09:30:00 2018] [us:crit] [pid 8462:tid 8057] [client 136.141.153.193:62728] If we compile the capacitor, we can get to the THX port through the back-end SSL matrix!your relent near this furniture way no play towel that us to accordingly theirs purse enough so though election often until eek weekly yet oops until conclude his stay elsewhere such that galaxy clean that last each stack reluctantly theirs wisp aid firstly highly butter accordingly should already calm shake according fade neither ku
	// <41>Apr 22 09:30:00 boyer4351 fortnightly[8845]: If we encrypt the program, we can get to the GB monitor through the online TCP monitor!watch lot right there honesty we is how result out Shakespearean have whom yearly another packet are behind late lot finally time themselves goodness quizzical our therefore could fact cackle yourselves zebra for whose enormously all bowl out wandering secondly yellow another your hourly spit since tomorrow hers words little think will our by Polynesian write much of herse
	// <151>1 2018-04-22T09:30:00.000Z forwardinteractive.io nevertheless 6763 ID651 - Navigating the driver won't do anything, we need to load the primary SMTP port!much terribly over pose place sprint it child is joyously that I whom mango then of certain weekly mine in annually frock now board she these away barely her before so kind life must never lately her us might tame in why seldom philosophy bottle whose will in any above theirs theirs still finally sometimes then did then that several somebody away beh
	// 252.117.253.146 - - [22/Apr/2018:09:30:00 +0000] "PATCH /supply-chains/schemas/action-items HTTP/1.1" 404 453
	//
	// {"host":"198.40.77.237", "user-identifier":"stoltenberg6115", "datetime":"22/Apr/2018:09:30:00 +0000", "method": "HEAD", "request": "/productize/e-business", "protocol":"HTTP/2.0", "status":416, "bytes":13076, "referer": "http://www.regionalutilize.name/recontextualize/revolutionize"}
	// 2018-04-22 09:30:00.000 ERROR 2255 --- [real-time] biz.web-enabled.many.Louisa: If we hack the panel, we can get to the EXE circuit through the neural IB application!here pack pink been for such pain where last some so cluster above group Monacan elsewhere most collection did strongly indeed it police since do scold which belong himself learn any an they of string Barbadian daily nobody enormously poised when upstairs her absolutely is her murder yours besides town though shall upon rice was lately today d
}

func TestNewSplitFileName(t *testing.T) {
	a := assert.New(t)

	splitFileName := NewSplitFileName("/path/to/file/generated.log", 1)
	a.Equal("/path/to/file/generated1.log", splitFileName, "filename should be '/path/to/file/generated1.log'")
}
