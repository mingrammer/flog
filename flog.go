package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Generate generates the logs with given options
func Generate(option *Option) error {
	logFileName := option.Output
	logFile, err := os.Create(logFileName)
	if err != nil {
		return err
	}

	delta := time.Duration(option.Sleep) * time.Second
	splitCount := 1

	if option.Bytes == 0 {
		// Generate the logs up to maximum number of lines
		for line := 0; line < option.Number; line++ {
			log := NewLog(option.Format, delta)
			logFile.WriteString(log)

			if (option.SplitBy > 0) && (line >= option.SplitBy*splitCount) {
				logFile.Close()
				fmt.Println(logFileName, "is created.")

				logFileName = NewSplitFileName(option.Output, splitCount)
				logFile, _ = os.Create(logFileName)
				splitCount++
			}

			delta += time.Duration(option.Sleep) * time.Second
		}
		logFile.Close()
		fmt.Println(logFileName, "is created.")
	} else {
		// Generate the logs up to maximum size in byte
		bytes := 0
		for bytes < option.Bytes {
			delta += time.Duration(option.Sleep) * time.Second
			log := NewLog(option.Format, delta)
			logFile.WriteString(log)

			bytes += len(log)
			if (option.SplitBy > 0) && (bytes >= option.SplitBy*splitCount) {
				logFile.Close()
				fmt.Println(logFileName, "is created.")

				logFileName = NewSplitFileName(option.Output, splitCount)
				logFile, _ = os.Create(logFileName)
				splitCount++
			}

			delta += time.Duration(option.Sleep) * time.Second
		}
		logFile.Close()
		fmt.Println(logFileName, "is created.")
	}
	return nil
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
