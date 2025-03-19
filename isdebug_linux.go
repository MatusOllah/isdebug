//go:build linux || illumos || solaris || android

package isdebug

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isDebug() bool {
	f, err := os.Open("/proc/self/status")
	if err != nil {
		return false // assume no debugger if we can't read the file
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ln := scanner.Text()
		if strings.HasPrefix(ln, "TracerPid:") {
			var tracerPid int
			fmt.Sscanf(ln, "TracerPid:\t%d", &tracerPid)
			return tracerPid != 0
		}
	}
	return false
}
