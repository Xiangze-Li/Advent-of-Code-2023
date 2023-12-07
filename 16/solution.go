package puzzle16

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	c "advent2023/util/constants"
)

type p struct {
	grid [][]byte
}

func (p *p) Init(filename string) {
	p.grid = util.GetGrid(filename)
}

type light struct {
	i, j int
	dir  c.Direction
}

//nolint:exhaustive,funlen,gocognit,gocyclo,cyclop // intentionally
func (p *p) handleLight(init light) int64 {
	lenI, lenJ := len(p.grid), len(p.grid[0])
	inBounds := func(i, j int) bool {
		return i >= 0 && i < lenI && j >= 0 && j < lenJ
	}

	lights := []light{init}
	visited := util.SliceND[c.Direction](lenI, lenJ).([][]c.Direction)

	for len(lights) > 0 {
		l := lights[0]
		lights = lights[1:]
		for {
			if visited[l.i][l.j]&l.dir != 0 {
				break
			}
			visited[l.i][l.j] |= l.dir

			switch p.grid[l.i][l.j] {
			case '.': // no-op
			case '/':
				switch l.dir {
				case c.N:
					l.dir = c.E
				case c.S:
					l.dir = c.W
				case c.W:
					l.dir = c.S
				case c.E:
					l.dir = c.N
				}
			case '\\':
				switch l.dir {
				case c.N:
					l.dir = c.W
				case c.S:
					l.dir = c.E
				case c.W:
					l.dir = c.N
				case c.E:
					l.dir = c.S
				}
			case '|':
				switch l.dir {
				case c.N, c.S: // no-op
				case c.W, c.E:
					l.dir = c.N
					lights = append(lights, light{l.i, l.j, c.S})
				}
			case '-':
				switch l.dir {
				case c.W, c.E: // no-op
				case c.N, c.S:
					l.dir = c.W
					lights = append(lights, light{l.i, l.j, c.E})
				}
			}

			switch l.dir {
			case c.N:
				l.i--
			case c.S:
				l.i++
			case c.W:
				l.j--
			case c.E:
				l.j++
			}
			if !inBounds(l.i, l.j) {
				break
			}
		}
	}

	var count int64
	for i := 0; i < lenI; i++ {
		for j := 0; j < lenJ; j++ {
			if visited[i][j] != 0 {
				count++
			}
		}
	}
	return count
}

func (p *p) Solve1() any {
	return p.handleLight(light{i: 0, j: 0, dir: c.E})
}

func (p *p) Solve2() any {
	var maxium int64
	lenI, lenJ := len(p.grid), len(p.grid[0])
	for i := 0; i < lenI; i++ {
		maxium = max(maxium, p.handleLight(light{i, 0, c.E}), p.handleLight(light{i, lenJ - 1, c.W}))
	}
	for j := 0; j < lenJ; j++ {
		maxium = max(maxium, p.handleLight(light{0, j, c.S}), p.handleLight(light{lenI - 1, j, c.N}))
	}
	return maxium
}

func init() {
	puzzle.Register(16, &p{})
}
