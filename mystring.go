package gocharge

import (
	"fmt"
	"strings"
	"unicode"
)

// MyString is a  type for csharp like string
type MyString string

// ToString converts the MyString to a standard Go string.
// It returns the underlying string value of the MyString type.
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

// Contains determines whether a substring occurs within this string
func (s *MyString) Contains(value string) bool {
	return strings.Contains(s.ToString(), value)
}

// StartsWith determines whether the beginning of this string matches the specified string
func (s *MyString) StartsWith(value string) bool {
	return strings.HasPrefix(s.ToString(), value)
}

// EndsWith determines whether the end of this string matches the specified string
func (s *MyString) EndsWith(value string) bool {
	return strings.HasSuffix(s.ToString(), value)
}

// ToLower returns a copy of this string converted to lowercase
func (s *MyString) ToLower() MyString {
	return MyString(strings.ToLower(s.ToString()))
}

// ToUpper returns a copy of this string converted to uppercase
func (s *MyString) ToUpper() MyString {
	return MyString(strings.ToUpper(s.ToString()))
}

// Trim removes all leading and trailing white-space characters from the current string
func (s *MyString) Trim(cutset ...string) MyString {
	if len(cutset) == 0 {
		return MyString(strings.TrimSpace(s.ToString()))
	}
	return MyString(strings.Trim(s.ToString(), cutset[0]))
}

// TrimStart removes all leading white-space characters from the current string
func (s *MyString) TrimStart(cutset ...string) MyString {
	if len(cutset) == 0 {
		return MyString(strings.TrimLeftFunc(s.ToString(), unicode.IsSpace))
	}
	return MyString(strings.TrimLeft(s.ToString(), cutset[0]))
}

// TrimEnd removes all trailing white-space characters from the current string
func (s *MyString) TrimEnd(cutset ...string) MyString {
	if len(cutset) == 0 {
		return MyString(strings.TrimRightFunc(s.ToString(), unicode.IsSpace))
	}
	return MyString(strings.TrimRight(s.ToString(), cutset[0]))
}

// Replace replaces all occurrences of a specified string in the current instance with another specified string
func (s *MyString) Replace(oldValue, newValue string) MyString {
	return MyString(strings.ReplaceAll(s.ToString(), oldValue, newValue))
}

// Split splits a string into substrings based on the separator
func (s *MyString) Split(separator string) []string {
	return strings.Split(s.ToString(), separator)
}

// Substring retrieves a substring from this instance
func (s *MyString) Substring(start int, length ...int) MyString {
	str := s.ToString()
	if len(length) == 0 {
		if start < 0 || start > len(str) {
			panic("Index out of range")
		}
		return MyString(str[start:])
	}

	end := start + length[0]
	if start < 0 || end > len(str) || length[0] < 0 {
		panic("Index out of range")
	}
	return MyString(str[start:end])
}

// IndexOf returns the zero-based index of the first occurrence of the specified string in this instance
func (s *MyString) IndexOf(value string) int {
	return strings.Index(s.ToString(), value)
}

// LastIndexOf returns the zero-based index of the last occurrence of the specified string in this instance
func (s *MyString) LastIndexOf(value string) int {
	return strings.LastIndex(s.ToString(), value)
}

// PadLeft returns a new string that right-aligns the characters in this instance by padding them on the left with a specified character
func (s *MyString) PadLeft(totalWidth int, paddingChar ...rune) MyString {
	str := s.ToString()
	if len(str) >= totalWidth {
		return *s
	}

	padding := ' '
	if len(paddingChar) > 0 {
		padding = paddingChar[0]
	}

	pad := strings.Repeat(string(padding), totalWidth-len(str))
	return MyString(pad + str)
}

// PadRight returns a new string that left-aligns the characters in this string by padding them on the right with a specified character
func (s *MyString) PadRight(totalWidth int, paddingChar ...rune) MyString {
	str := s.ToString()
	if len(str) >= totalWidth {
		return *s
	}

	padding := ' '
	if len(paddingChar) > 0 {
		padding = paddingChar[0]
	}

	pad := strings.Repeat(string(padding), totalWidth-len(str))
	return MyString(str + pad)
}

// Length returns the length of the string
func (s *MyString) Length() int {
	return len(s.ToString())
}

// Join concatenates the elements of a string array, using the specified separator between each element
func Join(separator string, values []string) MyString {
	return MyString(strings.Join(values, separator))
}

// Format formats a string using the specified format string and arguments
func Format(format string, args ...interface{}) MyString {
	return MyString(fmt.Sprintf(format, args...))
}

// CompareTo compares this instance with another string and returns an integer that indicates their relative position in the sort order
func (s *MyString) CompareTo(value string) int {
	return strings.Compare(s.ToString(), value)
}

// Equals determines whether two string instances are equal
func (s *MyString) Equals(value string) bool {
	return s.ToString() == value
}

// Clone creates a new MyString instance with the same value as this instance
func (s *MyString) Clone() MyString {
	return MyString(s.ToString())
}
