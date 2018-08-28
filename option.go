package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

const usage = `flog is a fake log generator for common log formats

Usage: flog [options]

Options:
  -f, --format string      Choose log format. ("apache_common"|"apache_combined"|"apache_error"|"rfc3164") (default "apache_common")
  -o, --output string      Output filename. Path-like is allowed. (default "generated.log")
  -t, --type string        Log output type. ("stdout"|"log"|"gz") (default "stdout")
  -n, --number integer     Number of lines to generate.
  -b, --bytes integer      Size of logs to generate (in bytes).
                           "bytes" will be ignored when "number" is set.
  -s, --sleep numeric      Sleep interval time between lines (in seconds). It does not actually sleep every log generation.
  -p, --split-by integer   Split the logs by this value in lines or bytes.
                           When "number" is set, it specifies the maximum number of lines for a log file.
                           When "bytes" is set, it specifies the maximum size of a log file.
  -w, --overwrite          [Warning] This will overwrite the existing file with new created logs.
`

var validFormats = []string{"apache_common", "apache_combined", "apache_error", "rfc3164"}
var validTypes = []string{"stdout", "log", "gz"}

// Option defines log generator options
type Option struct {
	Format    string
	Output    string
	Type      string
	Number    int
	Bytes     int
	Sleep     float64
	SplitBy   int
	Overwrite bool
}

func init() {
	pflag.Usage = printUsage
}

func printUsage() {
	fmt.Print(usage)
}

func errorExit(err error) {
	os.Stderr.WriteString(err.Error() + "\n")
	os.Exit(-1)
}

func defaultOptions() *Option {
	return &Option{
		Format:    "apache_common",
		Output:    "generated.log",
		Type:      "stdout",
		Number:    1000,
		Bytes:     0,
		Sleep:     0.0,
		SplitBy:   0,
		Overwrite: false,
	}
}

// ParseFormat validates the given format
func ParseFormat(format string) (string, error) {
	if !ContainsString(validFormats, format) {
		return "", fmt.Errorf("%s is not a valid format", format)
	}
	return format, nil
}

func ParseType(logType string) (string, error) {
	if !ContainsString(validTypes, logType) {
		return "", fmt.Errorf("%s is not a valid log type", logType)
	}
	return logType, nil
}

// ParseNumber validates the given number
func ParseNumber(lines int) (int, error) {
	if lines < 0 {
		return 0, errors.New("lines can not be negative")
	}
	return lines, nil
}

// ParseBytes validates the given bytes
func ParseBytes(bytes int) (int, error) {
	if bytes < 0 {
		return 0, errors.New("bytes can not be negative")
	}
	return bytes, nil
}

// ParseSleep validates the given sleep
func ParseSleep(sleep float64) (float64, error) {
	if sleep < 0 {
		return 0.0, errors.New("sleep can not be negative")
	}
	return sleep, nil
}

// ParseSplitBy validates the given split-by
func ParseSplitBy(splitBy int) (int, error) {
	if splitBy < 0 {
		return 0, errors.New("split-by can not be negative")
	}
	return splitBy, nil
}

// ParseOptions parses given parameters from command line
func ParseOptions() *Option {
	var err error

	opts := defaultOptions()

	help := pflag.BoolP("help", "h", false, "Show usage")
	format := pflag.StringP("format", "f", opts.Format, "Log format")
	output := pflag.StringP("output", "o", opts.Output, "Output filename. Path-like filename is allowed")
	logType := pflag.StringP("type", "t", opts.Type, "Log output type")
	number := pflag.IntP("number", "n", opts.Number, "Number of lines to generate")
	bytes := pflag.IntP("bytes", "b", opts.Bytes, "Size of logs to generate. (in bytes)")
	sleep := pflag.Float64P("sleep", "s", opts.Sleep, "Sleep interval time between lines. (in seconds)")
	splitBy := pflag.IntP("split", "p", opts.SplitBy, "Split by this value in lines or bytes")
	overwrite := pflag.BoolP("overwrite", "w", false, "To overwrite the existing file with new created logs")

	pflag.Parse()

	if *help {
		printUsage()
		os.Exit(0)
	}
	if opts.Format, err = ParseFormat(*format); err != nil {
		errorExit(err)
	}
	if opts.Type, err = ParseType(*logType); err != nil {
		errorExit(err)
	}
	if opts.Number, err = ParseNumber(*number); err != nil {
		errorExit(err)
	}
	if opts.Bytes, err = ParseBytes(*bytes); err != nil {
		errorExit(err)
	}
	if opts.Sleep, err = ParseSleep(*sleep); err != nil {
		errorExit(err)
	}
	if opts.SplitBy, err = ParseSplitBy(*splitBy); err != nil {
		errorExit(err)
	}
	opts.Output = *output
	opts.Overwrite = *overwrite
	return opts
}
