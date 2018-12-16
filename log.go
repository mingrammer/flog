package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
)

const (
	// ApacheCommonLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} HTTP/1.0" {response-code} {bytes}
	ApacheCommonLog = "%s - %s %d [%s] \"%s %s\" %d %d"
	// ApacheCombinedLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} HTTP/1.0" {response-code} {bytes} "{referrer}" "{agent}"
	ApacheCombinedLog = "%s - %s %d [%s] \"%s %s\" %d %d \"%s\" \"%s\""
	// ApacheErrorLog : [{timestamp}] [{module}:{severity}] [pid {pid}:tid {thread-id}] [client: %{client}] %{message}
	ApacheErrorLog = "[%s] [%s:%s] [pid %d:tid %d] [client: %s] %s"
	// RFC3164Log : <priority>{timestamp} {hostname} {application}[{pid}]: {message}
	RFC3164Log = "<%d>%s %s %s[%d]: %s"
)

// NewApacheCommonLog creates a log string with apache common log format
func NewApacheCommonLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheCommonLog,
		gofakeit.IPv4Address(),
		gofakeit.Username(),
		gofakeit.Number(0, 1000),
		t.Format(time.RFC3339),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
	)
}

// NewApacheCombinedLog creates a log string with apache combined log format
func NewApacheCombinedLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheCombinedLog,
		gofakeit.IPv4Address(),
		gofakeit.Username(),
		gofakeit.Number(1, 1000),
		t.Format(time.RFC3339),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		gofakeit.StatusCode(),
		gofakeit.Number(30, 100000),
		gofakeit.URL(),
		gofakeit.UserAgent(),
	)
}

// NewApacheErrorLog creates a log string with apache error log format
func NewApacheErrorLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheErrorLog,
		t.Format(time.RFC3339),
		gofakeit.Word(),
		gofakeit.LogLevel("apache"),
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 10000),
		gofakeit.IPv4Address(),
		gofakeit.HackerPhrase(),
	)
}

// NewRFC3164Log creates a log string with syslog (RFC3164) format
func NewRFC3164Log(t time.Time) string {
	return fmt.Sprintf(
		RFC3164Log,
		gofakeit.Number(0, 191),
		t.Format(RFC3164),
		gofakeit.Username(),
		gofakeit.BuzzWord(),
		gofakeit.Number(1, 10000),
		gofakeit.HackerPhrase(),
	)
}
