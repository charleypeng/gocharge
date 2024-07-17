package gocharge

import (
	"fmt"
)

type IStdLogger interface {
	ILogger
	LetMeGo() string
}

type Logger struct {
	std        string
	loggerType int
}

func (l Logger) LetMeGo() string {
	return "fwdfw"
}

func (l Logger) LogError(v any) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) LogDebug(v any) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) LogCritical(v any) {
	//TODO implement me
	panic("implement me")
}

func (l Logger) LogInfo(v any) {
	fmt.Println(v)
}
