// gocharge is a package for users to use go lang in a safe and convenient way
//
// copyright charleypeng(peng lei)
package gocharge

import (
	"fmt"
)

func Version() {
	fmt.Printf("v0.0.8\n")
	few := MyString("test string")
	few.IsNullOrWhiteSpace()
}
