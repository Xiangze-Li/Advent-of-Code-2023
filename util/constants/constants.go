//nolint:gochecknoglobals // constants
package c

// Direction is a direction in 2D space.
type Direction byte

// Directions in 2D space. Index `i` grows South, index `j` grows East.
const (
	N Direction = 1 << iota
	E
	S
	W
	NE = N | E
	SE = S | E
	NW = N | W
	SW = S | W
)

// Delta8 is a map of 8 directions to their index delta.
var Delta8 = map[Direction][2]int{
	N:  {-1, 0},
	NE: {-1, 1},
	E:  {0, 1},
	SE: {1, 1},
	S:  {1, 0},
	SW: {1, -1},
	W:  {0, -1},
	NW: {-1, -1},
}

// Delta4 is a map of 4 directions to their index delta.
//
//nolint:exhaustive // intentionally
var Delta4 = map[Direction][2]int{
	N: {-1, 0},
	E: {0, 1},
	S: {1, 0},
	W: {0, -1},
}

// Opposite maps a direction to its opposite direction.
var Opposite = map[Direction]Direction{
	N:  S,
	E:  W,
	S:  N,
	W:  E,
	NE: SW,
	SE: NW,
	NW: SE,
	SW: NE,
}

// ConvertFromUDLR converts a byte of U/D/L/R to a Direction.
var ConvertFromUDLR = map[byte]Direction{
	'U': N,
	'D': S,
	'L': W,
	'R': E,
}
