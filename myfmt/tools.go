package myfmt

import (
	"strings"
)

func IsNullOrEmpty(str string, trimSpace ...bool) bool {
	if len(str) == 0 {
		return true
	}
	if len(trimSpace) == 0 {
		return false
	}
	if len(trimSpace) > 1 {
		panic("only 1 param should be set")
	}
	if trimSpace[0] == true {
		if len(strings.TrimSpace(str)) == 0 {
			return true
		}

	}
	return false
}
