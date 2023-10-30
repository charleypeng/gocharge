package logging

import (
	"fmt"
	"gocharge/interfaces"
	"gocharge/logging/LoggerType"
)

type IStdLogger interface {
	gocharge.ILogger
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
	l.loggerType = LoggerType.Standard
	fmt.Println(v)
}
