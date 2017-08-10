package view

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/glesica/tictacgo/game"
	"github.com/glesica/tictacgo/game/marker"
	"github.com/glesica/tictacgo/game/square"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
)

var lineColor color.RGBA
var lineWidth float64

func init() {
	lineColor = colornames.Black
	lineWidth = 10
}

// Render draws the current game to the given screen.
func Render(state *game.State, window *pixelgl.Window) {
	width := window.Bounds().W()
	height := window.Bounds().H()

	w := width / 6
	h := height / 6

	// Draw the board

	makeBoard(width, height).Draw(window)

	// Draw the markers

	centers := []pixel.Vec{
		pixel.V(w, h*5),
		pixel.V(w*3, h*5),
		pixel.V(w*5, h*5),
		pixel.V(w, h*3),
		pixel.V(w*3, h*3),
		pixel.V(w*5, h*3),
		pixel.V(w, h),
		pixel.V(w*3, h),
		pixel.V(w*5, h),
	}

	var rad float64
	if width < height {
		rad = w
	} else {
		rad = h
	}

	for i, sq := range square.Positions {
		makeMarker(state.InSquare(sq), centers[i], rad, state.Winner() == state.InSquare(sq)).Draw(window)
	}
}

func makeBoard(width, height float64) *imdraw.IMDraw {
	brd := imdraw.New(nil)
	brd.Color = lineColor

	brd.Push(pixel.V(width/3, 0), pixel.V(width/3, height))
	brd.Line(lineWidth)

	brd.Push(pixel.V(width*2/3, 0), pixel.V(width*2/3, height))
	brd.Line(lineWidth)

	brd.Push(pixel.V(0, height/3), pixel.V(width, height/3))
	brd.Line(lineWidth)

	brd.Push(pixel.V(0, height*2/3), pixel.V(width, height*2/3))
	brd.Line(lineWidth)

	return brd
}

func makeMarker(symbol marker.T, center pixel.Vec, radius float64, isWinner bool) *imdraw.IMDraw {
	img := imdraw.New(nil)
	if isWinner {
		img.Color = colornames.Lawngreen
	} else {
		img.Color = colornames.Grey
	}

	// Define bounding boxes instead of center points, that way we get the size and location
	if symbol == marker.X {
		delta := radius / math.Sqrt2

		img.Push(pixel.V(center.X-delta, center.Y-delta))
		img.Push(pixel.V(center.X+delta, center.Y+delta))
		img.Line(lineWidth)

		img.Push(pixel.V(center.X-delta, center.Y+delta))
		img.Push(pixel.V(center.X+delta, center.Y-delta))
		img.Line(lineWidth)
	}

	if symbol == marker.O {
		img.Push(center)
		img.Circle(radius*0.8, lineWidth)
	}

	return img
}
