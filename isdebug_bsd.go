//go:build freebsd || openbsd || netbsd || dragonfly || aix

package isdebug

import "syscall"

func isDebug() bool {
	err := syscall.PtraceAttach(0)
	if err != nil {
		return true
	}
	defer syscall.PtraceDetach(0)
	return false
}
