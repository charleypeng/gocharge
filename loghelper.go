package gocharge

func LogError(logger *IStdLogger) {
	if logger == nil {
		return
	}
	(*logger).LogError("ewe")
}
