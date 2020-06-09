package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/liamg/raytracing/raytracer"
)

type Level struct {
	grid [32][32]bool
}

func NewLevel() *Level {
	return &Level{
		grid: [32][32]bool{},
	}
}

func (l *Level) Randomise() {
	l.grid[17][17] = true
}

func (l *Level) SetFromString(s string) {

	l.grid = [32][32]bool{}

	for x, line := range strings.Split(s, "\n") {
		if x >= 32 {
			break
		}
		for y, r := range line {
			if y >= 32 {
				continue
			}
			l.grid[x][y] = r != ' '
		}
	}
}

func (l *Level) Load(scene *raytracer.Scene) {

	scene.ResetObjects()

	//draw ground
	scene.AddObject(raytracer.NewBox(
		raytracer.Vector{X: -1000, Y: 0, Z: -1000},
		raytracer.Vector{X: 2000, Y: 1, Z: 2000},
		raytracer.Colour{R: 0.3, G: 0.3, B: 0.3},
		0,
		0,
	))

	rand.Seed(time.Now().UnixNano())

	// draw walls
	for x := 0; x < 32; x++ {
		for z := 0; z < 32; z++ {
			if l.grid[x][z] {
				scene.AddObject(raytracer.NewBox(
					raytracer.Vector{X: float64(x) - 16, Y: 0, Z: float64(z) - 16},
					raytracer.Vector{X: 1, Y: -1, Z: 1},
					raytracer.Colour{R: rand.Float64(), G: rand.Float64(), B: rand.Float64()},
					0,
					0,
				))
			}
		}
	}

}
