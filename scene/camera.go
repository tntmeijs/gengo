package scene

import (
	"math"

	. "github.com/tntmeijs/gengo/mathematics"
)

// A camera from which a scene can be rendered
type Camera struct {
	Position, direction Vec3
	nearPlane, farPlane float64
}

// Create a new camera with the following properties:
//	Position  : position of the camera in space
//	Focus     : point in space the camera is pointed towards
//	NearPlane : camera near plane (how close an object can be before it gets clipped)
//	FarPlane  : camera far plane (how far an object can be before it gets clipped)
func NewCamera(position Vec3, focus Vec3, nearPlane float64, farPlane float64) Camera {
	return Camera{position, Normalize(Sub(focus, position)), nearPlane, farPlane}
}

// Get the direction the camera is pointing into
func (c *Camera) GetDirection() Vec3 {
	return c.direction
}

// Set the point in space where the camera is pointed towards
func (c *Camera) SetFocusPoint(focus Vec3) {
	c.direction = Normalize(Sub(focus, c.Position))
}

// Generate a new ray at the center of the specified pixel
func (c *Camera) GenerateRayForPixelCenter(pixelX int, pixelY int, resolutionX int, resolutionY int) Ray {
	return c.GenerateRayForPixelWithOffset(pixelX, pixelY, 0.5, 0.5, resolutionX, resolutionY)
}

// Generate a new ray at the specified pixel with a relative normalized offset within the pixel.
// This allows for techniques like as anti-aliasing as it relies on multiple rays being cast through a pixel
// with varying offsets.
func (c *Camera) GenerateRayForPixelWithOffset(pixelX int, pixelY int, offsetX float64, offsetY float64, resolutionX int, resolutionY int) Ray {
	// Ensure the offsets never exceed the [0.0, 1.0]
	offsetX = math.Max(math.Min(1.0, offsetX), 0.0)
	offsetY = math.Max(math.Min(1.0, offsetY), 0.0)

	// Calculate screen-space coordinates and make sure that each ray goes through the center of a pixel
	screenX := (float64(pixelX) + offsetX) / float64(resolutionX)
	screenY := (float64(pixelY) + offsetY) / float64(resolutionY)

	// Convert screen-space coordinates from [0.0, 1.0] into [-1.0, 1.0]
	screenX = (screenX * 2.0) - 1.0
	screenY = (screenY * 2.0) - 1.0

	// Compensate the aspect ratio to ensure that the results do not look stretched
	aspectRatio := float64(resolutionX) / float64(resolutionY)

	// Offset the ray to make it trace through the imaginary image plane
	direction := c.GetDirection()
	direction.X += screenX * aspectRatio
	direction.Y -= screenY

	return Ray{Origin: Add(c.Position, MultiplyScalar(direction, c.nearPlane)), Direction: direction}
}

// Cast a ray into the scene and march towards the surface until an intersection is found, or until the
// ray passes the far plane of the camera
func (c *Camera) MarchAlongRay(ray Ray, scene Scene, stepSize float64) (bool, SurfaceHitInfo) {
	for distance := 0.0; distance < c.farPlane; distance += stepSize {
		pointInSpace := Add(ray.Origin, MultiplyScalar(ray.Direction, distance))

		if scene.DoesPointIntersectSurface(pointInSpace) {
			// Surface intersection found
			return true, scene.GetIntersectionPointSurfaceHitInfo(pointInSpace, distance)
		}
	}

	// No surface intersection found
	return false, SurfaceHitInfo{}
}
