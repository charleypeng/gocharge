package gocharge

import (
	"gocharge/logging"
	"testing"
)

func TestVersion(t *testing.T) {
	Version()
	var l = logging.IStdLogger(logging.Logger{})
	l.LetMeGo()
	//myfmt.(&l)
}
