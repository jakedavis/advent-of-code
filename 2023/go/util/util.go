package util

import "runtime"

const (
	// Insane things go makes you worry about dot com
	WindowsNewline = "\r\n"
	UnixNewline    = "\n"
)

func Newline() string {
	if runtime.GOOS == "windows " {
		return WindowsNewline
	} else {
		return UnixNewline
	}
}
