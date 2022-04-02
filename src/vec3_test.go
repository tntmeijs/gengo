package main

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

func TestMultiply(t *testing.T) {
	a := Vec3{1.0, 2.0, 3.0}
	result := Mul(a, 10.0)

	if result.X != 10.0 || result.Y != 20.0 || result.Z != 30.0 {
		t.Fatalf("Multiplication failure: {1.0, 2.0, 3.0} * 10.0 != %v", result)
	}
}

func TestMultiplyWith(t *testing.T) {
	a := Vec3{1.0, 2.0, 3.0}
	a.Mul(10.0)

	if a.X != 10.0 || a.Y != 20.0 || a.Z != 30.0 {
		t.Fatalf("Multiplication failure: {1.0, 2.0, 3.0} * 10.0 != %v", a)
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
