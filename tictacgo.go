package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/glesica/tictacgo/game"
	"github.com/glesica/tictacgo/game/marker"
	"github.com/glesica/tictacgo/view"
	"golang.org/x/image/colornames"
)

func run() {
	state := game.New()

	cfg := pixelgl.WindowConfig{
		Title:  "Tic Tac Go",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Clear(colornames.White)

		if win.JustReleased(pixelgl.MouseButton1) && state.Winner() == marker.Empty {
			sq := view.WhichSquare(win.MousePosition(), win.Bounds())
			_ = state.Move(sq)
		}

		view.Render(state, win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
