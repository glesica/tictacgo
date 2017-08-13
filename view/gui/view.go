package gui

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

const lineWidth = 10

var backgroundColor color.RGBA = colornames.White
var boardColor color.RGBA = colornames.Black
var player0Color color.RGBA = colornames.Red
var player1Color color.RGBA = colornames.Blue
var winnerColor color.RGBA = colornames.Darkgreen

type View struct {
	backgroundColor color.RGBA
	boardColor      color.RGBA
	player0Color    color.RGBA
	player1Color    color.RGBA
	window          *pixelgl.Window
	winnerColor     color.RGBA
}

// NewView constructs a new View struct that will draw on the
// window provided.
func NewView(window *pixelgl.Window) *View {
	return &View{
		backgroundColor: backgroundColor,
		boardColor:      boardColor,
		player0Color:    player0Color,
		player1Color:    player1Color,
		window:          window,
		winnerColor:     winnerColor,
	}
}

// Update the drawn version of the game state.
func (v *View) Update(state *game.State) {
	v.window.Clear(v.backgroundColor)

	width := v.window.Bounds().W()
	height := v.window.Bounds().H()

	widthSegment := width / 6
	heightSegment := height / 6

	// Draw the board

	v.makeBoard(width, height).Draw(v.window)

	// Draw the markers

	centers := []pixel.Vec{
		pixel.V(widthSegment, heightSegment*5),
		pixel.V(widthSegment*3, heightSegment*5),
		pixel.V(widthSegment*5, heightSegment*5),
		pixel.V(widthSegment, heightSegment*3),
		pixel.V(widthSegment*3, heightSegment*3),
		pixel.V(widthSegment*5, heightSegment*3),
		pixel.V(widthSegment, heightSegment),
		pixel.V(widthSegment*3, heightSegment),
		pixel.V(widthSegment*5, heightSegment),
	}

	var rad float64
	if width < height {
		rad = widthSegment
	} else {
		rad = heightSegment
	}

	for i, sq := range square.Positions {
		v.makeMarker(state.InSquare(sq), centers[i], rad, state.Winner() == state.InSquare(sq)).Draw(v.window)
	}
}

func (v *View) makeBoard(width, height float64) *imdraw.IMDraw {
	brd := imdraw.New(nil)
	brd.Color = v.boardColor

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

func (v *View) makeMarker(symbol marker.T, center pixel.Vec, radius float64, isWinner bool) *imdraw.IMDraw {
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
