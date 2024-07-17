package types

import (
	"strings"
)

// MyString is a  type for csharp like string
type MyString string

func (s *MyString) ToString() string {
	return string(*s)
}

// IsNullOrEmpty returns a bool to make sure if it is null or empty
func (s *MyString) IsNullOrEmpty(trimSpace ...bool) bool {
	var _s = s.ToString()

	if s == nil {
		return true
	}

	if len(*s) == 0 {
		return true
	}

	if len(trimSpace) > 1 {
		panic("only 1 param should be set")
	}

	if len(trimSpace) == 1 {
		if trimSpace[0] {
			if len(strings.TrimSpace(_s)) == 0 {
				return true
			}
		}
	}

	return false
}

// IsNullOrWhiteSpace returns a bool to make sure if it is null and whitespace
func (s *MyString) IsNullOrWhiteSpace() bool {
	var _s = s.ToString()
	if len(*s) == 0 {
		return true
	}

	if len(strings.TrimSpace(_s)) == 0 {
		return true
	}
	return false
}
