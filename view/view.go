package view

import "github.com/glesica/tictacgo/game"

type T interface {
	Update(state *game.State)
}
