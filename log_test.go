package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

var stopped = time.Date(2018, 04, 22, 9, 30, 0, 0, time.UTC)

func ExampleNewApacheCommonLog() {
	rand.Seed(11)
	gofakeit.Seed(11)

	fmt.Println(NewApacheCommonLog(stopped))
	// Output: 152.23.53.100 - - [22/Apr/2018:09:30:00 +0000] "HEAD /target/innovate/next-generation HTTP/2.0" 302 7174
}

func ExampleNewApacheCombinedLog() {
	rand.Seed(11)
	gofakeit.Seed(11)

	fmt.Println(NewApacheCombinedLog(stopped))
	// Output: 152.23.53.100 - - [22/Apr/2018:09:30:00 +0000] "HEAD /target/innovate/next-generation HTTP/2.0" 302 60732 "http://www.directsynergize.name/syndicate/holistic/scalable" "Mozilla/5.0 (Windows 98; Win 9x 4.90) AppleWebKit/5312 (KHTML, like Gecko) Chrome/37.0.804.0 Mobile Safari/5312"
}

func ExampleNewApacheErrorLog() {
	rand.Seed(11)
	gofakeit.Seed(11)

	fmt.Println(NewApacheErrorLog(stopped, 512))
	// Output: [Sun Apr 22 09:30:00 2018] [none:trace1-8] [pid 3590:tid 4801] [client 25.149.87.200:59704] The SSL sensor is down, decrypt the open-source program so we can transmit the SDD interface!by army party riches theirs instead here mine whichever that those instance growth has ouch enough Swiss us since down he she aha us you to upon how this this furniture way no play towel that us to accordingly theirs purse enough so though election often until eek weekly yet oops until conclude his stay elsewhere such that g
}

func ExampleNewRFC3164Log() {
	rand.Seed(11)
	gofakeit.Seed(11)

	fmt.Println(NewRFC3164Log(stopped, 512))
	// Output: <89>Apr 22 09:30:00 moen3645 keep[2362]: Copying the protocol won't do anything, we need to verify the redundant SAS program!still far unless army party riches theirs instead here mine whichever that those instance growth has ouch enough Swiss us since down he she aha us you to upon how this this furniture way no play towel that us to accordingly theirs purse enough so though election often until eek weekly yet oops until conclude his stay elsewhere such that galaxy clean that last each stack reluctantly t
}

func ExampleNewRFC5424Log() {
	rand.Seed(11)
	gofakeit.Seed(11)

	fmt.Println(NewRFC5424Log(stopped, 512))
	// Output: <89>1 2018-04-22T09:30:00.000Z futurefunctionalities.biz whom 4499 ID362 - Copying the protocol won't do anything, we need to verify the redundant SAS program!still far unless army party riches theirs instead here mine whichever that those instance growth has ouch enough Swiss us since down he she aha us you to upon how this this furniture way no play towel that us to accordingly theirs purse enough so though election often until eek weekly yet oops until conclude his stay elsewhere such that galaxy clean
}

func ExampleNewCommonLogFormat() {
	rand.Seed(11)
	gofakeit.Seed(11)

	fmt.Println(NewCommonLogFormat(stopped))
	// Output: 152.23.53.100 - - [22/Apr/2018:09:30:00 +0000] "HEAD /target/innovate/next-generation HTTP/2.0" 302 7174
}

func ExampleNewJSONLogFormat() {
	rand.Seed(11)
	gofakeit.Seed(11)

	fmt.Println(NewJSONLogFormat(stopped))
	// Output: {"host":"152.23.53.100", "user-identifier":"-", "datetime":"22/Apr/2018:09:30:00 +0000", "method": "HEAD", "request": "/target/innovate/next-generation", "protocol":"HTTP/2.0", "status":302, "bytes":7174, "referer": "http://www.directsynergize.name/syndicate/holistic/scalable"}
}

func ExampleNewSpringBootLogFormat() {
	rand.Seed(11)
	gofakeit.Seed(11)

	fmt.Println(NewSpringBootLogFormat(stopped, 512))
	// Output: 2018-04-22 09:30:00.000 ERROR 364 --- [info-mediaries] name.systems.trip.Santino: We need to verify the redundant SAS program!still far unless army party riches theirs instead here mine whichever that those instance growth has ouch enough Swiss us since down he she aha us you to upon how this this furniture way no play towel that us to accordingly theirs purse enough so though election often until eek weekly yet oops until conclude his stay elsewhere such that galaxy clean that last each stack reluctantly
}
