package mathematics

import "testing"

const epsilon = 0.0001

func TestAdd(t *testing.T) {
	a := Vec3{1.0, 2.0, 3.0}
	b := Vec3{10.0, 20.0, 30.0}
	result := Add(a, b)

	if result.X != 11.0 || result.Y != 22.0 || result.Z != 33.0 {
		t.Fatalf("Addition failure: {1.0, 2.0, 3.0} + {10.0, 20.0, 30.0} != %v", result)
	}
}

func TestAddTo(t *testing.T) {
	a := Vec3{1.0, 2.0, 3.0}
	b := Vec3{10.0, 20.0, 30.0}
	a.Add(b)

	if a.X != 11.0 || a.Y != 22.0 || a.Z != 33.0 {
		t.Fatalf("Addition failure: {1.0, 2.0, 3.0} + {10.0, 20.0, 30.0} != %v", a)
	}
}

func TestAddAll(t *testing.T) {
	a := Vec3{1.0, 2.0, 3.0}
	b := Vec3{2.0, 3.0, 4.0}
	c := Vec3{3.0, 4.0, 5.0}
	result := AddAll(a, b, c)

	if result.X != 6.0 || result.Y != 9 || result.Z != 12.0 {
		t.Fatalf("Addition failure: %v + %v + %v != %v", a, b, c, result)
	}
}

func TestSubtract(t *testing.T) {
	a := Vec3{1.0, 2.0, 3.0}
	b := Vec3{10.0, 20.0, 30.0}
	result := Sub(a, b)

	if result.X != -9.0 || result.Y != -18.0 || result.Z != -27.0 {
		t.Fatalf("Subtraction failure: {1.0, 2.0, 3.0} - {10.0, 20.0, 30.0} != %v", result)
	}
}

func TestSubtractFrom(t *testing.T) {
	a := Vec3{1.0, 2.0, 3.0}
	b := Vec3{10.0, 20.0, 30.0}
	a.Sub(b)

	if a.X != -9.0 || a.Y != -18.0 || a.Z != -27.0 {
		t.Fatalf("Subtraction failure: {1.0, 2.0, 3.0} + {10.0, 20.0, 30.0} != %v", a)
	}
}

func TestMultiplyScalar(t *testing.T) {
	a := Vec3{1.0, 2.0, 3.0}
	result := MultiplyScalar(a, 10.0)

	if result.X != 10.0 || result.Y != 20.0 || result.Z != 30.0 {
		t.Fatalf("Multiplication failure: %v * 10.0 != %v", a, result)
	}
}

func TestMultiplyWith(t *testing.T) {
	a := Vec3{1.0, 2.0, 3.0}
	a.MultiplyWith(10.0)

	if a.X != 10.0 || a.Y != 20.0 || a.Z != 30.0 {
		t.Fatalf("Multiplication failure: {1.0, 2.0, 3.0} * 10.0 != %v", a)
	}
}

func TestMultiply(t *testing.T) {
	a := Vec3{1.0, 2.0, 3.0}
	b := Vec3{1.0, 2.0, 3.0}
	c := Multiply(a, b)

	if c.X != 1.0 || c.Y != 4.0 || c.Z != 9.0 {
		t.Fatalf("Multiplication failure: %v * %v != %v", a, b, c)
	}
}

func TestDot(t *testing.T) {
	a := Vec3{-6.0, 8.0, 0.0}
	b := Vec3{5.0, 12.0, 0.0}
	scalar := Dot(a, b)

	if scalar != 66.0 {
		t.Fatalf("Dot product failure: expected 66.0 but got %f", scalar)
	}
}

func TestNegate(t *testing.T) {
	a := Vec3{2.0, 3.0, 4.0}
	result := Negate(a)

	if result.X != -2.0 || result.Y != -3.0 || result.Z != -4.0 {
		t.Fatalf("Negate failure: -%v != %v", a, result)
	}
}

func TestReflect(t *testing.T) {
	incident := Vec3{1.0, -1.0, 0.0}
	normal := Vec3{0.0, 1.0, 0.0}
	reflection := Reflect(incident, normal)

	if reflection.X != 1.0 || reflection.Y != 1.0 || reflection.Z != 0.0 {
		t.Fatalf("Reflect failure: %v relfected against %v != %v", incident, normal, reflection)
	}
}

func TestMagnitude(t *testing.T) {
	a := Vec3{3.0, 4.0, 0.0}
	magnitude := a.Magnitude()

	if magnitude != 25.0 {
		t.Fatalf("Magnitude failure: expected 25.0 but got %f", magnitude)
	}
}

func TestMagnitudeSqrt(t *testing.T) {
	a := Vec3{3.0, 4.0, 0.0}
	magnitude := a.MagnitudeSqrt()

	if magnitude != 5.0 {
		t.Fatalf("Magnitude square-root failure: expected 5.0 but got %f", magnitude)
	}
}

func TestNormalize(t *testing.T) {
	a := Vec3{3.0, 4.0, 5.0}
	normal := Normalize(a)
	length := normal.MagnitudeSqrt()

	if length-1.0 > epsilon {
		t.Fatalf("Normalize failure: expected 1.0 but got %f", length)
	}
}
