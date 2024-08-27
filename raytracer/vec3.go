package raytracer

import "math"

// Vec3 represents a 3D vector
type Vec3 struct {
	X, Y, Z float64
}

// Add adds two vectors
func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

// Sub subtracts two vectors
func (v Vec3) Sub(u Vec3) Vec3 {
	return Vec3{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

// Mul multiplies the vector by a scalar
func (v Vec3) Mul(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

// Dot returns the dot product of two vectors
func (v Vec3) Dot(u Vec3) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

// Length returns the length of the vector
func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Unit returns the unit vector
func (v Vec3) Unit() Vec3 {
	length := v.Length()
	return Vec3{v.X / length, v.Y / length, v.Z / length}
}

// Reflect returns the reflection of a vector off a surface
func (v Vec3) Reflect(normal Vec3) Vec3 {
	return v.Sub(normal.Mul(2 * v.Dot(normal)))
}
