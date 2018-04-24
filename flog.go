package main

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
	splitCount := 1
	delta := time.Duration(0)

	logFileName := option.Output
	writer, err := NewWriter(option.Type, logFileName)
	if err != nil {
		return err
	}

	if option.Bytes == 0 {
		// Generate the logs up to maximum number of lines
		for line := 0; line < option.Number; line++ {
			log := NewLog(option.Format, delta)
			writer.Write([]byte(log))

			if (option.Type != "stdout") && (option.SplitBy > 0) && (line > option.SplitBy*splitCount) {
				writer.Close()
				fmt.Println(logFileName, "is created.")

				logFileName = NewSplitFileName(option.Output, splitCount)
				writer, _ = NewWriter(option.Type, logFileName)

				splitCount++
			}

			delta += time.Duration(option.Sleep) * time.Second
		}
	} else {
		// Generate the logs up to maximum size in byte
		bytes := 0
		for bytes < option.Bytes {
			log := NewLog(option.Format, delta)
			writer.Write([]byte(log))

			bytes += len(log)
			if (option.Type != "stdout") && (option.SplitBy > 0) && (bytes > option.SplitBy*splitCount+1) {
				writer.Close()
				fmt.Println(logFileName, "is created.")

				logFileName = NewSplitFileName(option.Output, splitCount)
				writer, _ = NewWriter(option.Type, logFileName)

				splitCount++
			}

			delta += time.Duration(option.Sleep) * time.Second
		}
	}

	if option.Type != "stdout" {
		writer.Close()
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
func NewLog(format string, delta time.Duration) string {
	switch format {
	case "apache_common":
		return NewApacheCommonLog(delta)
	case "apache_combined":
		return NewApacheCombinedLog(delta)
	case "apache_error":
		return NewApacheErrorLog(delta)
	case "rfc3164":
		return NewRFC3164Log(delta)
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
