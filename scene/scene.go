package scene

import . "github.com/tntmeijs/gengo/mathematics"

const epsilon = 0.001

type sdf func(point Vec3) float64

// Information about the surface a ray hit
type SurfaceHitInfo struct {
	Point, Normal Vec3
	RayLength     float64
}

// Represents a scene that can be rendered
type Scene struct {
	sceneSDF sdf
}

// Create a new scene
func NewScene(sceneSDF sdf) Scene {
	return Scene{sceneSDF}
}

// Check whether the point in space intersects with the scene's surface
func (s *Scene) DoesPointIntersectSurface(point Vec3) bool {
	return s.sceneSDF(point) <= 0.0
}

// Calculate the information at the position a point intersects the scene's surface
func (s *Scene) GetIntersectionPointSurfaceHitInfo(point Vec3, rayLength float64) SurfaceHitInfo {
	return SurfaceHitInfo{point, s.approximateNormal(point), rayLength}
}

// Approximate the surface normal by samping points around the intersection point
//
// Reference: http://jamie-wong.com/2016/07/15/ray-marching-signed-distance-functions/
func (s *Scene) approximateNormal(point Vec3) Vec3 {
	normalX := s.sceneSDF(Vec3{X: point.X + epsilon, Y: point.Y, Z: point.Z}) - s.sceneSDF(Vec3{X: point.X - epsilon, Y: point.Y, Z: point.Z})
	normalY := s.sceneSDF(Vec3{X: point.X, Y: point.Y + epsilon, Z: point.Z}) - s.sceneSDF(Vec3{X: point.X, Y: point.Y - epsilon, Z: point.Z})
	normalZ := s.sceneSDF(Vec3{X: point.X, Y: point.Y, Z: point.Z + epsilon}) - s.sceneSDF(Vec3{X: point.X, Y: point.Y, Z: point.Z - epsilon})

	return Normalize(Vec3{X: normalX, Y: normalY, Z: normalZ})
}
