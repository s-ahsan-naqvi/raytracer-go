package raytracer

import "math"

type World struct {
	Objects []Sphere
}

func (w *World) Add(s Sphere) {
	w.Objects = append(w.Objects, s)
}

func (w *World) Hit(r Ray) (bool, float64, Sphere) {
	hitAnything := false
	closestSoFar := math.MaxFloat64
	var hitObject Sphere

	for _, obj := range w.Objects {
		if hit, t := obj.Hit(r); hit && t < closestSoFar {
			hitAnything = true
			closestSoFar = t
			hitObject = obj
		}
	}

	return hitAnything, closestSoFar, hitObject
}
