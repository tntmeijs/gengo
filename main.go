package main

import (
	"math"

	. "github.com/tntmeijs/gengo/mathematics"
	. "github.com/tntmeijs/gengo/scene"
	. "github.com/tntmeijs/gengo/utility"
)

// =============================================================================================================================
// =============================================================================================================================
// =============================================================================================================================

// Output image information
const ImageResolutionX = 640
const ImageResolutionY = 360
const ImageFileName = "output.png"

// Camera constants
const RayStepSize = 0.01
const CameraNearPlane = 0.001
const CameraFarPlane = 25.0

// Light parameters
const AmbientStrength = 0.25
const SpecularStrength = 0.5
const SpecularShininess = 32

// Camera transformation
var cameraPosition = Vec3{X: 0.0, Y: 0, Z: -1.525}
var cameraLookAt = Vec3{X: 0.0, Y: 0.0, Z: 0.0}

// Lights
var ambientLightPosition = Vec3{X: -10.0, Y: 10.0, Z: -10.0}

// Colors
var ambientColor = Color{Red: 255, Green: 255, Blue: 255, Alpha: 255}
var surfaceColor = Color{Red: 26, Green: 188, Blue: 156, Alpha: 255}

// =============================================================================================================================
// =============================================================================================================================
// =============================================================================================================================

// Entire scene represented as a signed distance function
func sceneSDF(point Vec3) float64 {
	return MandelbulbSDF(point, 7, 8, 5.0)
}

// Simple Blinn-Phong lighting model
//
// Reference: https://learnopengl.com/Advanced-Lighting/Advanced-Lighting
func calculatePixelColor(surfaceInfo SurfaceHitInfo, camera Camera) Color {
	ambientLightDirection := Normalize(Sub(ambientLightPosition, surfaceInfo.Point))
	viewDirection := Normalize(Sub(camera.Position, surfaceInfo.Point))
	halfwayDirection := Normalize(Add(ambientLightDirection, viewDirection))

	ambient := MultiplyScalar(ambientColor.AsNormalizedVec3(), AmbientStrength)
	diffuse := MultiplyScalar(ambientColor.AsNormalizedVec3(), math.Max(Dot(surfaceInfo.Normal, ambientLightDirection), 0.0))
	specular := MultiplyScalar(ambientColor.AsNormalizedVec3(), math.Pow(math.Max(Dot(surfaceInfo.Normal, halfwayDirection), 0.0), SpecularShininess)*SpecularStrength)

	lightColor := AddAll(ambient, diffuse, specular)
	outputColor := Multiply(lightColor, surfaceColor.AsNormalizedVec3())

	return ColorFromNormalizedVec3(outputColor)
}

// Application entry point
func main() {
	scene := NewScene(sceneSDF)
	image := NewPngImage(ImageResolutionX, ImageResolutionY, ImageFileName)
	camera := NewCamera(cameraPosition, cameraLookAt, CameraNearPlane, CameraFarPlane)

	for y := 0; y < ImageResolutionY; y++ {
		for x := 0; x < ImageResolutionX; x++ {
			ray := camera.GenerateRayForPixelCenter(x, y, ImageResolutionX, ImageResolutionY)
			pixelColor := Color{Red: 0, Green: 0, Blue: 0, Alpha: 0}

			didHit, hitInfo := camera.MarchAlongRay(ray, scene, RayStepSize)

			if didHit {
				pixelColor = calculatePixelColor(hitInfo, camera)
			}

			image.SetPixelColor(x, y, pixelColor)
		}
	}

	image.WritePngToFile()
}
