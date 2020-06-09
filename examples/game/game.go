package main

import (
	"github.com/gdamore/tcell"
	"github.com/liamg/raytracing/raytracer"
)

type Game struct {
	screen tcell.Screen
	scene  *raytracer.Scene
	camera Camera
}

type Camera struct {
	position  raytracer.Vector
	xRotation float64
	yRotation float64
}

func (game *Game) Init() error {
	var err error
	game.screen, err = tcell.NewScreen()
	if err != nil {
		return err
	}

	if err := game.screen.Init(); err != nil {
		return err
	}

	w, h := game.screen.Size()
	canvasSize := raytracer.Vector{X: float64(w), Y: float64(h)}

	game.scene = raytracer.NewScene(raytracer.WithCanvasSize(canvasSize))
	game.scene.AddLight(raytracer.NewAmbientLight(raytracer.Colour{1.0, 1.0, 1.0}))
	game.camera.position.Y = -1

	return nil
}

func (game *Game) Close() {
	game.screen.Fini()
}

func (game *Game) PlayLevel(level *Level) error {

	reflectionDepth := 0

	w, h := game.screen.Size()
	level.Load(game.scene)

	var style tcell.Style

	for {

		game.screen.Clear()

		for x := -w / 2; x < w/2; x++ {
			for y := -h / 2; y < h/2; y++ {
				dest := game.scene.CanvasToViewport(raytracer.Vector{X: float64(x), Y: float64(y)})
				dest = raytracer.RotateXY(dest, game.camera.xRotation, game.camera.yRotation)

				colour, _ := game.scene.TraceRay(game.camera.position, dest, 0.5, 100, reflectionDepth)
				style = tcell.StyleDefault.Background(tcell.NewRGBColor(int32(colour.R*255), int32(colour.G*255), int32(colour.B*255)))
				game.screen.SetContent(
					x+(w/2),
					y+(h/2),
					' ',
					nil,
					style,
				)
			}
		}

		game.screen.Show()

		ev := game.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyUp:
				dest := game.scene.CanvasToViewport(raytracer.Vector{X: float64(0), Y: float64(0)})
				dest = raytracer.RotateXY(dest, game.camera.xRotation, game.camera.yRotation)
				if _, t := game.scene.TraceRay(game.camera.position, dest, 1.00, 2.00, 0); t >= 2 {
					game.camera.position = game.camera.position.Add(dest)
				}
			case tcell.KeyDown:
				dest := game.scene.CanvasToViewport(raytracer.Vector{X: float64(0), Y: float64(0)})
				dest = raytracer.RotateXY(dest, game.camera.xRotation, game.camera.yRotation).Reverse()
				if _, t := game.scene.TraceRay(game.camera.position, dest, 1.00, 2.00, 0); t >= 2 {
					game.camera.position = game.camera.position.Add(dest)
				}
			case tcell.KeyLeft:
				game.camera.yRotation -= 0.1
			case tcell.KeyRight:
				game.camera.yRotation += 0.1
			case tcell.KeyEscape:
				// pause menu
				return nil
			}
		case *tcell.EventResize:
			w, h = ev.Size()
			game.scene.Resize(w, h)
		}

	}

	return nil
}
