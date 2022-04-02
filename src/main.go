package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
)

type sdf func(point Vec3) bool

type HitInfo struct {
	Point, Normal Vec3
	Distance      float64
}

type Color struct {
	Red, Green, Blue, Alpha uint8
}

type Ray struct {
	Origin, Direction Vec3
}

const Width = 640
const Height = 360
const ComponentsPerPixel = 4
const StepSize = 0.01
const NearPlane = 0.001
const FarPlane = 25.0
const Epsilon = 0.001

var cameraPosition = Vec3{0.0, 5.0, 0.0}
var cameraFocusPoint = Vec3{0.0, 0.0, 15.0}
var sky = Color{135, 206, 235, 255}

// setPixel is used to assign a pixel specific red, green, and blue components.
// Each pixel is assumed to use the RGBA uint8 color format. A color can be defined
// by specifying a value between 0 and 255 (inclusive) for each color component.
func setPixel(x int, y int, color Color, pixels *[]uint8) {
	firstColorIndex := (x * ComponentsPerPixel) + (Width * y * ComponentsPerPixel)

	(*pixels)[firstColorIndex] = color.Red
	(*pixels)[firstColorIndex+1] = color.Green
	(*pixels)[firstColorIndex+2] = color.Blue
	(*pixels)[firstColorIndex+3] = color.Alpha
}

func marchRay(ray Ray, sdf sdf) (bool, HitInfo) {
	for distance := NearPlane; distance < FarPlane; distance += StepSize {
		// Tip of the ray
		pointInSpace := Add(ray.Origin, Mul(ray.Direction, distance))

		// Evaluate the surface
		if sdf(pointInSpace) {
			// TODO: Calculate the surface normal by sampling points next to the hit location
			return true, HitInfo{Point: pointInSpace, Normal: Vec3{}, Distance: distance}
		}
	}

	// No surface intersection found
	return false, HitInfo{}
}

func generateRayForPixel(x int, y int) Ray {
	// Calculate screen-space coordinates and make sure that each ray goes through the center of a pixel
	screenX := (float64(x) + 0.5) / Width
	screenY := (float64(y) + 0.5) / Height

	// Convert screen-space coordinates from 0.0 --> 1.0 into -1.0 --> 1.0
	screenX = (screenX * 2.0) - 1.0
	screenY = (screenY * 2.0) - 1.0

	// Offset the ray to make it trace through the imaginary image plane
	direction := Normalize(Sub(cameraFocusPoint, cameraPosition))
	direction.X += screenX
	direction.Y -= screenY

	return Ray{cameraPosition, direction}
}

func doesPointIntersectSurface(point Vec3) bool {
	return point.Y < math.Sin(point.X)*math.Sin(point.Z)
}

func main() {
	// Create a blank image
	data := image.NewRGBA(image.Rect(0, 0, Width, Height))

	// Set all pixels to baby blue
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			ray := generateRayForPixel(x, y)
			pixelColor := sky

			didHit, hitInfo := marchRay(ray, doesPointIntersectSurface)

			if didHit {
				pixelColor = Color{uint8(hitInfo.Distance * 235.0), 164, 135, 255}
			}

			setPixel(x, y, pixelColor, &data.Pix)
		}
	}

	// Create a new file
	file, error := os.Create("output.png")
	if error != nil {
		fmt.Println("Error while creating output image: ", error)
	}

	// Write the image data as a .png file
	png.Encode(file, data)

	// Close the file handle
	file.Close()
}
