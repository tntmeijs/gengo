package mathematics

import "math"

type Vec3 struct {
	X, Y, Z float64
}

// Add a vector to this vector's values
func (v *Vec3) Add(other Vec3) {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}

// Add two vectors together
func Add(a Vec3, b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Add all vectors together
func AddAll(vectors ...Vec3) Vec3 {
	result := Vec3{}

	for _, vector := range vectors {
		result.Add(vector)
	}

	return result
}

// Subtract a vector from this vector's values
func (v *Vec3) Sub(other Vec3) {
	v.X -= other.X
	v.Y -= other.Y
	v.Z -= other.Z
}

// Subtract two vectors from each other
func Sub(a Vec3, b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Multiply a vector with a scalar value
func (v *Vec3) MultiplyWith(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
}

// Multiply a vector with a scalar value
func MultiplyScalar(a Vec3, scalar float64) Vec3 {
	return Vec3{a.X * scalar, a.Y * scalar, a.Z * scalar}
}

// Multiply a vector with another vector
func Multiply(a Vec3, b Vec3) Vec3 {
	return Vec3{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

// Calculate the dot product scalar value of two vectors
func Dot(a Vec3, b Vec3) float64 {
	return (a.X * b.X) + (a.Y * b.Y) + (a.Z * b.Z)
}

// Returns a negated copy of the input vector
func Negate(a Vec3) Vec3 {
	return Vec3{-a.X, -a.Y, -a.Z}
}

// Reflect a vector
//
// Reference: https://www.khronos.org/registry/OpenGL-Refpages/gl4/html/reflect.xhtml
func Reflect(incident Vec3, normal Vec3) Vec3 {
	scalarMultiplyDotMultiplyNormal := MultiplyScalar(normal, 2.0*Dot(normal, incident))

	reflectX := incident.X - scalarMultiplyDotMultiplyNormal.X
	reflectY := incident.Y - scalarMultiplyDotMultiplyNormal.Y
	reflectZ := incident.Z - scalarMultiplyDotMultiplyNormal.Z

	return Vec3{reflectX, reflectY, reflectZ}
}

// Calculate the non-squared magnitude of a vector.
//
// This function calculates the Pythagorean Theorem but without squaring the result.
//
// If the magnitude is only used to compare vectors,
// it is much quicker to use the non-squared magnitude
// as it will save a square-root operation.
func (v *Vec3) Magnitude() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

// Calculate the squared magnitude of a vector
//
// This function returns the result of the Pythagorean Theorem
func (v *Vec3) MagnitudeSqrt() float64 {
	return math.Sqrt(v.Magnitude())
}

// Normalize a vector
func Normalize(v Vec3) Vec3 {
	magnitude := v.MagnitudeSqrt()
	v.X /= magnitude
	v.Y /= magnitude
	v.Z /= magnitude
	return v
}
