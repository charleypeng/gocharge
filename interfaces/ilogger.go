package gocharge

type ILogger interface {
	LogInfo(v any)
	LogError(v any)
	LogDebug(v any)
	LogCritical(v any)
}
