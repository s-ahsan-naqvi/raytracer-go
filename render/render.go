package render

import (
	"image"
	"image/color"
	"log"
	"math/rand"
	"raytracer/raytracer"
	"time"
)

// RayColor returns the color of the ray based on what it hits
func RayColor(r raytracer.Ray, world raytracer.World, depth int) color.RGBA {
	if depth <= 0 {
		return color.RGBA{0, 0, 0, 255} // Return black if recursion depth is exceeded
	}

	if hit, t, obj := world.Hit(r); hit {
		hitPoint := r.Origin.Add(r.Direction.Mul(t))
		normal := hitPoint.Sub(obj.Center).Unit()
		reflectedDirection := r.Direction.Reflect(normal)
		reflectedRay := raytracer.Ray{hitPoint, reflectedDirection}

		reflectedColor := RayColor(reflectedRay, world, depth-1)
		lightDir := raytracer.Vec3{1, 1, -1}.Unit()
		lightIntensity := max(0, normal.Dot(lightDir))

		finalColor := raytracer.Vec3{
			X: float64(obj.Color.R) * lightIntensity,
			Y: float64(obj.Color.G) * lightIntensity,
			Z: float64(obj.Color.B) * lightIntensity,
		}.Mul(0.5).Add(raytracer.Vec3{
			X: float64(reflectedColor.R),
			Y: float64(reflectedColor.G),
			Z: float64(reflectedColor.B),
		}.Mul(0.5))

		return color.RGBA{
			R: uint8(finalColor.X),
			G: uint8(finalColor.Y),
			B: uint8(finalColor.Z),
			A: 255,
		}
	}

	unitDirection := r.Direction.Unit()
	t := 0.5 * (unitDirection.Y + 1.0)
	white := raytracer.Vec3{1.0, 1.0, 1.0}
	blue := raytracer.Vec3{X: 0.5, Y: 0.7, Z: 1.0}
	blend := white.Mul(1.0 - t).Add(blue.Mul(t))
	return color.RGBA{
		R: uint8(255 * blend.X),
		G: uint8(255 * blend.Y),
		B: uint8(255 * blend.Z),
		A: 255,
	}
}

// max returns the maximum of two float64 values
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// RenderScene renders the scene and returns an image
func RenderScene(imageWidth, imageHeight, samplesPerPixel int, world raytracer.World) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	origin := raytracer.Vec3{X: 0, Y: 0, Z: 0}
	lowerLeftCorner := raytracer.Vec3{X: -2, Y: -1, Z: -1}
	horizontal := raytracer.Vec3{4, 0, 0}
	vertical := raytracer.Vec3{0, 2, 0}

	rand.Seed(time.Now().UnixNano())

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			var pixelColor raytracer.Vec3

			// Take multiple samples per pixel
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / float64(imageWidth)
				v := (float64(j) + rand.Float64()) / float64(imageHeight)
				ray := raytracer.Ray{
					Origin:    origin,
					Direction: lowerLeftCorner.Add(horizontal.Mul(u)).Add(vertical.Mul(v)),
				}
				colorVec := raytracer.Vec3{
					X: float64(RayColor(ray, world, 50).R),
					Y: float64(RayColor(ray, world, 50).G),
					Z: float64(RayColor(ray, world, 50).B),
				}
				pixelColor = pixelColor.Add(colorVec)
				log.Println("Pixel color:", pixelColor)
			}

			// Average the color and write to the image
			scale := 1.0 / float64(samplesPerPixel)
			finalColor := pixelColor.Mul(scale)

			img.Set(i, imageHeight-j-1, color.RGBA{
				R: uint8(finalColor.X),
				G: uint8(finalColor.Y),
				B: uint8(finalColor.Z),
				A: 255,
			})
		}
	}

	return img
}