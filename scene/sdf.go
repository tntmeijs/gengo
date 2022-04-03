package scene

import (
	"math"

	. "github.com/tntmeijs/gengo/mathematics"
)

// Sphere centered on the origin
func SphereSDF(point Vec3, radius float64) float64 {
	return math.Sqrt((point.X*point.X)+(point.Y*point.Y)+(point.Z*point.Z)) - radius
}

// Mandelbulb fractal
// Adapted from: http://blog.hvidtfeldts.net/index.php/2011/09/distance-estimated-3d-fractals-v-the-mandelbulb-different-de-approximations/
func MandelbulbSDF(point Vec3, iterations int, power int, bailout float64) float64 {
	z := point
	dr := 1.0
	r := 0.0
	powerFloat := float64(power)

	for i := 0; i < iterations; i++ {
		r = z.MagnitudeSqrt()
		if r > bailout {
			break
		}

		// Convert to polar coordinates
		theta := math.Acos(z.Z / r)
		phi := math.Atan2(z.Y, z.X)
		dr = math.Pow(r, powerFloat-1.0)*powerFloat*dr + 1.0

		// Scale and rotate the point
		zr := math.Pow(r, powerFloat)
		theta = theta * powerFloat
		phi = phi * powerFloat

		// Convert back to cartesian coordinates
		z = MultiplyScalar(Vec3{X: math.Sin(theta) * math.Cos(phi), Y: math.Sin(phi) * math.Sin(theta), Z: math.Cos(theta)}, zr)
		z.Add(point)
	}

	return 0.5 * math.Log(r) * r / dr
}
