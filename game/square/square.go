package square

import (
	"github.com/pkg/errors"
	"strings"
)

type T string

const (
	NW T = "NW"
	N  T = "N"
	NE T = "NE"
	W  T = "W"
	C  T = "C"
	E  T = "E"
	SW T = "SW"
	S  T = "S"
	SE T = "SE"
)

var Positions [9]T = [9]T{
	NW,
	N,
	NE,
	W,
	C,
	E,
	SW,
	S,
	SE,
}

var WinPositions [8][3]T = [8][3]T{
	[3]T{NW, N, NE},
	[3]T{W, C, E},
	[3]T{SW, S, SE},
	[3]T{NW, W, SW},
	[3]T{N, C, S},
	[3]T{NE, E, SE},
	[3]T{NW, C, SE},
	[3]T{NE, C, SW},
}

var stringToPosition map[string]T = map[string]T{
	"NW": NW,
	"N":  N,
	"NE": NE,
	"W":  W,
	"C":  C,
	"E":  E,
	"SW": SW,
	"S":  S,
	"SE": SE,
}

// Convert a string to a position. Trims leading and trailing
// whitespace and converts to upper case first. Returns an error
// if the string does not correspond to a position.
func FromString(str string) (T, error) {
	trimmedStr := strings.ToUpper(strings.TrimSpace(str))
	pos, ok := stringToPosition[trimmedStr]
	if ok {
		return pos, nil
	}

	return "", errors.Errorf("Invalid position: %s", str)
}
