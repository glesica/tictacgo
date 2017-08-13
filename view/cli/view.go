package cli

import (
	"github.com/glesica/tictacgo/game"
	"github.com/glesica/tictacgo/game/square"
)

type View struct{}

func NewView() *View {
	return &View{}
}

func (v *View) Update(state *game.State) {
	println(state.InSquare(square.NW) + "|" + state.InSquare(square.N) + "|" + state.InSquare(square.NE))
	println("-----")
	println(state.InSquare(square.W) + "|" + state.InSquare(square.C) + "|" + state.InSquare(square.E))
	println("-----")
	println(state.InSquare(square.SW) + "|" + state.InSquare(square.S) + "|" + state.InSquare(square.SE))
}
