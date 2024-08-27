package raytracer

import (
	"image/color"
	"math"
)

// Sphere represents a sphere in 3D space
type Sphere struct {
	Center Vec3
	Radius float64
	Color  color.RGBA
}

// Hit checks if a ray hits the sphere
func (s Sphere) Hit(r Ray) (bool, float64) {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		return true, (-b - math.Sqrt(discriminant)) / (2.0 * a)
	}
	return false, -1
}
