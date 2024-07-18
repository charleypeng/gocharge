package gocharge

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
	_s := s.ToString()
	return IsStringNull(&_s, trimSpace...)
}

// IsNullOrWhiteSpace returns a bool to make sure if it is null or whitespace
func (s *MyString) IsNullOrWhiteSpace() bool {
	return s.IsNullOrEmpty(true)
}

func IsStringNull(s *string, trimSpace ...bool) bool {
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
			if len(strings.TrimSpace(*s)) == 0 {
				return true
			}
		}
	}

	return false
}

func IsStringNullOrWhiteSpace(s *string) bool {
	return IsStringNull(s, true)
}
