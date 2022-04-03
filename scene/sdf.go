package scene

import (
	"math"

	. "github.com/tntmeijs/gengo/mathematics"
)

// Sphere centered on the origin
func SphereSDF(point Vec3, radius float64) float64 {
	return math.Sqrt((point.X*point.X)+(point.Y*point.Y)+(point.Z*point.Z)) - radius
}
