package gocharge

import (
	"gocharge/logging"
	"gocharge/types"
	"testing"
)

func TestVersion(t *testing.T) {
	var l = logging.IStdLogger(logging.Logger{})
	l.LetMeGo()
	var we types.MyString
	we = "3423"

	if we.IsNullOrEmpty(true) {
		println("wow")
	} else {
		println("now!")
	}
	//myfmt.(&l)
}
