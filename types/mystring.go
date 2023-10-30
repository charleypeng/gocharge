package types

import (
	"fmt"
	"strings"
)

// MyString is a test for csharp like string
type MyString string

func (s *MyString) ToString() string {
	return fmt.Sprintf("%v", *s)
}

func (s *MyString) IsNullOrEmpty(trimSpace ...bool) bool {
	var _s = s.ToString()

	if s == nil {
		return true
	}
	if len(*s) == 0 {
		return true
	}
	if len(trimSpace) == 0 {
		return false
	}
	if len(trimSpace) > 1 {
		panic("only 1 param should be set")
	}
	if trimSpace[0] == true {
		if len(strings.TrimSpace(_s)) == 0 {
			fmt.Println("trim hitted")
			return true
		}

	}
	return false
}

func (s *MyString) IsNullOrEmptySpace() bool {
	var _s = s.ToString()
	if len(*s) == 0 {
		fmt.Println("null hitted")
		return true
	}

	if len(strings.TrimSpace(_s)) == 0 {
		fmt.Println("trim hitted")
		return true
	}
	return false
}
