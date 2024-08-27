package main

import (
	"image/color"
	"os"
	"image/png"
	"raytracer/raytracer"
	"raytracer/render"
)

func main() {
	const (
		imageWidth  = 800
		imageHeight = 400
		samplesPerPixel = 100 // Number of samples per pixel for anti-aliasing
	)

	// Create a world and add spheres
	world := raytracer.World{}
	world.Add(raytracer.Sphere{
		Center: raytracer.Vec3{0, 0, -1},
		Radius: 0.5,
		Color:  color.RGBA{255, 0, 0, 255},
	})
	world.Add(raytracer.Sphere{
		Center: raytracer.Vec3{0, -100.5, -1},
		Radius: 100,
		Color:  color.RGBA{0, 255, 0, 255},
	})

	// Render the scene with anti-aliasing
	img := render.RenderScene(imageWidth, imageHeight, samplesPerPixel, world)

	// Save the image to a file
	file, err := os.Create("dist/output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)
}
