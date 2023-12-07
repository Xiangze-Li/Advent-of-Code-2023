package puzzle21

import (
	"advent2023/internal/puzzle"
	"advent2023/util"
	c "advent2023/util/constants"
)

type p struct {
	grid  [][]byte
	entry [2]int
}

func (p *p) Init(filename string) {
	p.grid = util.GetGrid(filename)

	for i, row := range p.grid {
		for j, c := range row {
			if c == 'S' {
				p.entry = [2]int{i, j}
				p.grid[i][j] = '.'
				return
			}
		}
	}
}

func (p *p) Solve1() any {
	const timeLimit = 64

	vis := map[[2]int]bool{}
	queue := [][3]int{{p.entry[0], p.entry[1], 0}}
	even := 0
	lenI, lenJ := len(p.grid), len(p.grid[0])

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur[2] > timeLimit {
			continue
		}
		if vis[[2]int{cur[0], cur[1]}] {
			continue
		}
		vis[[2]int{cur[0], cur[1]}] = true

		if cur[2]%2 == 0 {
			even++
		}

		for _, delta := range c.Delta4 {
			i, j := cur[0]+delta[0], cur[1]+delta[1]
			if i < 0 || i >= lenI || j < 0 || j >= lenJ {
				continue
			}
			if p.grid[i][j] == '#' {
				continue
			}
			queue = append(queue, [3]int{i, j, cur[2] + 1})
		}
	}
	return even
}

func (p *p) Solve2() any {
	const timeLimit = 26501365
	length := len(p.grid)

	queue := map[[2]int]bool{p.entry: true}
	a := []int{}

	for step := 1; true; step++ {
		nextQ := map[[2]int]bool{}

		for cur := range queue {
			for _, delta := range c.Delta4 {
				nextI, nextJ := cur[0]+delta[0], cur[1]+delta[1]
				i, j := nextI%length, nextJ%length
				if i < 0 {
					i += length
				}
				if j < 0 {
					j += length
				}
				if p.grid[i][j] == '#' {
					continue
				}
				nextQ[[2]int{nextI, nextJ}] = true
			}
		}

		queue = nextQ
		if step%length == timeLimit%length {
			a = append(a, len(queue))
			if len(a) == 3 {
				break
			}
		}
	}

	poly := func(n int) int {
		a0, a1, a2 := a[0], a[1], a[2]
		return a0 + (a1-a0)*n + (a2-2*a1+a0)*n*(n-1)/2
	}

	return poly(timeLimit / length)
}

func init() {
	puzzle.Register(21, &p{})
}
