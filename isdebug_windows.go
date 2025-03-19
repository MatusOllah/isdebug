//go:build windows

package isdebug

import "golang.org/x/sys/windows"

var (
	kernel32          = windows.NewLazySystemDLL("kernel32.dll")
	isDebuggerPresent = kernel32.NewProc("IsDebuggerPresent")
)

func isDebug() bool {
	value, _, _ := isDebuggerPresent.Call()
	return value != 0
}
