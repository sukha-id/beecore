//go:build !linux
// +build !linux

package dailylogger

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
