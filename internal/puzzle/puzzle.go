package puzzle

import "fmt"

type Puzzle interface {
	Init(filename string)
	Solve1() any
	Solve2() any
}

var puzzles = make(map[int]Puzzle) //nolint:gochecknoglobals // global registry

func Register(day int, puzzle Puzzle) {
	if _, exist := puzzles[day]; exist {
		panic(fmt.Errorf("duplicate registration for puzzle %d", day))
	}
	puzzles[day] = puzzle
}

func Get(day int) (Puzzle, error) {
	puzzle, exist := puzzles[day]
	if !exist {
		return nil, fmt.Errorf("no puzzle registered for day %d", day)
	}
	return puzzle, nil
}
