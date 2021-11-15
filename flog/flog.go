package flog

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Generate generates the logs with given options
func Generate(option *Option) error {
	var (
		splitCount = 1
		created    = time.Now()

		interval time.Duration
		delay    time.Duration
	)

	if option.Delay > 0 {
		interval = option.Delay
		delay = interval
	}
	if option.Sleep > 0 {
		interval = option.Sleep
	}

	logFileName := option.Output
	writer, err := NewWriter(option.Type, logFileName)
	if err != nil {
		return err
	}

	if option.Forever {
		for {
			time.Sleep(delay)
			log := NewLog(option.Format, created)
			_, _ = writer.Write([]byte(log + "\n"))
			created = created.Add(interval)
		}
	}

	if option.Bytes == 0 {
		// Generates the logs until the certain number of lines is reached
		for line := 0; line < option.Number; line++ {
			time.Sleep(delay)
			log := NewLog(option.Format, created)
			_, _ = writer.Write([]byte(log + "\n"))

			if (option.Type != "stdout") && (option.SplitBy > 0) && (line > option.SplitBy*splitCount) {
				_ = writer.Close()
				fmt.Println(logFileName, "is created.")

				logFileName = NewSplitFileName(option.Output, splitCount)
				writer, _ = NewWriter(option.Type, logFileName)

				splitCount++
			}
			created = created.Add(interval)
		}
	} else {
		// Generates the logs until the certain size in bytes is reached
		bytes := 0
		for bytes < option.Bytes {
			time.Sleep(delay)
			log := NewLog(option.Format, created)
			_, _ = writer.Write([]byte(log + "\n"))

			bytes += len(log)
			if (option.Type != "stdout") && (option.SplitBy > 0) && (bytes > option.SplitBy*splitCount+1) {
				_ = writer.Close()
				fmt.Println(logFileName, "is created.")

				logFileName = NewSplitFileName(option.Output, splitCount)
				writer, _ = NewWriter(option.Type, logFileName)

				splitCount++
			}
			created = created.Add(interval)
		}
	}

	if option.Type != "stdout" {
		_ = writer.Close()
		fmt.Println(logFileName, "is created.")
	}
	return nil
}

// NewWriter returns a closeable writer corresponding to given log type
func NewWriter(logType string, logFileName string) (io.WriteCloser, error) {
	switch logType {
	case "stdout":
		return os.Stdout, nil
	case "log":
		logFile, err := os.Create(logFileName)
		if err != nil {
			return nil, err
		}
		return logFile, nil
	case "gz":
		logFile, err := os.Create(logFileName)
		if err != nil {
			return nil, err
		}
		return gzip.NewWriter(logFile), nil
	default:
		return nil, nil
	}
}

// NewLog creates a log for given format
func NewLog(format string, t time.Time) string {
	switch format {
	case "apache_common":
		return NewApacheCommonLog(t)
	case "apache_combined":
		return NewApacheCombinedLog(t)
	case "apache_error":
		return NewApacheErrorLog(t)
	case "rfc3164":
		return NewRFC3164Log(t)
	case "rfc5424":
		return NewRFC5424Log(t)
	case "common_log":
		return NewCommonLogFormat(t)
	case "json":
		return NewJSONLogFormat(t)
	case "logfmt":
		return NewLogFmtLogFormat(t)
	default:
		return ""
	}
}

// NewSplitFileName creates a new file path with split count
func NewSplitFileName(path string, count int) string {
	logFileNameExt := filepath.Ext(path)
	pathWithoutExt := strings.TrimSuffix(path, logFileNameExt)
	return pathWithoutExt + strconv.Itoa(count) + logFileNameExt
}
