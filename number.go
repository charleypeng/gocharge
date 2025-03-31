package gocharge

import "fmt"

type Number interface {
	~int64 | ~float64
}

// SumNumbers sums all the values in a map[K]V
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbersBy sums all the values in a map[K]V using a function to combine them
func SumNumbersBy[K comparable, V Number](m map[K]V, f func(V, V) V) V {
	var s V
	for _, v := range m {
		s = f(s, v)
	}
	return s
}

// Max returns the larger of two numbers
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of two numbers
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Clamp returns a value clamped to the inclusive range of min and max
func Clamp[T Number](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// IsInRange checks if a value is within the specified range (inclusive)
func IsInRange[T Number](value, min, max T) bool {
	return value >= min && value <= max
}

// Abs returns the absolute value of a number
func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Sign returns:
// -1 if x <  0
//
//	0 if x == 0
//	1 if x >  0
func Sign[T Number](x T) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

// IsZero checks if a number is zero
func IsZero[T Number](x T) bool {
	return x == 0
}

// IsPositive checks if a number is greater than zero
func IsPositive[T Number](x T) bool {
	return x > 0
}

// IsNegative checks if a number is less than zero
func IsNegative[T Number](x T) bool {
	return x < 0
}

// Between checks if a number is strictly between min and max (exclusive)
func Between[T Number](value, min, max T) bool {
	return value > min && value < max
}

// Compare compares two numbers and returns:
// -1 if a < b
//
//	0 if a == b
//	1 if a > b
func Compare[T Number](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// DivRem performs integer division and returns both quotient and remainder
// Note: This function is only available for int64 type
func DivRem(a, b int64) (quotient, remainder int64) {
	quotient = a / b
	remainder = a % b
	return
}

// Round returns the nearest integer value as float64
// Rounds half away from zero
func Round(x float64) float64 {
	if x < 0 {
		return float64(int64(x - 0.5))
	}
	return float64(int64(x + 0.5))
}

// Floor returns the largest integer value less than or equal to x
func Floor(x float64) float64 {
	return float64(int64(x))
}

// Ceiling returns the smallest integer value greater than or equal to x
func Ceiling(x float64) float64 {
	if float64(int64(x)) == x {
		return x
	}
	if x < 0 {
		return float64(int64(x))
	}
	return float64(int64(x) + 1)
}

// ParseInt64 attempts to parse a string to int64
func ParseInt64(s string) (int64, bool) {
	var result int64
	_, err := fmt.Sscanf(s, "%d", &result)
	return result, err == nil
}

// ParseFloat64 attempts to parse a string to float64
func ParseFloat64(s string) (float64, bool) {
	var result float64
	_, err := fmt.Sscanf(s, "%f", &result)
	return result, err == nil
}
