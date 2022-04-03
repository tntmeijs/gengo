package scene

import . "github.com/tntmeijs/gengo/mathematics"

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
	return SurfaceHitInfo{point, Vec3{}, rayLength}
}
