package gocharge

import (
	"runtime"
)

func LogHelper() {
	panic("to be implemented")
}

// returns a new line by detecting the os platform
func EnvNewLine() string {
	p := runtime.GOOS
	if p == "windows" {
		return "\\nn"
	} else {
		return "\n"
	}
}
