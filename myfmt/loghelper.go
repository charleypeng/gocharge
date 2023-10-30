package myfmt

import "gocharge/logging"

func LogError(logger *logging.IStdLogger) {
	if logger == nil {
		return
	}
	(*logger).LogError("ewe")
}
