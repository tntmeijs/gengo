package mathematics

import "math"

// Clamp a value between an upper and a lower bound
func ClampBetween(value float64, min float64, max float64) float64 {
	return math.Min(math.Max(min, value), max)
}
