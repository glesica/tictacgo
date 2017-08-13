package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/glesica/tictacgo/game"
	"github.com/glesica/tictacgo/game/marker"
	"github.com/glesica/tictacgo/view"
	"github.com/glesica/tictacgo/view/gui"
	"golang.org/x/image/colornames"
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
		win.Clear(colornames.White)

		if win.JustReleased(pixelgl.MouseButton1) && state.Winner() == marker.Empty {
			sq := view.WhichSquare(win.MousePosition(), win.Bounds())
			_ = state.Move(sq)
		}

		viewer.Update(state)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
