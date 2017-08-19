package cli

import (
	"bufio"
	"fmt"
	"github.com/glesica/tictacgo/game"
	"github.com/glesica/tictacgo/game/square"
	"github.com/glesica/tictacgo/view"
	"os"
)

type cliView struct{}

func New() view.T {
	return &cliView{}
}

func (v *cliView) Update(state game.T) {
	println(state.InSquare(square.NW) + "|" + state.InSquare(square.N) + "|" + state.InSquare(square.NE))
	println("-----")
	println(state.InSquare(square.W) + "|" + state.InSquare(square.C) + "|" + state.InSquare(square.E))
	println("-----")
	println(state.InSquare(square.SW) + "|" + state.InSquare(square.S) + "|" + state.InSquare(square.SE))

	fmt.Printf("Next move (%s): ", state.Turn())

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	pos, err := square.FromString(input)
	if err != nil {
		return
	}

	state.Move(pos)
}
