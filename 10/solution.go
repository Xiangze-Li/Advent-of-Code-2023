//nolint:gocyclo, cyclop // It's just that much complicated.
package puzzle10

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	c "advent2023/util/constants"
)

type p struct {
	grid  [][]pipe
	start [2]int
}

func (p *p) calcStartDir(sI, sJ int) c.Direction {
	pipeDirSS := []c.Direction{}
	lenI, lenJ := len(p.grid), len(p.grid[0])

	for dir, delta := range c.Delta4 {
		i, j := sI+delta[0], sJ+delta[1]
		if i < 0 || i >= lenI || j < 0 || j >= lenJ {
			continue
		}
		pipeDir := p.grid[i][j].direction
		if pipeDir == conn00 {
			continue
		}
		if (dir == c.N && (pipeDir == conn13 || pipeDir == conn23 || pipeDir == conn34)) ||
			(dir == c.E && (pipeDir == conn24 || pipeDir == conn34 || pipeDir == conn41)) ||
			(dir == c.S && (pipeDir == conn12 || pipeDir == conn13 || pipeDir == conn41)) ||
			(dir == c.W && (pipeDir == conn12 || pipeDir == conn23 || pipeDir == conn24)) {
			pipeDirSS = append(pipeDirSS, dir)
		}
	}

	util.Assert(len(pipeDirSS) == 2, "invalid pipeDirSS")

	if pipeDirSS[0] == c.N && pipeDirSS[1] == c.E {
		p.grid[sI][sJ].direction = conn12
		return c.N
	}
	if pipeDirSS[0] == c.N && pipeDirSS[1] == c.S {
		p.grid[sI][sJ].direction = conn13
		return c.N
	}
	if pipeDirSS[0] == c.N && pipeDirSS[1] == c.W {
		p.grid[sI][sJ].direction = conn41
		return c.N
	}
	if pipeDirSS[0] == c.E && pipeDirSS[1] == c.S {
		p.grid[sI][sJ].direction = conn23
		return c.E
	}
	if pipeDirSS[0] == c.E && pipeDirSS[1] == c.W {
		p.grid[sI][sJ].direction = conn24
		return c.E
	}
	if pipeDirSS[0] == c.S && pipeDirSS[1] == c.W {
		p.grid[sI][sJ].direction = conn34
		return c.S
	}

	return 0
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.grid = util.SliceND[pipe](len(lines), len(lines[0])).([][]pipe)

	for i, line := range lines {
		for j, c := range line {
			dir := convertDirection(c)
			p.grid[i][j] = pipe{direction: dir, i: i, j: j}
			if dir == connSS {
				p.start = [2]int{i, j}
			}
		}
	}

	i, j := p.start[0], p.start[1]
	curPipe := &p.grid[i][j]
	var prevDir c.Direction
	for {
		var nextDir c.Direction
		switch curPipe.direction {
		case conn00: // no-op
		case connSS:
			nextDir = p.calcStartDir(i, j)
		case conn24, conn13:
			nextDir = prevDir
		case conn12:
			if prevDir == c.S {
				nextDir = c.E
			} else {
				nextDir = c.N
			}
		case conn34:
			if prevDir == c.N {
				nextDir = c.W
			} else {
				nextDir = c.S
			}
		case conn23:
			if prevDir == c.N {
				nextDir = c.E
			} else {
				nextDir = c.S
			}
		case conn41:
			if prevDir == c.S {
				nextDir = c.W
			} else {
				nextDir = c.N
			}
		}
		i, j = i+c.Delta4[nextDir][0], j+c.Delta4[nextDir][1]
		curPipe.next = &(p.grid[i][j])
		curPipe = curPipe.next
		prevDir = nextDir
		if i == p.start[0] && j == p.start[1] {
			break
		}
	}
}

func (p *p) Solve1() any {
	ptrFast := &(p.grid[p.start[0]][p.start[1]])
	step := uint64(0)
	for {
		ptrFast = ptrFast.next.next
		step++
		if ptrFast.i == p.start[0] && ptrFast.j == p.start[1] {
			break
		}
	}
	return step
}

func (p *p) Solve2() any {
	loop := map[[2]int]bool{}
	for p := &(p.grid[p.start[0]][p.start[1]]); true; {
		key := [2]int{p.i, p.j}
		if loop[key] {
			break
		}
		loop[key] = true
		p = p.next
	}

	count := 0

	for i := 0; i < len(p.grid); i++ {
		crossed := 0
		for j := 0; j < len(p.grid[0]); j++ {
			if loop[[2]int{i, j}] {
				switch p.grid[i][j].direction {
				case conn12, conn13, conn41:
					crossed++
				default: // no-op
				}
			} else if crossed%2 == 1 {
				count++
			}
		}
	}

	return count
}

func init() {
	puzzle.Register(10, &p{})
}
