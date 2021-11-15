package flog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseFormat(t *testing.T) {
	a := assert.New(t)

	format, err := ParseFormat("apache_common")
	a.Equal("apache_common", format, "format should be apache_common")
	a.NoError(err, "there should be no error")

	format, err = ParseFormat("unknown")
	a.Equal("", format, "format should be empty string when invalid format is given")
	a.Error(err, "there should be an error when invalid format is given")
}

func TestParseNumber(t *testing.T) {
	a := assert.New(t)

	number, err := ParseNumber(1000)
	a.Equal(1000, number, "number should be 1000")
	a.NoError(err, "there should be no error")

	number, err = ParseNumber(-1000)
	a.Equal(0, number, "number should be 0 when negative is given")
	a.Error(err, "there should be an error when negative is given")
}

func TestParseBytes(t *testing.T) {
	a := assert.New(t)

	bytes, err := ParseBytes(1024 * 1024)
	a.Equal(1024*1024, bytes, "bytes should be 1048576")
	a.NoError(err, "there should be no error")

	bytes, err = ParseBytes(-1024)
	a.Equal(0, bytes, "bytes should be 0 when negative is given")
	a.Error(err, "there should be an error when negative is given")
}

func TestParseSleep(t *testing.T) {
	a := assert.New(t)

	sleep, err := ParseSleep("10")
	a.Equal(10*time.Second, sleep, "sleep should be 10s")
	a.NoError(err, "there should be no error")

	sleep, err = ParseSleep("20ms")
	a.Equal(20*time.Millisecond, sleep, "sleep should be 20ms")
	a.NoError(err, "there should be no error")

	sleep, err = ParseSleep("3s")
	a.Equal(3*time.Second, sleep, "sleep should be 3s")
	a.NoError(err, "there should be no error")

	sleep, err = ParseSleep("5.5")
	a.Equal(time.Duration(5.5*float64(time.Second)), sleep, "sleep should be 5.5s")
	a.NoError(err, "there should be no error")

	sleep, err = ParseSleep("-10")
	a.Equal(time.Duration(0), sleep, "sleep should be 0 when negative is given")
	a.Error(err, "there should be an error when negative is given")
}

func TestParseDelay(t *testing.T) {
	a := assert.New(t)

	delay, err := ParseDelay("10")
	a.Equal(10*time.Second, delay, "delay should be 10s")
	a.NoError(err, "there should be no error")

	delay, err = ParseDelay("20ms")
	a.Equal(20*time.Millisecond, delay, "delay should be 20ms")
	a.NoError(err, "there should be no error")

	delay, err = ParseDelay("3s")
	a.Equal(3*time.Second, delay, "delay should be 3s")
	a.NoError(err, "there should be no error")

	delay, err = ParseDelay("5.5")
	a.Equal(time.Duration(5.5*float64(time.Second)), delay, "delay should be 5.5s")
	a.NoError(err, "there should be no error")

	delay, err = ParseDelay("-10")
	a.Equal(time.Duration(0), delay, "delay should be 0 when negative is given")
	a.Error(err, "there should be an error when negative is given")
}

func TestParseSplitBy(t *testing.T) {
	a := assert.New(t)

	splitBy, err := ParseSplitBy(200)
	a.Equal(200, splitBy, "split-by should be 200")
	a.NoError(err, "there should be no error")

	splitBy, err = ParseSplitBy(-200)
	a.Equal(0, splitBy, "split-by should be 0 when negative is given")
	a.Error(err, "there should be an error when negative is given")
}

func TestParseOptions(t *testing.T) {
	a := assert.New(t)

	option := ParseOptions()
	a.Equal(defaultOptions(), option, "without flags, option should be default one")
}
