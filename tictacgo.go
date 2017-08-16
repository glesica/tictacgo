package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/glesica/tictacgo/game"
	"github.com/glesica/tictacgo/view/cli"
	"github.com/glesica/tictacgo/view/gui"
	"flag"
	"os"
)

var state game.T

func runGui() {
	cfg := pixelgl.WindowConfig{
		Title:  "Tic-Tac-Go",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	view := gui.New(win)

	for !win.Closed() {
		view.Update(state)
		win.Update()
	}
}

func runCli() {
	view := cli.New()

	for !state.IsOver() {
		view.Update(state)
	}
}

func main() {
	showHelp := flag.Bool("help", false, "Show help")
	useGui := flag.Bool("gui", false, "Run the OpenGL GUI")

	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	state = game.New()

	if *useGui {
		pixelgl.Run(runGui)
	} else {
		runCli()
	}
}
