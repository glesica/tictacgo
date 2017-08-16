package game

import (
	"errors"
)

import (
	"github.com/glesica/tictacgo/game/marker"
	"github.com/glesica/tictacgo/game/square"
)

type T interface {
	// The marker current in the given square.
	InSquare(sq square.T) marker.T

	// Place the next player's marker in the given square. Return
	// an error if the position is already occupied.
	Move(sq square.T) error

	// Whether the game is over (either a draw or win).
	IsOver() bool

	// Which marker is up next?
	Turn() marker.T

	// Which marker has won (empty if none).
	Winner() marker.T
}

type state struct {
	board map[square.T]marker.T
	next  marker.T
}

func New() T {
	state := state{
		board: make(map[square.T]marker.T),
		next:  marker.X,
	}

	for _, sq := range square.Positions {
		state.board[sq] = marker.Empty
	}

	return &state
}

func (s *state) InSquare(sq square.T) marker.T {
	return s.board[sq]
}

func (s *state) Move(sq square.T) error {
	if s.board[sq] != marker.Empty {
		return errors.New("Square already occupied")
	}

	s.board[sq] = s.next

	if s.next == marker.X {
		s.next = marker.O
	} else {
		s.next = marker.X
	}

	return nil
}

func (s *state) IsOver() bool {
	for _, mkr := range s.board {
		if mkr == marker.Empty {
			return s.Winner() != marker.Empty
		}
	}

	return true
}

func (s *state) Turn() marker.T {
	return s.next
}

func (s *state) Winner() marker.T {
	for _, positions := range square.WinPositions {
		p0 := positions[0]
		p1 := positions[1]
		p2 := positions[2]
		if s.board[p0] == s.board[p1] && s.board[p1] == s.board[p2] {
			return s.board[p0]
		}
	}

	return marker.Empty
}
