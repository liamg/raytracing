package main

import (
	"github.com/fogleman/gg"
	"github.com/liamg/raytracing/raytracer"
)

func main() {

	canvasSize := raytracer.Vector{X: 1000, Y: 1000}

	scene := raytracer.NewScene(raytracer.WithCanvasSize(canvasSize), raytracer.WithBackgroundColour(raytracer.Colour{}))
	scene.AddObject(raytracer.NewSphere(
		raytracer.Vector{X: 0, Y: 0, Z: 3},
		1,
		raytracer.Colour{R: 1, G: 0, B: 0},
		50,
		0.4,
	))
	scene.AddObject(raytracer.NewSphere(
		raytracer.Vector{X: -2, Y: 1, Z: 3},
		1,
		raytracer.Colour{R: 0, G: 0, B: 1},
		100,
		0.3,
	))
	scene.AddObject(raytracer.NewSphere(
		raytracer.Vector{X: 0.9, Y: -0.5, Z: 2},
		0.1,
		raytracer.Colour{R: 0.5, G: 1, B: 0},
		100,
		0.3,
	))
	scene.AddObject(raytracer.NewSphere(
		raytracer.Vector{X: 0.6, Y: -0.7, Z: 2},
		0.15,
		raytracer.Colour{R: 1, G: 1, B: 0},
		100,
		0.3,
	))
	scene.AddObject(raytracer.NewSphere(
		raytracer.Vector{X: 0.4, Y: -0.9, Z: 2},
		0.2,
		raytracer.Colour{R: 1, G: 0.4, B: 0},
		100,
		0.3,
	))
	scene.AddObject(raytracer.NewBox(
		raytracer.Vector{X: -1, Y: 1, Z: 2},
		raytracer.Vector{X: 5, Y: 0.5, Z: 5},
		raytracer.Colour{R: 1, G: 1, B: 1},
		100,
		0.7,
	))

	scene.AddLight(raytracer.NewAmbientLight(raytracer.Colour{0.4, 0.4, 0.4}))
	scene.AddLight(raytracer.NewPointLight(raytracer.Colour{0.7, 0.7, 0.7}, raytracer.Vector{7, -70, 0}))
	scene.AddLight(raytracer.NewPointLight(raytracer.Colour{0.5, 0.5, 0.5}, raytracer.Vector{80, -7, 0}))
	scene.AddLight(raytracer.NewPointLight(raytracer.Colour{0.3, 0.4, 0.4}, raytracer.Vector{0, 0, 0}))
	scene.AddLight(raytracer.NewPointLight(raytracer.Colour{0.1, 0.4, 0.4}, raytracer.Vector{-2, -2, 0}))

	dc := gg.NewContext(int(canvasSize.X), int(canvasSize.Y))

	origin := raytracer.Vector{
		X: -2,
		Y: -1,
	}

	for x := -int(canvasSize.X) / 2; x < int(canvasSize.X)/2; x++ {
		for y := -int(canvasSize.Y) / 2; y < int(canvasSize.Y)/2; y++ {
			dest := scene.CanvasToViewport(raytracer.Vector{X: float64(x), Y: float64(y)})
			dest = raytracer.RotateXY(dest, -0.2, 0.6)
			colour, _ := scene.TraceRay(origin, dest, 1, 100, 4)
			dc.SetRGB(colour.R, colour.G, colour.B)
			dc.SetPixel(x+(int(canvasSize.X)/2), (int(canvasSize.Y)/2)+y)
		}
	}

	dc.SavePNG("demo.png")
}
