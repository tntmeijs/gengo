package main

import (
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

var camera = NewCamera(Vec3{X: 0.0, Y: 0.0, Z: -1.5}, Vec3{X: 0.0, Y: 0.0, Z: 0.0}, NearPlane, FarPlane)
var scene = NewScene(sceneSDF)

// Entire scene represented as a signed distance function
func sceneSDF(point Vec3) float64 {
	return SphereSDF(point, 1)
}

// Application entry point
func main() {
	image := NewPngImage(ImageResolutionX, ImageResolutionY, ImageFileName)

	for y := 0; y < ImageResolutionY; y++ {
		for x := 0; x < ImageResolutionX; x++ {
			ray := camera.GenerateRayForPixelCenter(x, y, ImageResolutionX, ImageResolutionY)
			pixelColor := Color{Red: 0, Green: 0, Blue: 0, Alpha: 0}

			didHit, hitInfo := camera.MarchAlongRay(ray, scene, StepSize)

			if didHit {
				pixelColor = Color{Red: uint8(hitInfo.RayLength * 200.0), Green: 0, Blue: 0, Alpha: 255}
			}

			image.SetPixelColor(x, y, pixelColor)
		}
	}

	image.WritePngToFile()
}
