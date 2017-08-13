package cli

import "github.com/glesica/tictacgo/game"

type View struct {}

func NewView() *View {
	return &View{}
}

func (v *View) Update(state *game.State) {

}
