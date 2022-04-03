package utility

// Represents an RGBA color
type Color struct {
	Red, Green, Blue, Alpha uint8
}

// Sky blue: (135, 206, 235)
func SkyBlue() Color {
	return Color{135, 206, 235, 255}
}
