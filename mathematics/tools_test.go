package mathematics

import "testing"

func TestClampBetween(t *testing.T) {
	result := ClampBetween(50.0, 0.0, 100.0)

	if result != 50.0 {
		t.Fatalf("ClampBetween failed to clamp the value within bounds")
	}
}

func TestClampBetweenUpperBound(t *testing.T) {
	result := ClampBetween(150.0, 0.0, 100.0)

	if result != 100.0 {
		t.Fatalf("ClampBetween failed to clamp the value within bounds")
	}
}

func TestClampBetweenUpperBoundExact(t *testing.T) {
	result := ClampBetween(100.0, 0.0, 100.0)

	if result != 100.0 {
		t.Fatalf("ClampBetween failed to clamp the value within bounds")
	}
}

func TestClampBetweenLowerBoundExact(t *testing.T) {
	result := ClampBetween(0.0, 0.0, 100.0)

	if result != 0.0 {
		t.Fatalf("ClampBetween failed to clamp the value within bounds")
	}
}
