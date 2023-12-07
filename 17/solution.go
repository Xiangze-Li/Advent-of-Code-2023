package puzzle17

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	c "advent2023/util/constants"
	"container/heap"
	"errors"
	"strings"
)

type p struct {
	grids [][]uint64
}

func (p *p) Init(filename string) {
	lines := util.GetLines(filename)
	p.grids = make([][]uint64, 0, len(lines))
	for _, line := range lines {
		p.grids = append(p.grids, util.ArrayStrToUint64(strings.Split(line, "")))
	}
}

func (p *p) solve(mustTurn func(prev, curr c.Direction, count int) bool) uint64 {
	lenI, lenJ := len(p.grids), len(p.grids[0])
	inBounds := func(i, j int) bool {
		return i >= 0 && i < lenI && j >= 0 && j < lenJ
	}
	closed := util.SliceND[map[string]bool](lenI, lenJ).([][]map[string]bool)
	for i := 0; i < lenI; i++ {
		for j := 0; j < lenJ; j++ {
			closed[i][j] = make(map[string]bool)
		}
	}
	open := stateHeap{
		{0, 0, []c.Direction{c.E}, 0},
		{0, 0, []c.Direction{c.S}, 0},
	}
	heap.Init(&open)

	for len(open) > 0 {
		curr := heap.Pop(&open).(state)
		currDir := curr.going[len(curr.going)-1]
		if closed[curr.i][curr.j][string(curr.going)] {
			continue
		}
		closed[curr.i][curr.j][string(curr.going)] = true
		for dir, delta := range c.Delta4 {
			i, j := curr.i+delta[0], curr.j+delta[1]
			if !inBounds(i, j) ||
				mustTurn(currDir, dir, len(curr.going)) ||
				c.Opposite[dir] == currDir {
				continue
			}
			path := []c.Direction{dir}
			if dir == currDir {
				path = append(curr.going, dir) //nolint:gocritic // intentional
			}
			if closed[i][j][string(path)] {
				continue
			}
			cost := curr.cost + p.grids[i][j]
			if i == lenI-1 && j == lenJ-1 {
				return cost
			}
			heap.Push(&open, state{i, j, path, cost})
		}
	}

	return util.Must[uint64](0, errors.New("no path found"))
}

func (p *p) Solve1() any {
	return p.solve(func(prev, curr c.Direction, count int) bool {
		return prev == curr && count >= 3
	})
}

func (p *p) Solve2() any {
	return p.solve(func(prev, curr c.Direction, count int) bool {
		return (prev != curr && count < 4) || (prev == curr && count >= 10)
	})
}

func init() {
	puzzle.Register(17, &p{})
}
