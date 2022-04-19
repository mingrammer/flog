package flog

import (
	"errors"
	"os"
	"path/filepath"
)

// Run checks overwrite flag and generates logs with given options
func Run(option *Option) error {
	logDir := filepath.Dir(option.Output)
	if err := os.MkdirAll(logDir, 0766); err != nil {
		return err
	}
	if _, err := os.Stat(option.Output); err == nil && !option.Overwrite {
		return errors.New(option.Output + " already exists. You can overwrite with -w option")
	}
	return Generate(option)
}
