package types

import (
	"gocharge/interfaces"
)

type Number interface {
	~int64 | ~float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func Log[T gocharge.ILogger](lover T) {
	lover.LogInfo("wjkk")
}
