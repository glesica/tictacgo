package gui

import (
	"github.com/faiface/pixel"
	"github.com/glesica/tictacgo/game/square"
)

func WhichSquare(position pixel.Vec, bounds pixel.Rect) square.T {
	if position.X < bounds.W()/3 {
		if position.Y < bounds.H()/3 {
			return square.SW
		}

		if position.Y < bounds.H()*2/3 {
			return square.W
		}

		return square.NW
	}

	if position.X < bounds.W()*2/3 {
		if position.Y < bounds.H()/3 {
			return square.S
		}

		if position.Y < bounds.H()*2/3 {
			return square.C
		}

		return square.N
	}

	if position.Y < bounds.H()/3 {
		return square.SE
	}

	if position.Y < bounds.H()*2/3 {
		return square.E
	}

	return square.NE
}
