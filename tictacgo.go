package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/glesica/tictacgo/game"
	"github.com/glesica/tictacgo/view"
	"github.com/glesica/tictacgo/view/cli"
	"github.com/glesica/tictacgo/view/gui"
)

var viewer view.T

func run() {
	state := game.New()

	cfg := pixelgl.WindowConfig{
		Title:  "Tic-Tac-Go",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	viewer = gui.NewView(win)

	for !win.Closed() {
		viewer.Update(state)

		win.Update()
	}

	viewer = cli.NewView()

	viewer.Update(state)
}

func main() {
	pixelgl.Run(run)
}
