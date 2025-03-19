//go:build darwin || ios

package isdebug

import (
	"syscall"
	"unsafe"
)

func isDebug() bool {
	var info int32
	size := uintptr(unsafe.Sizeof(info))

	// CTL_KERN (1), KERN_PROC (14), KERN_PROC_PID (1), current PID
	mib := [4]int32{1, 14, 1, int32(syscall.Getpid())}

	_, _, errno := syscall.Syscall6(syscall.SYS___SYSCTL, uintptr(unsafe.Pointer(&mib[0])), 4, uintptr(unsafe.Pointer(&info)), uintptr(unsafe.Pointer(&size)), 0, 0)
	if errno != 0 {
		return false // if sysctl fails, assume no debugger
	}

	const P_TRACED = 0x00000800
	return (info & P_TRACED) != 0
}
