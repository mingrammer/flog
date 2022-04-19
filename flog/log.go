package flog

import (
	"fmt"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
)

const (
	// ApacheCommonLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	ApacheCommonLog = "%s - %s [%s] \"%s %s %s\" %d %d"
	// ApacheCombinedLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes} "{referrer}" "{agent}"
	ApacheCombinedLog = "%s - %s [%s] \"%s %s %s\" %d %d \"%s\" \"%s\""
	// ApacheErrorLog : [{timestamp}] [{module}:{severity}] [pid {pid}:tid {thread-id}] [client %{client}:{port}] %{message}
	ApacheErrorLog = "[%s] [%s:%s] [pid %d:tid %d] [client %s:%d] %s"
	// RFC3164Log : <priority>{timestamp} {hostname} {application}[{pid}]: {message}
	RFC3164Log = "<%d>%s %s %s[%d]: %s"
	// RFC5424Log : <priority>{version} {iso-timestamp} {hostname} {application} {pid} {message-id} {structured-data} {message}
	RFC5424Log = "<%d>%d %s %s %s %d ID%d %s %s"
	// CommonLogFormat : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	CommonLogFormat = "%s - %s [%s] \"%s %s %s\" %d %d"
	// JSONLogFormat : {"host": "{host}", "user-identifier": "{user-identifier}", "datetime": "{datetime}", "method": "{method}", "request": "{request}", "protocol": "{protocol}", "status", {status}, "bytes": {bytes}, "referer": "{referer}"}
	JSONLogFormat = `{"host":"%s", "user-identifier":"%s", "datetime":"%s", "method": "%s", "request": "%s", "protocol":"%s", "status":%d, "bytes":%d, "referer": "%s"}`
	// LogFmtLogFormat : host={host} user={user-identifier} timestamp={datetime} method={method} request="{request}" protocol={protocol} status={status} bytes={bytes} referer="{referer}"
	LogFmtLogFormat = `host="%s" user=%s timestamp=%s method=%s request="%s" protocol=%s status=%d bytes=%d referer="%s"`
)

// NewApacheCommonLog creates a log string with apache common log format
func NewApacheCommonLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheCommonLog,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(Apache),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
	)
}

// NewApacheCombinedLog creates a log string with apache combined log format
func NewApacheCombinedLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheCombinedLog,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(Apache),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
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
		t.Format(ApacheError),
		gofakeit.Word(),
		gofakeit.LogLevel("apache"),
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 10000),
		gofakeit.IPv4Address(),
		gofakeit.Number(1, 65535),
		gofakeit.HackerPhrase(),
	)
}

// NewRFC3164Log creates a log string with syslog (RFC3164) format
func NewRFC3164Log(t time.Time) string {
	return fmt.Sprintf(
		RFC3164Log,
		gofakeit.Number(0, 191),
		t.Format(RFC3164),
		strings.ToLower(gofakeit.Username()),
		gofakeit.Word(),
		gofakeit.Number(1, 10000),
		gofakeit.HackerPhrase(),
	)
}

// NewRFC5424Log creates a log string with syslog (RFC5424) format
func NewRFC5424Log(t time.Time) string {
	return fmt.Sprintf(
		RFC5424Log,
		gofakeit.Number(0, 191),
		gofakeit.Number(1, 3),
		t.Format(RFC5424),
		gofakeit.DomainName(),
		gofakeit.Word(),
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 1000),
		"-", // TODO: structured data
		gofakeit.HackerPhrase(),
	)
}

// NewCommonLogFormat creates a log string with common log format
func NewCommonLogFormat(t time.Time) string {
	return fmt.Sprintf(
		CommonLogFormat,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(CommonLog),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
	)
}

// NewJSONLogFormat creates a log string with json log format
func NewJSONLogFormat(t time.Time) string {
	return fmt.Sprintf(
		JSONLogFormat,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(CommonLog),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
		gofakeit.URL(),
	)
}

// NewLogFmtLogFormat creates a log string with logfmt log format
func NewLogFmtLogFormat(t time.Time) string {
	return fmt.Sprintf(
		LogFmtLogFormat,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(RFC5424),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
		gofakeit.URL(),
	)
}
