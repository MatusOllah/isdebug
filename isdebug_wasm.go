//go:build wasm

package isdebug

func isDebug() bool {
	return false // no reliable way to check for debuggers in WASM
}
