// +build !windows

package flog

import (
	"errors"
	"os"
	"path/filepath"
	"syscall"
)

// Run checks overwrite flag and generates logs with given options
func Run(option *Option) error {
	logDir := filepath.Dir(option.Output)
	oldMask := syscall.Umask(0000)
	if err := os.MkdirAll(logDir, 0766); err != nil {
		return err
	}
	syscall.Umask(oldMask)
	if _, err := os.Stat(option.Output); err == nil && !option.Overwrite {
		return errors.New(option.Output + " already exists. You can overwrite with -w option")
	}
	return Generate(option)
}
