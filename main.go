package main

import (
	"math"

	. "github.com/tntmeijs/gengo/mathematics"
	. "github.com/tntmeijs/gengo/scene"
	. "github.com/tntmeijs/gengo/utility"
)

const ImageResolutionX = 640
const ImageResolutionY = 360
const ImageFileName = "output.png"
const StepSize = 0.01
const NearPlane = 0.001
const FarPlane = 25.0

var camera = NewCamera(Vec3{X: 0.0, Y: 5.0, Z: 0.0}, Vec3{X: 0.0, Y: 0.0, Z: 15.0}, NearPlane, FarPlane)
var scene = NewScene(sceneSDF)
var sky = SkyBlue()

// Entire scene represented as a signed distance function
func sceneSDF(point Vec3) float64 {
	return point.Y - math.Sin(point.X)*math.Sin(point.Z)
}

// Application entry point
func main() {
	image := NewPngImage(ImageResolutionX, ImageResolutionY, ImageFileName)

	for y := 0; y < ImageResolutionY; y++ {
		for x := 0; x < ImageResolutionX; x++ {
			ray := camera.GenerateRayForPixelCenter(x, y, ImageResolutionX, ImageResolutionY)
			pixelColor := sky

			didHit, hitInfo := camera.MarchAlongRay(ray, scene, StepSize)

			if didHit {
				pixelColor = Color{Red: uint8(hitInfo.Normal.X * 235.0), Green: uint8(hitInfo.Normal.Y * 164.0), Blue: uint8(hitInfo.Normal.Z * 135.0), Alpha: 255}
			}

			image.SetPixelColor(x, y, pixelColor)
		}
	}

	image.WritePngToFile()
}
