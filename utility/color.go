package utility

import . "github.com/tntmeijs/gengo/mathematics"

// Represents an RGBA color
type Color struct {
	Red, Green, Blue, Alpha uint8
}

// Normalize the color and return it as a Vec3
func (c *Color) AsNormalizedVec3() Vec3 {
	x := ClampBetween(float64(c.Red)/255.0, 0.0, 1.0)
	y := ClampBetween(float64(c.Green)/255.0, 0.0, 1.0)
	z := ClampBetween(float64(c.Blue)/255.0, 0.0, 1.0)

	return Vec3{X: x, Y: y, Z: z}
}

// Convert a normalized Vec3 into a fully opaque color
func ColorFromNormalizedVec3(vector Vec3) Color {
	red := ClampBetween(vector.X*255.0, 0.0, 255.0)
	green := ClampBetween(vector.Y*255.0, 0.0, 255.0)
	blue := ClampBetween(vector.Z*255.0, 0.0, 255.0)

	return Color{Red: uint8(red), Green: uint8(green), Blue: uint8(blue), Alpha: 255}
}
