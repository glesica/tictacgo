package game

import (
	"errors"
)

import (
	"github.com/glesica/tictacgo/game/marker"
	"github.com/glesica/tictacgo/game/square"
)

type State struct {
	board map[square.T]marker.T
	next  marker.T
}

func New() *State {
	state := State{
		board: make(map[square.T]marker.T),
		next:  marker.X,
	}

	// TODO: If we made Empty the zero value would we get this for free?
	for _, sq := range square.Positions {
		state.board[sq] = marker.Empty
	}

	return &state
}

// Retrieve the marker present in the given square.
func (s *State) InSquare(sq square.T) marker.T {
	return s.board[sq]
}

// Check to see whether the given square is occupied.
func (s *State) IsEmpty(sq square.T) bool {
	return s.board[sq] == marker.Empty
}

// Place a marker on the game board on behalf of the player whose
// turn is next. Return an error if the square is already occupied.
func (s *State) Move(sq square.T) error {
	if !s.IsEmpty(sq) {
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

// Return the marker belonging to the winning player or, if there is
// no winner, return the Empty marker.
func (s *State) Winner() marker.T {
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
