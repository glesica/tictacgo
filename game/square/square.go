package square

type T int

const (
	NW T = iota
	N
	NE
	W
	C
	E
	SW
	S
	SE
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
