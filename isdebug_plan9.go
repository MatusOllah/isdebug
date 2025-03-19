//go:build plan9

package isdebug

import "os"

func isDebug() bool {
	_, err := os.Stat("/dev/debug")
	return err == nil
}
